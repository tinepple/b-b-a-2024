package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetFlatsByHouseID(ctx context.Context, houseID int64, status string) ([]Flat, error) {
	builder := sq.Select(
		"id",
		"house_id",
		"status",
		"price",
		"rooms_count",
	).From(flatsTableName).
		Where(sq.Eq{"house_id": houseID}).
		PlaceholderFormat(sq.Dollar)

	if status != "" {
		builder = builder.Where(sq.Eq{"status": status})
	}

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var dest []Flat

	err = s.db.SelectContext(ctx, &dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
