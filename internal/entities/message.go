package entities

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type Message struct {
	Name      string
	Content   string
	Timestamp time.Time
}

type MessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	// Unmarshal the Message
	var req Message
	if err := json.Unmarshal(m.Body, &req); err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)
		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return err
	}

	//Print the Message
	log.Println("Message Received")
	log.Println("Name : ", req.Name)
	log.Println("Content : ", req.Content)
	log.Println("Timestamp : ", req.Timestamp)
	log.Println("")

	// Will automatically set the message as finish
	return nil
}
