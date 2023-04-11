// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: farm.sql

package models

import (
	"context"
)

const createFarm = `-- name: CreateFarm :one
INSERT INTO farms (
        farm_code,
        farm_coordinates,
        farm_airspace,
        farm_location,
        farm_geolocation,
        farm_contact
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, farm_code, farm_coordinates, farm_airspace, farm_location, farm_geolocation, farm_creation_date, farm_contact
`

type CreateFarmParams struct {
	FarmCode        string
	FarmCoordinates string
	FarmAirspace    string
	FarmLocation    string
	FarmGeolocation string
	FarmContact     int64
}

func (q *Queries) CreateFarm(ctx context.Context, arg CreateFarmParams) (Farm, error) {
	row := q.db.QueryRowContext(ctx, createFarm,
		arg.FarmCode,
		arg.FarmCoordinates,
		arg.FarmAirspace,
		arg.FarmLocation,
		arg.FarmGeolocation,
		arg.FarmContact,
	)
	var i Farm
	err := row.Scan(
		&i.ID,
		&i.FarmCode,
		&i.FarmCoordinates,
		&i.FarmAirspace,
		&i.FarmLocation,
		&i.FarmGeolocation,
		&i.FarmCreationDate,
		&i.FarmContact,
	)
	return i, err
}

const getFarm = `-- name: GetFarm :one
SELECT id, farm_code, farm_coordinates, farm_airspace, farm_location, farm_geolocation, farm_creation_date, farm_contact
FROM farms
WHERE farm_code = $1
LIMIT 1
`

func (q *Queries) GetFarm(ctx context.Context, farmCode string) (Farm, error) {
	row := q.db.QueryRowContext(ctx, getFarm, farmCode)
	var i Farm
	err := row.Scan(
		&i.ID,
		&i.FarmCode,
		&i.FarmCoordinates,
		&i.FarmAirspace,
		&i.FarmLocation,
		&i.FarmGeolocation,
		&i.FarmCreationDate,
		&i.FarmContact,
	)
	return i, err
}

const getFarmForUpdate = `-- name: GetFarmForUpdate :one
SELECT id, farm_code, farm_coordinates, farm_airspace, farm_location, farm_geolocation, farm_creation_date, farm_contact
FROM farms
WHERE farm_code = $1
LIMIT 1 FOR NO KEY
UPDATE
`

func (q *Queries) GetFarmForUpdate(ctx context.Context, farmCode string) (Farm, error) {
	row := q.db.QueryRowContext(ctx, getFarmForUpdate, farmCode)
	var i Farm
	err := row.Scan(
		&i.ID,
		&i.FarmCode,
		&i.FarmCoordinates,
		&i.FarmAirspace,
		&i.FarmLocation,
		&i.FarmGeolocation,
		&i.FarmCreationDate,
		&i.FarmContact,
	)
	return i, err
}

const listFarms = `-- name: ListFarms :many
SELECT id, farm_code, farm_coordinates, farm_airspace, farm_location, farm_geolocation, farm_creation_date, farm_contact
FROM farms
ORDER BY farm_code
LIMIT $1
OFFSET $2
`

type ListFarmsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListFarms(ctx context.Context, arg ListFarmsParams) ([]Farm, error) {
	rows, err := q.db.QueryContext(ctx, listFarms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Farm
	for rows.Next() {
		var i Farm
		if err := rows.Scan(
			&i.ID,
			&i.FarmCode,
			&i.FarmCoordinates,
			&i.FarmAirspace,
			&i.FarmLocation,
			&i.FarmGeolocation,
			&i.FarmCreationDate,
			&i.FarmContact,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}