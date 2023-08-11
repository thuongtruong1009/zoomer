package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/infrastructure/queue/rmq/adapter"
	"sync"
	"time"
)

const (
	_defaultWaitTime = 5 * time.Second
	_defaultAttempts = 10
	_defaultTimeOut  = 2 * time.Second
)

type Message struct {
	Queue         string
	Priority      uint8
	ContentType   string
	Body          []byte
	ReplyTo       string
	CorrelationID string
}

type pendingCall struct {
	done   chan struct{}
	status string
	body   []byte
}

type Client struct {
	conn           *adapter.RmqConnection
	serverExchange string
	error          chan error
	stop           chan struct{}

	rw      sync.RWMutex
	calls   map[string]*pendingCall
	timeout time.Duration
}

func New(url, serverExchange, clientExchange string, opts ...Option) (*Client, error) {
	cfg := adapter.RmqConfig{
		URL:      url,
		WaitTime: _defaultWaitTime,
		Attempts: _defaultAttempts,
	}

	c := &Client{
		conn:           adapter.New(clientExchange, cfg),
		serverExchange: serverExchange,
		error:          make(chan error),
		stop:           make(chan struct{}),
		calls:          make(map[string]*pendingCall),
		timeout:        _defaultTimeOut,
	}

	for _, opt := range opts {
		opt(c)
	}

	err := c.conn.AttemptsConnect()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RabbitMQ client: %w", err)
	}

	go c.consumer()
	return c, nil
}

func (c *Client) publish(corrID, handler string, request interface{}) error {
	var (
		requestBody []byte
		err         error
	)

	if request != nil {
		requestBody, err = json.Marshal(request)
		if err != nil {
			return err
		}
	}

	err = c.conn.Channel.Publish(c.serverExchange, "", false, false, amqp.Publishing{
		ContentType:   "application/json",
		CorrelationId: corrID,
		ReplyTo:       c.conn.ConsumerExchange,
		Type:          handler,
		Body:          requestBody,
	})

	if err != nil {
		return fmt.Errorf("Failed to create publish channel: %w", err)
	}
	return nil
}

func (c *Client) RemoteCall(handler string, request, response interface{}) error {
	select {
	case <-c.stop:
		time.Sleep(c.timeout)
		select {
		case <-c.stop:
			return errors.New("Client is stopped")
		default:
		}
	default:
	}

	corrID := uuid.New().String()
	err := c.publish(corrID, handler, request)
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %w", err)
	}

	call := &pendingCall{
		done: make(chan struct{}),
	}

	c.addCall(corrID, call)
	defer c.deleteCall(corrID)

	select {
	case <-time.After(c.timeout):
		return errors.New("Timeout")
	case <-call.done:
	}

	if call.status == constants.Success {
		err = json.Unmarshal(call.body, &response)
		if err != nil {
			return fmt.Errorf("Failed to unmarshal response: %w", err)
		}
		return nil
	}

	if call.status == constants.ErrorBadRequest.Error() {
		return constants.ErrorBadRequest
	}

	if call.status == constants.ErrorInternalServer.Error() {
		return constants.ErrorInternalServer
	}

	return nil
}

func (c *Client) consumer() {
	for {
		select {
		case <-c.stop:
			return
		case d, opened := <-c.conn.Delivery:
			if !opened {
				c.reconnect()
				return
			}

			_ = d.Ack(false)
			c.getCall(&d)
		}
	}
}

func (c *Client) reconnect() {
	close(c.stop)

	err := c.conn.AttemptsConnect()
	if err != nil {
		c.error <- err
		close(c.error)
		return
	}
	c.stop = make(chan struct{})
	go c.consumer()
}

func (c *Client) getCall(d *amqp.Delivery) {
	c.rw.RLock()
	call, ok := c.calls[d.CorrelationId]
	c.rw.RUnlock()

	if !ok {
		return
	}

	call.status = d.Type
	call.body = d.Body
	close(call.done)
}

func (c *Client) addCall(corrID string, call *pendingCall) {
	c.rw.Lock()
	c.calls[corrID] = call
	c.rw.Unlock()
}

func (c *Client) deleteCall(corrID string) {
	c.rw.Lock()
	delete(c.calls, corrID)
	c.rw.Unlock()
}

func (c *Client) Notify() <-chan error {
	return c.error
}

func (c *Client) Shutdown() error {
	select {
	case <-c.error:
		return nil
	default:
	}

	close(c.stop)
	time.Sleep(c.timeout)

	err := c.conn.Connection.Close()
	if err != nil {
		return fmt.Errorf("Failed to close connection: %w", err)
	}

	return nil
}
