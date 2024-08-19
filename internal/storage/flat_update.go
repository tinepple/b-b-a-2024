package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) UpdateFlat(ctx context.Context, flat Flat) (Flat, error) {
	query, params, err := sq.Update(flatsTableName).
		SetMap(map[string]interface{}{
			"price":        flat.Price,
			"rooms_count":  flat.RoomsCount,
			"status":       flat.Status,
			"moderator_id": flat.ModeratorID.String,
		}).
		Where(sq.Eq{"id": flat.ID}).
		Suffix("returning id, house_id, status, number, price, rooms_count").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return Flat{}, err
	}

	var dest Flat

	err = s.db.QueryRowContext(ctx, s.db.Rebind(query), params...).Scan(
		&dest.ID,
		&dest.HouseID,
		&dest.Status,
		&dest.Number,
		&dest.Price,
		&dest.RoomsCount,
	)
	if err != nil {
		return Flat{}, err
	}

	return dest, nil
}
