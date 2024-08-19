// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/deps.go

// Package mocks is a generated GoMock package.
package mocks

import (
	storage "backend-bootcamp-assignment-2024/internal/storage"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockiStorage is a mock of iStorage interface.
type MockiStorage struct {
	ctrl     *gomock.Controller
	recorder *MockiStorageMockRecorder
}

// MockiStorageMockRecorder is the mock recorder for MockiStorage.
type MockiStorageMockRecorder struct {
	mock *MockiStorage
}

// NewMockiStorage creates a new mock instance.
func NewMockiStorage(ctrl *gomock.Controller) *MockiStorage {
	mock := &MockiStorage{ctrl: ctrl}
	mock.recorder = &MockiStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiStorage) EXPECT() *MockiStorageMockRecorder {
	return m.recorder
}

// CreateFlat mocks base method.
func (m *MockiStorage) CreateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlat", ctx, flat)
	ret0, _ := ret[0].(storage.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlat indicates an expected call of CreateFlat.
func (mr *MockiStorageMockRecorder) CreateFlat(ctx, flat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlat", reflect.TypeOf((*MockiStorage)(nil).CreateFlat), ctx, flat)
}

// CreateHouse mocks base method.
func (m *MockiStorage) CreateHouse(ctx context.Context, house storage.House) (storage.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHouse", ctx, house)
	ret0, _ := ret[0].(storage.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHouse indicates an expected call of CreateHouse.
func (mr *MockiStorageMockRecorder) CreateHouse(ctx, house interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHouse", reflect.TypeOf((*MockiStorage)(nil).CreateHouse), ctx, house)
}

// CreateHouseUserSubscription mocks base method.
func (m *MockiStorage) CreateHouseUserSubscription(ctx context.Context, houseID int64, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHouseUserSubscription", ctx, houseID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHouseUserSubscription indicates an expected call of CreateHouseUserSubscription.
func (mr *MockiStorageMockRecorder) CreateHouseUserSubscription(ctx, houseID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHouseUserSubscription", reflect.TypeOf((*MockiStorage)(nil).CreateHouseUserSubscription), ctx, houseID, userID)
}

// CreateUser mocks base method.
func (m *MockiStorage) CreateUser(ctx context.Context, user storage.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockiStorageMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockiStorage)(nil).CreateUser), ctx, user)
}

// GetFlatByID mocks base method.
func (m *MockiStorage) GetFlatByID(ctx context.Context, flatID int64) (storage.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlatByID", ctx, flatID)
	ret0, _ := ret[0].(storage.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlatByID indicates an expected call of GetFlatByID.
func (mr *MockiStorageMockRecorder) GetFlatByID(ctx, flatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlatByID", reflect.TypeOf((*MockiStorage)(nil).GetFlatByID), ctx, flatID)
}

// GetFlatsByHouseID mocks base method.
func (m *MockiStorage) GetFlatsByHouseID(ctx context.Context, houseID int64, status string) ([]storage.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlatsByHouseID", ctx, houseID, status)
	ret0, _ := ret[0].([]storage.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlatsByHouseID indicates an expected call of GetFlatsByHouseID.
func (mr *MockiStorageMockRecorder) GetFlatsByHouseID(ctx, houseID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlatsByHouseID", reflect.TypeOf((*MockiStorage)(nil).GetFlatsByHouseID), ctx, houseID, status)
}

// GetUserByEmail mocks base method.
func (m *MockiStorage) GetUserByEmail(ctx context.Context, email string) (storage.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(storage.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockiStorageMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockiStorage)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockiStorage) GetUserByID(ctx context.Context, userID string) (storage.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(storage.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockiStorageMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockiStorage)(nil).GetUserByID), ctx, userID)
}

// UpdateFlat mocks base method.
func (m *MockiStorage) UpdateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlat", ctx, flat)
	ret0, _ := ret[0].(storage.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlat indicates an expected call of UpdateFlat.
func (mr *MockiStorageMockRecorder) UpdateFlat(ctx, flat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlat", reflect.TypeOf((*MockiStorage)(nil).UpdateFlat), ctx, flat)
}

// MockauthService is a mock of authService interface.
type MockauthService struct {
	ctrl     *gomock.Controller
	recorder *MockauthServiceMockRecorder
}

// MockauthServiceMockRecorder is the mock recorder for MockauthService.
type MockauthServiceMockRecorder struct {
	mock *MockauthService
}

// NewMockauthService creates a new mock instance.
func NewMockauthService(ctrl *gomock.Controller) *MockauthService {
	mock := &MockauthService{ctrl: ctrl}
	mock.recorder = &MockauthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockauthService) EXPECT() *MockauthServiceMockRecorder {
	return m.recorder
}

// GenerateJWT mocks base method.
func (m *MockauthService) GenerateJWT(userRole string, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWT", userRole, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWT indicates an expected call of GenerateJWT.
func (mr *MockauthServiceMockRecorder) GenerateJWT(userRole, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWT", reflect.TypeOf((*MockauthService)(nil).GenerateJWT), userRole, userID)
}

// GetUserID mocks base method.
func (m *MockauthService) GetUserID(jwtToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", jwtToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockauthServiceMockRecorder) GetUserID(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockauthService)(nil).GetUserID), jwtToken)
}

// ValidateClientRoleJWT mocks base method.
func (m *MockauthService) ValidateClientRoleJWT(jwtToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateClientRoleJWT", jwtToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateClientRoleJWT indicates an expected call of ValidateClientRoleJWT.
func (mr *MockauthServiceMockRecorder) ValidateClientRoleJWT(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateClientRoleJWT", reflect.TypeOf((*MockauthService)(nil).ValidateClientRoleJWT), jwtToken)
}

// ValidateModeratorRoleJWT mocks base method.
func (m *MockauthService) ValidateModeratorRoleJWT(jwtToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateModeratorRoleJWT", jwtToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateModeratorRoleJWT indicates an expected call of ValidateModeratorRoleJWT.
func (mr *MockauthServiceMockRecorder) ValidateModeratorRoleJWT(jwtToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateModeratorRoleJWT", reflect.TypeOf((*MockauthService)(nil).ValidateModeratorRoleJWT), jwtToken)
}

// MockkafkaService is a mock of kafkaService interface.
type MockkafkaService struct {
	ctrl     *gomock.Controller
	recorder *MockkafkaServiceMockRecorder
}

// MockkafkaServiceMockRecorder is the mock recorder for MockkafkaService.
type MockkafkaServiceMockRecorder struct {
	mock *MockkafkaService
}

// NewMockkafkaService creates a new mock instance.
func NewMockkafkaService(ctrl *gomock.Controller) *MockkafkaService {
	mock := &MockkafkaService{ctrl: ctrl}
	mock.recorder = &MockkafkaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockkafkaService) EXPECT() *MockkafkaServiceMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockkafkaService) Produce(houseID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", houseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockkafkaServiceMockRecorder) Produce(houseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockkafkaService)(nil).Produce), houseID)
}

// Mocklogger is a mock of logger interface.
type Mocklogger struct {
	ctrl     *gomock.Controller
	recorder *MockloggerMockRecorder
}

// MockloggerMockRecorder is the mock recorder for Mocklogger.
type MockloggerMockRecorder struct {
	mock *Mocklogger
}

// NewMocklogger creates a new mock instance.
func NewMocklogger(ctrl *gomock.Controller) *Mocklogger {
	mock := &Mocklogger{ctrl: ctrl}
	mock.recorder = &MockloggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocklogger) EXPECT() *MockloggerMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *Mocklogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockloggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*Mocklogger)(nil).Errorf), varargs...)
}
