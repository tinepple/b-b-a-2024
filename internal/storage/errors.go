package storage

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("row not found")

func handleError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}
