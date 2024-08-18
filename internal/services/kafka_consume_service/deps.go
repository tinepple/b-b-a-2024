package kafka_consume_service

import "context"

type senderService interface {
	SendEmail(ctx context.Context, recipient string, message string) error
}

type iStorage interface {
	GetHouseUserSubscriptionsEmails(ctx context.Context, houseID int64) ([]string, error)
}

type logger interface {
	Errorf(format string, args ...interface{})
}
