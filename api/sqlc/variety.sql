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