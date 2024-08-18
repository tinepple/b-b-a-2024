package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/mocks"
	"backend-bootcamp-assignment-2024/internal/storage"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestHandler_FlatCreate(t *testing.T) {
	ctrl := gomock.NewController(t)

	type expectedResult struct {
		statusCode int
		body       string
	}
	type args struct {
		houseID int64
		price   int64
		rooms   int64
	}
	type handlerFields struct {
		storage      iStorage
		authService  authService
		kafkaService kafkaService
		logger       logger
	}
	tests := []struct {
		name           string
		handlerFields  handlerFields
		args           args
		expectedResult expectedResult
	}{
		{
			name: "Ошибка доступа при проверке токена",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)
					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(errors.New("internal error"))

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)

					m.EXPECT().Errorf("authService.ValidateClientRoleJWT error: %v", errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   2,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка валидации houseID",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)
					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)
					return m
				}(t),
			},
			args: args{
				houseID: 0,
				price:   2,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка валидации price",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)
					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)
					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   -10,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка валидации rooms",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)
					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)
					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   2,
				rooms:   0,
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка при создании квартиры",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().CreateFlat(gomock.Any(), storage.Flat{
						HouseID:    1,
						Price:      2,
						RoomsCount: 3,
					}).Return(storage.Flat{}, errors.New("internal error"))

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)

					m.EXPECT().Errorf("handler.FlatCreate,storage.CreateFlat error: %v", errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   2,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 500,
				body:       `{"code":500,"message":"error creating flat"}`,
			},
		},
		{
			name: "Ошибка при отправке сообщения в кафку",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().CreateFlat(gomock.Any(), storage.Flat{
						HouseID:    1,
						Price:      2,
						RoomsCount: 3,
					}).Return(storage.Flat{
						ID:         1,
						HouseID:    1,
						Price:      2,
						RoomsCount: 3,
						Status:     "created",
					}, nil)

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				kafkaService: func(t *testing.T) kafkaService {
					m := mocks.NewMockkafkaService(ctrl)

					m.EXPECT().Produce(int64(1)).Return(errors.New("internal error"))

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)

					m.EXPECT().Errorf("handler.FlatCreate,kafkaService.Produce error: %v", errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   2,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 500,
				body:       `{"code":500,"message":"error producing message"}`,
			},
		},
		{
			name: "Успешное создание квартиры",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().CreateFlat(gomock.Any(), storage.Flat{
						HouseID:    1,
						Price:      2,
						RoomsCount: 3,
					}).Return(storage.Flat{
						ID:         1,
						HouseID:    1,
						Price:      2,
						RoomsCount: 3,
						Status:     "created",
					}, nil)

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)

					return m
				}(t),
				kafkaService: func(t *testing.T) kafkaService {
					m := mocks.NewMockkafkaService(ctrl)

					m.EXPECT().Produce(int64(1)).Return(nil)

					return m
				}(t),
			},
			args: args{
				houseID: 1,
				price:   2,
				rooms:   3,
			},
			expectedResult: expectedResult{
				statusCode: 200,
				body:       `{"id":1,"house_id":1,"price":2,"rooms":3,"status":"created"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := New(
				tt.handlerFields.storage,
				tt.handlerFields.authService,
				tt.handlerFields.kafkaService,
				tt.handlerFields.logger,
			)

			w := httptest.NewRecorder()

			body, _ := json.Marshal(Flat{
				HouseID: tt.args.houseID,
				Price:   tt.args.price,
				Rooms:   tt.args.rooms,
			})
			req, _ := http.NewRequest("POST", "/flat/create", strings.NewReader(string(body)))
			req.Header.Set("Authorization", "Bearer someToken")
			handler.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedResult.statusCode, w.Code)
			assert.Equal(t, tt.expectedResult.body, w.Body.String())
		})
	}
}
