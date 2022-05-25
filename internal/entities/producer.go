package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"simple-go-nsq-app/internal/configs"

	"github.com/nsqio/go-nsq"
)

type Publisher struct {
	Producer *nsq.Producer
}

func NewPublisher(conf *nsq.Config) *Publisher {
	var nsqConf configs.NsqConfig = configs.NewNsqConfig()

	producer, err := nsq.NewProducer(
		fmt.Sprintf("%s:%s", nsqConf.Host, nsqConf.NsqdPort),
		conf,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &Publisher{
		Producer: producer,
	}
}

func (p *Publisher) Publish(topic string, msg *Message) {
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Producer.Publish(topic, payload); err != nil {
		log.Fatal(err)
	}
}
