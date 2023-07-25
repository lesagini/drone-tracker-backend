-- name: CreateField :one
INSERT INTO fields (
        field_name,
        field_type,
        field_farm_id,
        field_variety_id,
        field_polygon,
        field_area,
        field_dieback,
        field_stage_name,
        field_status,
        field_notes
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
ORDER BY field_name;
-- name: UpdateField :one
UPDATE fields
set field_name = $1,
    field_farm_id = $2,
    field_type = $3,
    field_variety_id = $4,
    field_polygon = $5,
    field_area = $6,
    field_dieback = $7,
    field_stage_name = $8,
    field_status = $9,
    field_notes = $10
WHERE field_name = $11 AND field_farm_id = $12
RETURNING *;