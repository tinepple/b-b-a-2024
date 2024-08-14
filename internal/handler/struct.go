package handler

import "time"

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
