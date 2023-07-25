-- name: CreateFlight :one
INSERT INTO flights (
        flight_farm_location,
        flight_farm_id,
        flight_farm_geolocation,
        flight_duration,
        flight_pilot,
        flight_acreage
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;