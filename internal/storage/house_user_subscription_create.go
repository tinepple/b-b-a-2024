package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateHouseUserSubscription(ctx context.Context, houseID int64, userID string) error {
	query, params, err := sq.Insert(houseUserSubscriptionsTableName).
		Columns(
			"house_id",
			"user_id",
		).
		Values(houseID, userID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, s.db.Rebind(query), params...)
	if err != nil {
		return err
	}

	return nil
}
