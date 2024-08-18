package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetUser(ctx context.Context, userID int64) (User, error) {
	query, params, err := sq.Select(
		"id",
		"email",
		"password",
		"role",
	).From(usersTableName).
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return User{}, err
	}

	var dest User

	err = s.db.QueryRowContext(ctx, s.db.Rebind(query), params...).Scan(
		&dest.ID,
		&dest.Email,
		&dest.Password,
		&dest.Role,
	)
	if err != nil {
		return User{}, err
	}

	return dest, nil
}
