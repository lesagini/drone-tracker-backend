package models

import (
	"context"
	"drones/utils"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestFlightTx(t *testing.T) {
	store := NewTransaction(testDb)
	d = utils.GetDetails()
	farm := createRandomFarm(t)
	createRandomOperator(t)
	pilot := createRandomPilot(t)
	fd := gofakeit.Float32Range(1.0, 8.0)
	fa := fd * gofakeit.Float32Range(0.8, 1.2) * 2000

	n := 5

	errs := make(chan error)
	results := make(chan FlightTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.FlightTx(context.Background(), CreateFlightParams{
				FlightFarmID:          farm.FarmCode,
				FlightFarmLocation:    farm.FarmLocation,
				FlightFarmGeolocation: farm.FarmGeolocation,
				FlightDuration:        fmt.Sprintf("%f", fd),
				FlightPilot:           pilot.PilotID,
				FlightAcreage:         fmt.Sprintf("%f", fa),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)

		newPilot := result.NewPilot

		require.NotEmpty(t, newPilot)
	}

}
