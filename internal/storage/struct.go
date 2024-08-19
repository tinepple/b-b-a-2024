package storage

import (
	"database/sql"
	"time"
)

type House struct {
	ID        int64          `db:"id"`
	Address   string         `db:"address"`
	Year      int64          `db:"year"`
	Developer sql.NullString `db:"developer"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type Flat struct {
	ID          int64          `db:"id"`
	HouseID     int64          `db:"house_id"`
	Status      string         `db:"status"`
	Number      sql.NullInt64  `db:"number"` // про поле написано в описании, но нигде нет в контрактах
	Price       int64          `db:"price"`
	RoomsCount  int64          `db:"rooms_count"`
	ModeratorID sql.NullString `db:"moderator_id"`
}

type User struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
