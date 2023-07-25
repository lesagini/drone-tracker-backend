-- name: CreateFarm :one
INSERT INTO fields (
        field_name,
        field_type,
        field_farm_id,
        field_variety_id,
        field_polygon,
        field_area,
        field_dieback,
        field_stage_name,
        field_status
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetField :one
SELECT *
FROM fields
WHERE field_name = $1 AND field_farm_id = $2
LIMIT 1;
-- name: GetFieldForUpdate :one
SELECT *
FROM fields
WHERE field_name = $1 AND field_farm_id = $2
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListFields :many
SELECT *
FROM fields
ORDER BY field_name
LIMIT $1
OFFSET $2;