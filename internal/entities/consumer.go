package entities

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"simple-go-nsq-app/internal/configs"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	consumer *nsq.Consumer
}

func NewConsumer(conf *nsq.Config, topic string, channel string) *Consumer {
	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.Fatal(err)
	}

	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Consume(handler *MessageHandler) {
	var nsqConf configs.NsqConfig = configs.NewNsqConfig()
	c.consumer.AddHandler(handler)

	if err := c.consumer.ConnectToNSQD(fmt.Sprintf("%s:%s", nsqConf.Host, nsqConf.NsqdPort)); err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	c.consumer.Stop()
}
