package kafka_consume_service

import (
	"backend-bootcamp-assignment-2024/internal/kafka_messages"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *service) Consume(ctx context.Context) error {
	for {
		select {
		case msg, ok := <-s.consumer.Messages():
			if !ok {
				return errors.New("channel closed, exiting")
			}

			var receivedMessage kafka_messages.Message
			err := json.Unmarshal(msg.Value, &receivedMessage)
			if err != nil {
				s.logger.Errorf("Error unmarshaling JSON: %v", err)
				continue
			}

			emails, err := s.storage.GetHouseUserSubscriptionsEmails(ctx, receivedMessage.HouseID)
			if err != nil {
				s.logger.Errorf("storage.GetHouseUserSubscriptionsEmails, error: %v", err)
				continue
			}

			for _, email := range emails {
				err := s.senderService.SendEmail(ctx, email, fmt.Sprintf("в доме %d появились новые квартиры", receivedMessage.HouseID))
				if err != nil {
					s.logger.Errorf("senderService.SendEmail, error: %v", err)
					continue
				}
			}
		}
	}
}
