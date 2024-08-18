package handler

import (
	"backend-bootcamp-assignment-2024/internal/handler/mocks"
	"backend-bootcamp-assignment-2024/internal/storage"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestHandler_HouseGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	type expectedResult struct {
		statusCode int
		body       string
	}
	type args struct {
		houseID int64
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
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка при валидации houseID",
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

					m.EXPECT().Errorf("handler.HouseGet,invalid id: %s", "0")

					return m
				}(t),
			},
			args: args{
				houseID: 0,
			},
			expectedResult: expectedResult{
				statusCode: 400,
			},
		},
		{
			name: "Ошибка при получении квартир за клиента",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().GetFlatsByHouseID(gomock.Any(), int64(1), "approved").
						Return([]storage.Flat{
							{
								ID:         1,
								HouseID:    1,
								Price:      1000,
								RoomsCount: 3,
								Status:     "created",
							},
							{
								ID:         2,
								HouseID:    1,
								Price:      500,
								RoomsCount: 2,
								Status:     "approved",
							},
						}, errors.New("internal error"))

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)
					m.EXPECT().ValidateModeratorRoleJWT("someToken").Return(errors.New("internal error"))

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)

					m.EXPECT().Errorf("handler.HouseGet,storage.GetFlatsByHouseID error: %v", errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
			},
			expectedResult: expectedResult{
				statusCode: 500,
				body:       `{"code":500,"message":"error getting flats"}`,
			},
		},
		{
			name: "Ошибка при получении квартир за модератора",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().GetFlatsByHouseID(gomock.Any(), int64(1), "").
						Return([]storage.Flat{
							{
								ID:         1,
								HouseID:    1,
								Price:      1000,
								RoomsCount: 3,
								Status:     "created",
							},
							{
								ID:         2,
								HouseID:    1,
								Price:      500,
								RoomsCount: 2,
								Status:     "approved",
							},
						}, errors.New("internal error"))

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)
					m.EXPECT().ValidateModeratorRoleJWT("someToken").Return(nil)

					return m
				}(t),
				logger: func(t *testing.T) logger {
					m := mocks.NewMocklogger(ctrl)

					m.EXPECT().Errorf("handler.HouseGet,storage.GetFlatsByHouseID error: %v", errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
			},
			expectedResult: expectedResult{
				statusCode: 500,
				body:       `{"code":500,"message":"error getting flats"}`,
			},
		},
		{
			name: "Успешное получение квартир за модератора",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().GetFlatsByHouseID(gomock.Any(), int64(1), "").
						Return([]storage.Flat{
							{
								ID:         1,
								HouseID:    1,
								Price:      1000,
								RoomsCount: 3,
								Status:     "created",
							},
							{
								ID:         2,
								HouseID:    1,
								Price:      500,
								RoomsCount: 2,
								Status:     "approved",
							},
						}, nil)

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)
					m.EXPECT().ValidateModeratorRoleJWT("someToken").Return(nil)

					return m
				}(t),
			},
			args: args{
				houseID: 1,
			},
			expectedResult: expectedResult{
				statusCode: 200,
				body:       `{"flats":[{"id":1,"house_id":1,"price":1000,"rooms":3,"status":"created"},{"id":2,"house_id":1,"price":500,"rooms":2,"status":"approved"}]}`,
			},
		},
		{
			name: "Успешное получение квартир за клиента",
			handlerFields: handlerFields{
				storage: func(t *testing.T) iStorage {
					m := mocks.NewMockiStorage(ctrl)

					m.EXPECT().GetFlatsByHouseID(gomock.Any(), int64(1), "approved").
						Return([]storage.Flat{
							{
								ID:         1,
								HouseID:    1,
								Price:      1000,
								RoomsCount: 3,
								Status:     "created",
							},
							{
								ID:         2,
								HouseID:    1,
								Price:      500,
								RoomsCount: 2,
								Status:     "approved",
							},
						}, nil)

					return m
				}(t),
				authService: func(t *testing.T) authService {
					m := mocks.NewMockauthService(ctrl)

					m.EXPECT().ValidateClientRoleJWT("someToken").Return(nil)
					m.EXPECT().ValidateModeratorRoleJWT("someToken").Return(errors.New("internal error"))

					return m
				}(t),
			},
			args: args{
				houseID: 1,
			},
			expectedResult: expectedResult{
				statusCode: 200,
				body:       `{"flats":[{"id":1,"house_id":1,"price":1000,"rooms":3,"status":"created"},{"id":2,"house_id":1,"price":500,"rooms":2,"status":"approved"}]}`,
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
			req, _ := http.NewRequest("GET", fmt.Sprintf("/house/%d", tt.args.houseID), nil)
			req.Header.Set("Authorization", "Bearer someToken")
			handler.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedResult.statusCode, w.Code)
			assert.Equal(t, tt.expectedResult.body, w.Body.String())
		})
	}
}
