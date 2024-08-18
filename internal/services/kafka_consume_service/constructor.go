package kafka_consume_service

import (
	"context"

	"github.com/IBM/sarama"
)

type service struct {
	consumer      sarama.PartitionConsumer
	senderService senderService
	storage       iStorage
	logger        logger
}

type Service interface {
	Consume(ctx context.Context) error
}

func New(consumer sarama.PartitionConsumer, senderService senderService, storage iStorage, logger logger) Service {
	return &service{
		consumer:      consumer,
		senderService: senderService,
		storage:       storage,
		logger:        logger,
	}
}
