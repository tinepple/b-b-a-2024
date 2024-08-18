package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetFlatByID(ctx context.Context, flatID int64) (Flat, error) {
	query, params, err := sq.Select(
		"status",
		"moderator_id",
	).From(flatsTableName).
		Where(sq.Eq{"id": flatID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var dest Flat

	err = s.db.QueryRowContext(ctx, s.db.Rebind(query), params...).Scan(&dest.Status, &dest.ModeratorID)
	if err != nil {
		return Flat{}, err
	}

	return dest, nil
}
