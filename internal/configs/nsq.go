package configs

import (
	"simple-go-nsq-app/internal/utils"
)

type NsqConfig struct {
	Host string
	NsqdPort string
	NsqlookupdPort string
}

func NewNsqConfig() NsqConfig {
	return NsqConfig{
		Host: utils.LookupEnv("NSQ_HOST", "127.0.0.1"),
		NsqdPort: utils.LookupEnv("NSQD_PORT", "4150"),
		NsqlookupdPort: utils.LookupEnv("NSQLOOKUPD_PORT", "4161"),
	}
}