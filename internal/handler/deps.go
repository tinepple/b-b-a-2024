package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"context"
)

type iStorage interface {
	CreateHouse(ctx context.Context, house storage.House) (storage.House, error)
	CreateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error)
	UpdateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error)
	CreateUser(ctx context.Context, user storage.User) (string, error)
	GetUserByID(ctx context.Context, userID string) (storage.User, error)
	GetUserByEmail(ctx context.Context, email string) (storage.User, error)
	GetFlatsByHouseID(ctx context.Context, houseID int64, status string) ([]storage.Flat, error)
	CreateHouseUserSubscription(ctx context.Context, houseID int64, userID string) error
	GetFlatByID(ctx context.Context, flatID int64) (storage.Flat, error)
}

type authService interface {
	GenerateJWT(userRole string, userID string) (string, error)
	ValidateModeratorRoleJWT(jwtToken string) error
	ValidateClientRoleJWT(jwtToken string) error
	GetUserID(jwtToken string) (string, error)
}

type kafkaService interface {
	Produce(houseID int64) error
}

type logger interface {
	Errorf(format string, args ...interface{})
}
