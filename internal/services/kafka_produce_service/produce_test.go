package kafka_produce_service

import (
	"backend-bootcamp-assignment-2024/internal/kafka_messages"
	"backend-bootcamp-assignment-2024/internal/services/kafka_produce_service/mocks"
	"encoding/json"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestService_Produce(t *testing.T) {
	ctrl := gomock.NewController(t)

	type serviceFields struct {
		producer producer
	}
	type args struct {
		houseID int64
	}

	tests := []struct {
		name          string
		serviceFields serviceFields
		args          args
		expectedErr   error
	}{
		{
			name: "Ошибка при отправке сообщения в кафку",
			serviceFields: serviceFields{
				producer: func(t *testing.T) producer {
					m := mocks.NewMockproducer(ctrl)

					bytes, _ := json.Marshal(kafka_messages.Message{
						HouseID: 2,
					})

					m.EXPECT().SendMessage(&sarama.ProducerMessage{
						Topic: "ping",
						Value: sarama.ByteEncoder(bytes),
					}).Return(int32(0), int64(0), errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 2,
			},
			expectedErr: errors.New("internal error"),
		},
		{
			name: "Успешная отправка сообщения в кафку",
			serviceFields: serviceFields{
				producer: func(t *testing.T) producer {
					m := mocks.NewMockproducer(ctrl)

					bytes, _ := json.Marshal(kafka_messages.Message{
						HouseID: 2,
					})

					m.EXPECT().SendMessage(&sarama.ProducerMessage{
						Topic: "ping",
						Value: sarama.ByteEncoder(bytes),
					}).Return(int32(0), int64(0), nil)

					return m
				}(t),
			},
			args: args{
				houseID: 2,
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := New(tt.serviceFields.producer)
			err := service.Produce(tt.args.houseID)
			if tt.expectedErr != nil {
				assert.Equal(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
