package services

import (
	"simple-go-nsq-app/internal/entities"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

type NsqServiceContract interface {
	Consume()
	Publish()
}

type NsqService struct {
	Consumer  *entities.Consumer
	Publisher *entities.Publisher
	Message   *entities.Message
}

func NewNsqService(conf nsq.Config) NsqServiceContract {
	return &NsqService{
		Consumer:  entities.NewConsumer(&conf, "test", "test_channel"),
		Publisher: entities.NewPublisher(&conf),
	}
}

func (s *NsqService) Consume() {
	var messageHandler *entities.MessageHandler = &entities.MessageHandler{}

	s.Consumer.Consume(messageHandler)
}

func (s *NsqService) Publish() {
	for i := 0; i < 100; i++ {
		s.Publisher.Publish("test", &entities.Message{
			Name:      "This is a test " + strconv.Itoa(i),
			Content:   "Content of a test number " + strconv.Itoa(i),
			Timestamp: time.Now(),
		})
	}
}
