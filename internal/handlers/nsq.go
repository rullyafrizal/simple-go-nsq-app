package handlers

import (
	"simple-go-nsq-app/internal/services"

	"github.com/nsqio/go-nsq"
)

type NsqHandler struct {
	NsqService services.NsqServiceContract
}

func NewNsqHandler() *NsqHandler {
	return &NsqHandler{
		NsqService: services.NewNsqService(*nsq.NewConfig()),
	}
}

func (h *NsqHandler) StartProducing() {
	h.NsqService.Publish()
}

func (h *NsqHandler) StartConsuming() {
	h.NsqService.Consume()
}
