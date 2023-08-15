// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: flight.sql

package models

import (
	"context"
)

const createFlight = `-- name: CreateFlight :one
INSERT INTO flights (
        flight_farm_id,
        flight_duration,
        flight_pilot,
        flight_acreage
    )
VALUES ($1, $2, $3, $4)
RETURNING id, flight_date, flight_farm_id, flight_duration, flight_pilot, flight_acreage
`

type CreateFlightParams struct {
	FlightFarmID   string
	FlightDuration string
	FlightPilot    string
	FlightAcreage  string
}

func (q *Queries) CreateFlight(ctx context.Context, arg CreateFlightParams) (Flight, error) {
	row := q.db.QueryRowContext(ctx, createFlight,
		arg.FlightFarmID,
		arg.FlightDuration,
		arg.FlightPilot,
		arg.FlightAcreage,
	)
	var i Flight
	err := row.Scan(
		&i.ID,
		&i.FlightDate,
		&i.FlightFarmID,
		&i.FlightDuration,
		&i.FlightPilot,
		&i.FlightAcreage,
	)
	return i, err
}
