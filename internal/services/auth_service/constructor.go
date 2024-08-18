package auth_service

import (
	"os"
)

type service struct {
	privateKey []byte
}

type Service interface {
	GenerateJWT(userRole string, userID int64) (string, error)
	ValidateModeratorRoleJWT(jwtToken string) error
	ValidateClientRoleJWT(jwtToken string) error
	GetUserID(jwtToken string) (int64, error)
}

func New() Service {
	return &service{privateKey: []byte(os.Getenv("JWT_PRIVATE_KEY"))}
}
