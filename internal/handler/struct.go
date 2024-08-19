package handler

import "time"

var ValidUserTypes = map[string]struct{}{
	"client":    {},
	"moderator": {},
}

var ValidFlatStatuses = map[string]struct{}{
	"created":       {},
	"approved":      {},
	"declined":      {},
	"on moderation": {},
}

type HouseCreateRequest struct {
	Address   string `json:"address"`
	Year      int64  `json:"year"`
	Developer string `json:"developer"`
}

type HouseCreateResponse struct {
	ID        int64     `json:"id"`
	Address   string    `json:"address"`
	Year      int64     `json:"year"`
	Developer string    `json:"developer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type HouseGetResponse struct {
	Flats []Flat `json:"flats"`
}

type HouseSubscribeRequest struct {
	Email string `json:"email"`
}

type Flat struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int64  `json:"rooms"`
	Status  string `json:"status"`
}

type FlatCreateRequest struct {
	HouseID int64 `json:"house_id"`
	Price   int64 `json:"price"`
	Rooms   int64 `json:"rooms"`
}

type FlatCreateResponse struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int64  `json:"rooms"`
	Status  string `json:"status"`
}

type FlatUpdateRequest struct {
	ID     int64  `json:"id"`
	Price  int64  `json:"price"`
	Rooms  int64  `json:"rooms"`
	Status string `json:"status"`
}

type FlatUpdateResponse struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int64  `json:"rooms"`
	Status  string `json:"status"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type DummyLoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID   int64
	Role string
}

type InternalErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
