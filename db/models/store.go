package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
)

type Transaction struct {
	*Queries
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{
		db:      db,
		Queries: New(db),
	}
}

func (transaction *Transaction) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := transaction.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rbErr: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}


type FlightTxResult struct {
	Flight   Flight
	OldPilot Pilot
	NewPilot Pilot
}

func (transaction *Transaction) FlightTx(ctx context.Context, arg CreateFlightParams) (FlightTxResult, error) {
	var result FlightTxResult

	err := transaction.execTX(ctx, func(q *Queries) error {
		var err error
		result.Flight, err = q.CreateFlight(ctx, CreateFlightParams{
			FlightFarmID:          arg.FlightFarmID,
			FlightFarmLocation:    arg.FlightFarmLocation,
			FlightFarmGeolocation: arg.FlightFarmGeolocation,
			FlightDuration:        arg.FlightDuration,
			FlightPilot:           arg.FlightPilot,
			FlightAcreage:         arg.FlightAcreage,
		})
		if err != nil {
			return err
		}
		result.OldPilot, err = q.GetPilotForUpdate(ctx, arg.FlightPilot)

		if err != nil {
			return err
		}
		pilotFlightHours, err := strconv.ParseFloat(result.OldPilot.PilotFlightHours, 32)

		if err != nil {
			return err
		}
		pilotCoveredAcreage, err := strconv.ParseFloat(result.OldPilot.PilotCoveredAcreage, 32)

		if err != nil {
			return err
		}
		FlightHours, err := strconv.ParseFloat(arg.FlightDuration, 32)
		if err != nil {
			return err
		}
		FlightAcreage, err := strconv.ParseFloat(arg.FlightAcreage, 32)
		if err != nil {
			return err
		}

		result.NewPilot, err = q.UpdatePilot(ctx, UpdatePilotParams{
			PilotID:             arg.FlightPilot,
			PilotFlightHours:    fmt.Sprintf("%f", FlightHours+pilotFlightHours),
			PilotCoveredAcreage: fmt.Sprintf("%f", FlightAcreage+pilotCoveredAcreage),
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
