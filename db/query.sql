-- name: GetPilot :one
SELECT *
FROM pilots
WHERE pilot_id = $1
LIMIT 1;
-- name: GetPilotForUpdate :one
SELECT *
FROM pilots
WHERE pilot_id = $1
LIMIT 1
FOR NO KEY UPDATE;
-- name: ListFarms :many
SELECT *
FROM farms
ORDER BY farm_code;
-- name: CreatePilot :one
INSERT INTO pilots (
        pilot_id,
        pilot_operator_id,
        pilot_initials,
        pilot_number,
        pilot_full_name,
        pilot_license_number,
        pilot_farm_location_code,
        pilot_farm_location,
        pilot_status,
        pilot_classification,
        pilot_flight_hours,
        pilot_covered_acreage
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    )
RETURNING *;
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
-- name: CreateOperator :one
INSERT INTo operators (
        operator_id,
        operator_name,
        operator_headquater,
        operator_number_pilots_deployed,
        opertor_contact
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

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

-- name: CreateVariety :one
INSERT INTO varieties (
    variety_internal_identity,
    variety_botanical_name,
    variety_farm_id,
    variety_acreage,
    variety_type,
    variety_iterval_code
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdatePilot :one
UPDATE pilots
  set pilot_flight_hours = $2,
  pilot_covered_acreage = $3
WHERE pilot_id = $1
RETURNING *;