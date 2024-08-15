package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateUser(ctx context.Context, user User) (int64, error) {
	query, params, err := sq.Insert(usersTableName).
		Columns(
			"email",
			"password",
			"role",
		).
		Values(user.Email, user.Password, user.Role).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	var dest int64

	err = s.db.QueryRowContext(ctx, s.db.Rebind(query), params...).Scan(
		&dest,
	)
	if err != nil {
		return 0, err
	}

	return dest, nil
}
