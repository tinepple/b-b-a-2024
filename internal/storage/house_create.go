package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateHouse(ctx context.Context, house House) (House, error) {
	query, params, err := sq.Insert(housesTableName).
		Columns(
			"address",
			"year",
			"developer",
		).
		Values(
			house.Address,
			house.Year,
			house.Developer,
		).
		Suffix("returning id, adress,year, developer, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return House{}, err
	}

	var dest House

	err = s.db.QueryRow(query, params...).Scan(
		&dest.ID,
		&dest.Address,
		&dest.Year,
		&dest.Developer,
		&dest.CreatedAt,
		&dest.UpdatedAt,
	)
	if err != nil {
		return House{}, err
	}
	return dest, err
}
