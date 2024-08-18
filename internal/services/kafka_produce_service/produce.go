package kafka_produce_service

import (
	"backend-bootcamp-assignment-2024/internal/kafka_messages"
	"encoding/json"

	"github.com/IBM/sarama"
)

func (s *service) Produce(houseID int64) error {
	message := kafka_messages.Message{
		HouseID: houseID,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "ping",
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
