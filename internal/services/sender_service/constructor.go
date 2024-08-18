package sender_service

import "context"

type sender struct{}

type Sender interface {
	SendEmail(ctx context.Context, recipient string, message string) error
}

func New() Sender {
	return &sender{}
}
