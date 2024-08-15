package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateFlat(ctx context.Context, flat Flat) (Flat, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return Flat{}, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query, params, err := sq.Insert(flatsTableName).
		Columns(
			"house_id",
			"price",
			"rooms_count",
		).
		Values(flat.HouseID, flat.Price, flat.RoomsCount).
		Suffix("returning id, house_id, status, number, price, rooms_count").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		fmt.Println(31)
		return Flat{}, err
	}

	var dest Flat

	err = tx.QueryRowContext(ctx, tx.Rebind(query), params...).Scan(
		&dest.ID,
		&dest.HouseID,
		&dest.Status,
		&dest.Number,
		&dest.Price,
		&dest.RoomsCount,
	)
	if err != nil {
		fmt.Println(46)
		return Flat{}, err
	}

	query, params, err = sq.Update(housesTableName).
		Set(
			"updated_at", "now()",
		).
		Where(sq.Eq{"id": dest.HouseID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		fmt.Println(58)
		return Flat{}, err
	}

	_, err = tx.ExecContext(ctx, tx.Rebind(query), params...)
	if err != nil {
		fmt.Println(65)
		return Flat{}, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(70)
		return Flat{}, err
	}

	return dest, nil
}
