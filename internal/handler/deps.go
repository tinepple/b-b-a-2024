package handler

import (
	"backend-bootcamp-assignment-2024/internal/storage"
	"context"
)

type iStorage interface {
	CreateHouse(ctx context.Context, house storage.House) (storage.House, error)
	CreateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error)
	UpdateFlat(ctx context.Context, flat storage.Flat) (storage.Flat, error)
	GetFlatsByHouseID(ctx context.Context, houseID int64, status string) ([]storage.Flat, error)
	CreateUser(ctx context.Context, user storage.User) (int64, error)
	GetUserByID(ctx context.Context, userID int64) (storage.User, error)
	GetUserByEmail(ctx context.Context, email string) (storage.User, error)
}

type authService interface {
	GenerateJWT(userRole string) (string, error)
	ValidateModeratorRoleJWT(jwtToken string) error
	ValidateClientRoleJWT(jwtToken string) error
}
