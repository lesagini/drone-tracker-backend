-- name: CreateFarm :one
INSERT INTO farms (
        farm_code,
        farm_coordinates,
        farm_airspace,
        farm_location,
        farm_geolocation,
        farm_contact
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFarm :one
SELECT *
FROM farms
WHERE farm_code = $1
LIMIT 1;
-- name: GetFarmForUpdate :one
SELECT *
FROM farms
WHERE farm_code = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListFarms :many
SELECT *
FROM farms
ORDER BY farm_code
LIMIT $1
OFFSET $2;