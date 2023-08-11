package server

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/infrastructure/queue/rmq/adapter"
	"log"
	"time"
)

const (
	_defaultWaitTime = 5 * time.Second
	_defaultAttempts = 10
	_defaultTimeout  = 2 * time.Second
)

type CallHandler func(*amqp.Delivery) (interface{}, error)

type Server struct {
	conn    *adapter.RmqConnection
	error   chan error
	stop    chan struct{}
	router  map[string]CallHandler
	timeout time.Duration
}

func NewServer(url, serverExchange string, router map[string]CallHandler, opts ...Option) (*Server, error) {
	cfg := adapter.RmqConfig{
		URL:      url,
		WaitTime: _defaultWaitTime,
		Attempts: _defaultAttempts,
	}

	s := &Server{
		conn:    adapter.New(serverExchange, cfg),
		error:   make(chan error),
		stop:    make(chan struct{}),
		router:  router,
		timeout: _defaultTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	err := s.conn.AttemptsConnect()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ server: %w", err)
	}

	go s.consumer()

	return s, nil
}

func (s *Server) consumer() {
	for {
		select {
		case <-s.stop:
			return
		case d, opened := <-s.conn.Delivery:
			if !opened {
				s.reconnect()
				return
			}
			_ = d.Ack(false)
			s.serverCall(&d)
		}
	}
}

func (s *Server) serverCall(d *amqp.Delivery) {
	callHandler, ok := s.router[d.Type]
	if !ok {
		s.publish(d, nil, constants.ErrorBadRequest.Error())
		return
	}

	response, err := callHandler(d)
	if err != nil {
		s.publish(d, nil, constants.ErrorInternalServer.Error())
		log.Printf("Failed to handle call: %s", err)
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %s", err)
	}
	s.publish(d, body, constants.Success)
}

func (s *Server) publish(d *amqp.Delivery, body []byte, status string) {
	err := s.conn.Channel.Publish(d.ReplyTo, "", false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: d.CorrelationId,
		Type:          status,
		Body:          body,
	})
	if err != nil {
		log.Printf("Failed to publish response: %s", err)
	}
}

func (s *Server) reconnect() {
	close(s.stop)
	err := s.conn.AttemptsConnect()
	if err != nil {
		s.error <- err
		close(s.error)
		return
	}

	s.stop = make(chan struct{})
	go s.consumer()
}

func (s *Server) Notify() <-chan error {
	return s.error
}

func (s *Server) Shutdown() error {
	select {
	case <-s.error:
		return nil
	default:
	}

	close(s.stop)
	time.Sleep(s.timeout)

	err := s.conn.Connection.Close()
	if err != nil {
		return fmt.Errorf("Failed to close RabbitMQ connection: %w", err)
	}
	return nil
}
