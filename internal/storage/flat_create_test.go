//go:build integration

package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_CreateFlat(t *testing.T) {
	testCases := []struct {
		name           string
		initDbQueries  []string
		args           Flat
		expectedResult Flat
		expectedErr    error
	}{
		{
			name: "Успешно создана квартира",
			initDbQueries: []string{`
					insert into houses (
						id,
						address,
						year
					) values (3, 'address', 2024)
				`,
			},
			args: Flat{
				HouseID:    3,
				Price:      123,
				RoomsCount: 4,
			},
			expectedResult: Flat{
				ID:         1,
				HouseID:    3,
				Status:     "created",
				Price:      123,
				RoomsCount: 4,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareDB(tc.initDbQueries...); err != nil {
				t.Fatal(err)
			}

			res, err := storageRepo.CreateFlat(context.Background(), tc.args)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedResult, res)
		})
	}
}
