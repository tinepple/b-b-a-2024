//go:build integration

package storage

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage_CreateHouse(t *testing.T) {
	testCases := []struct {
		name           string
		args           House
		expectedResult House
		expectedErr    error
	}{
		{
			name: "Успешно создан дом",
			args: House{
				Address: "some_address",
				Year:    2024,
				Developer: sql.NullString{
					String: "gk samolet",
					Valid:  true,
				},
			},
			expectedResult: House{
				ID:      1,
				Address: "some_address",
				Year:    2024,
				Developer: sql.NullString{
					String: "gk samolet",
					Valid:  true,
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			house, err := storageRepo.CreateHouse(context.Background(), tc.args)
			assert.Equal(t, tc.expectedErr, err)
			compareHouses(t, tc.expectedResult, house)
		})
	}
}

func compareHouses(t *testing.T, expected, actual House) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Address, actual.Address)
	assert.Equal(t, expected.Year, actual.Year)
	assert.Equal(t, expected.Developer, actual.Developer)
	assert.Equal(t, expected.CreatedAt.Truncate(time.Hour).UTC(), actual.CreatedAt.Truncate(time.Hour).UTC())
	assert.Equal(t, expected.UpdatedAt.Truncate(time.Hour).UTC(), actual.UpdatedAt.Truncate(time.Hour).UTC())
}
