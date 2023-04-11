// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: flight.sql

package models

import (
	"context"
)

const createFlight = `-- name: CreateFlight :one
INSERT INTO flights (
        flight_farm_location,
        flight_farm_id,
        flight_farm_geolocation,
        flight_duration,
        flight_pilot,
        flight_acreage
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, flight_date, flight_farm_location, flight_farm_id, flight_farm_geolocation, flight_duration, flight_pilot, flight_acreage
`

type CreateFlightParams struct {
	FlightFarmLocation    string
	FlightFarmID          string
	FlightFarmGeolocation string
	FlightDuration        string
	FlightPilot           string
	FlightAcreage         string
}

func (q *Queries) CreateFlight(ctx context.Context, arg CreateFlightParams) (Flight, error) {
	row := q.db.QueryRowContext(ctx, createFlight,
		arg.FlightFarmLocation,
		arg.FlightFarmID,
		arg.FlightFarmGeolocation,
		arg.FlightDuration,
		arg.FlightPilot,
		arg.FlightAcreage,
	)
	var i Flight
	err := row.Scan(
		&i.ID,
		&i.FlightDate,
		&i.FlightFarmLocation,
		&i.FlightFarmID,
		&i.FlightFarmGeolocation,
		&i.FlightDuration,
		&i.FlightPilot,
		&i.FlightAcreage,
	)
	return i, err
}
