package storage

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetHouseUserSubscriptionsEmails(ctx context.Context, houseID int64) ([]string, error) {
	query, params, err := sq.Select(
		"u.email",
	).From(fmt.Sprintf("%s hus", houseUserSubscriptionsTableName)).
		InnerJoin(fmt.Sprintf("%s u on u.id = hus.user_id", usersTableName)).
		Where(sq.Eq{"hus.house_id": houseID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var dest []string

	err = s.db.SelectContext(ctx, &dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
