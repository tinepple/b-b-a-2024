package kafka_produce_service

import (
	"github.com/IBM/sarama"
)

type producer interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type service struct {
	producer producer
}

type Service interface {
	Produce(houseID int64) error
}

func New(producer producer) Service {
	return &service{
		producer: producer,
	}
}
