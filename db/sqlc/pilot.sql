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
-- name: GetPilot :one
SELECT *
FROM pilots
WHERE pilot_id = $1
LIMIT 1;
-- name: GetPilotForUpdate :one
SELECT *
FROM pilots
WHERE pilot_id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListPilots :many
SELECT *
FROM pilots
ORDER BY pilot_id;
-- name: UpdatePilot :one
UPDATE pilots
set pilot_flight_hours = $2,
    pilot_covered_acreage = $3
WHERE pilot_id = $1
RETURNING *;