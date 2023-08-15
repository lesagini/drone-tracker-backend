-- name: CreateFlight :one
INSERT INTO flights (
        flight_farm_id,
        flight_duration,
        flight_pilot,
        flight_acreage
    )
VALUES ($1, $2, $3, $4)
RETURNING *;