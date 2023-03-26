package example

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/streadway/amqp"
// )

// type Offer struct {
// 	Type string `json:"type"`
// 	SDP  string `json:"sdp"`
// 	From string `json:"from"`
// 	To   string `json:"to"`
// }

// type IceCandidate struct {
// 	Type           string `json:"type"`
// 	Candidate      string `json:"candidate"`
// 	SDPMid         string `json:"sdpMid"`
// 	SDPMLineIndex  int    `json:"sdpMLineIndex"`
// 	From           string `json:"from"`
// 	To             string `json:"to"`
// }

// func main() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a channel: %v", err)
// 	}
// 	defer ch.Close()

// 	exchangeName := "webrtc-exchange"
// 	queueName := "webrtc-queue"

// 	err = ch.ExchangeDeclare(exchangeName, "fanout", false, false, false, false, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to declare exchange: %v", err)
// 	}

// 	q, err := ch.QueueDeclare(queueName, false, true, true, false, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to declare queue: %v", err)
// 	}

// 	err = ch.QueueBind(q.Name, "", exchangeName, false, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to bind queue: %v", err)
// 	}

// 	offers := make(map[string]*webrtc.PeerConnection)
// 	iceCandidates := make(map[string][]IceCandidate)

// 	// Handle incoming offers and ice candidates
// 	go func() {
// 		for {
// 			msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
// 			if err != nil {
// 				log.Fatalf("Failed to consume messages: %v", err)
// 			}

// 			for msg := range msgs {
// 				switch msg.ContentType {
// 				case "offer":
// 					offer := &Offer{}
// 					err := json.Unmarshal(msg.Body, offer)
// 					if err != nil {
// 						log.Printf("Failed to unmarshal offer: %v", err)
// 						continue
// 					}

// 					pc, err := webrtc.NewPeerConnection(webrtc.Configuration{})
// 					if err != nil {
// 						log.Printf("Failed to create peer connection: %v", err)
// 						continue
// 					}

// 					offers[offer.From] = pc

// 					// Add video tracks to the peer connection
// 					_, err = pc.AddTransceiver(webrtc.RTPCodecTypeVideo)
// 					if err != nil {
// 						log.Printf("Failed to add video transceiver: %v", err)
// 						continue
// 					}

// 					// Set the remote description from the offer
// 					err = pc.SetRemoteDescription(webrtc.SessionDescription{
// 						Type: webrtc.SDPTypeOffer,
// 						SDP:  offer.SDP,
// 					})
// 					if err != nil {
// 						log.Printf("Failed to set remote description: %v", err)
// 						continue
// 					}

// 					// Create an answer and
// 					answer, err := pc.CreateAnswer(nil)
// 					if err != nil {
// 						log.Printf("Failed to create answer: %v", err)
// 						continue
// 					}

// 					// Set the local description of the peer connection
// 					err = pc.SetLocalDescription(answer)
// 					if err != nil {
// 						log.Printf("Failed to set local description: %v", err)
// 						continue
// 					}

// 					// Publish the answer
// 					answerMsg, err := json.Marshal(Offer{
// 						Type: "answer",
// 						SDP:  answer.SDP,
// 						From: offer.To,
// 						To:   offer.From,
// 					})
// 					if err != nil {
// 						log.Printf("Failed to marshal answer: %v", err)
// 						continue
// 					}

// 					err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
// 						ContentType: "answer",
// 						Body:        answerMsg,
// 					})
// 					if err != nil {
// 						log.Printf("Failed to publish answer: %v", err)
// 						continue
// 					}

// 				case "icecandidate":
// 					iceCandidate := &IceCandidate{}
// 					err := json.Unmarshal(msg.Body, iceCandidate)
// 					if err != nil {
// 						log.Printf("Failed to unmarshal ice candidate: %v", err)
// 						continue
// 					}

// 					if pc, ok := offers[iceCandidate.To]; ok {
// 						err = pc.AddICECandidate(webrtc.ICECandidateInit{
// 							Candidate:     iceCandidate.Candidate,
// 							SDPMid:        iceCandidate.SDPMid,
// 							SDPMLineIndex: iceCandidate.SDPMLineIndex,
// 						})
// 						if err != nil {
// 							log.Printf("Failed to add ICE candidate: %v", err)
// 							continue
// 						}
// 					} else {
// 						iceCandidates[iceCandidate.To] = append(iceCandidates[iceCandidate.To], *iceCandidate)
// 					}
// 				}
// 			}
// 		}
// 	}()

// 	// Handle outgoing offers and ice candidates
// 	for {
// 		fmt.Println("Enter command (offer <to> | icecandidate <to> <candidate> <sdpMid> <sdpMLineIndex>):")

// 		var cmd, to, candidate, sdpMid string
// 		var sdpMLineIndex int
// 		fmt.Scan(&cmd)

// 		switch cmd {
// 		case "offer":
// 			fmt.Scan(&to)

// 			pc, err := webrtc.NewPeerConnection(webrtc.Configuration{})
// 			if err != nil {
// 				log.Printf("Failed to create peer connection: %v", err)
// 				continue
// 			}

// 			offers[to] = pc

// 			// Add video tracks to the peer connection
// 			_, err = pc.AddTransceiver(webrtc.RTPCodecTypeVideo)
// 			if err != nil {
// 				log.Printf("Failed to add video transceiver: %v", err)
// 				continue
// 			}

// 			// Create an offer
// 			offer, err := pc.CreateOffer(nil)
// 			if err != nil {
// 				log.Printf("Failed to create offer: %v", err)
// 				continue
// 			}

// 			// Set the local description of the peer connection
// 			err = pc.SetLocalDescription(offer)
// 			if err != nil {
// 				log.Printf("Failed to set local description: %v", err)
// 				continue
// 			}

// 			// Publish the offer
// 			offerMsg, err := json.Marshal(Offer{
// 				Type: "offer",
// 				SDP:  offer.SDP,
// 				From: "me",
// 				To:   to,
// 			})
// 			if err != nil {
// 				log.Printf
// 				("Failed to marshal offer: %v", err)
// 				continue
// 			}

// 			err = ch.Publish(exchangeName, "", false, false, amqp.Publishing{
// 				ContentType: "offer",
// 				Body:        offerMsg,
// 			})
// 			if err != nil {
// 				log.Printf("Failed to publish offer: %v", err)
// 				continue
// 			}

// 		case "icecandidate":
// 			fmt.Scan(&to, &candidate, &sdpMid, &sdpMLineIndex)

// 			if pc, ok := offers[to]; ok {
// 				err := pc.AddICECandidate(webrtc.ICECandidateInit{
// 					Candidate:     candidate,
// 					SDPMid:        sdpMid,
// 					SDPMLineIndex: sdpMLineIndex,
// 				})
// 				if err != nil {
// 					log.Printf("Failed to add ICE candidate: %v", err)
// 					continue
// 				}
// 			} else {
// 				iceCandidates[to] = append(iceCandidates[to], IceCandidate{
// 					Candidate:     candidate,
// 					SDPMid:        sdpMid,
// 					SDPMLineIndex: sdpMLineIndex,
// 					To:            to,
// 				})
// 			}
// 		}
// 	}
// }

// // signal
// if *roomName == "" {
// 	log.Fatalf("Must specify room name")
// }

// connectStr := os.Getenv("RABBITMQ_URL")
// if connectStr == "" {
// 	connectStr = "amqp://guest:guest@localhost:5672/"
// }

// conn, err := amqp.Dial(connectStr)
// if err != nil {
// 	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
// }
// defer conn.Close()

// ch, err := conn.Channel()
// if err != nil {
// 	log.Fatalf("Failed to open RabbitMQ channel: %v", err)
// }
// defer ch.Close()

// err = ch.ExchangeDeclare(
// 	exchangeName,
// 	amqp.ExchangeTopic,
// 	true,
// 	false,
// 	false,
// 	false,
// 	nil,
// )
// if err != nil {
// 	log.Fatalf("Failed to declare RabbitMQ exchange: %v", err)
// }

// err = ch.QueueBind(
// 	queueName,
// 	"*",
// 	exchangeName,
// 	false,
// 	nil,
// )
// if err != nil {
// 	log.Fatalf("Failed to bind RabbitMQ queue: %v", err)
// }

// log.Printf("Listening for messages in room %q...", *roomName)
// handleMessages(ch)
