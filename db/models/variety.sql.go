// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: variety.sql

package models

import (
	"context"
)

const createVariety = `-- name: CreateVariety :one
INSERT INTO varieties (
        variety_internal_identity,
        variety_botanical_name,
        variety_farm_id,
        variety_acreage,
        variety_type,
        variety_iterval_code
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, variety_internal_identity, variety_botanical_name, variety_farm_id, variety_creation_date, variety_acreage, variety_type, variety_iterval_code
`

type CreateVarietyParams struct {
	VarietyInternalIdentity string
	VarietyBotanicalName    string
	VarietyFarmID           string
	VarietyAcreage          int64
	VarietyType             VarietyTypes
	VarietyItervalCode      string
}

func (q *Queries) CreateVariety(ctx context.Context, arg CreateVarietyParams) (Variety, error) {
	row := q.db.QueryRowContext(ctx, createVariety,
		arg.VarietyInternalIdentity,
		arg.VarietyBotanicalName,
		arg.VarietyFarmID,
		arg.VarietyAcreage,
		arg.VarietyType,
		arg.VarietyItervalCode,
	)
	var i Variety
	err := row.Scan(
		&i.ID,
		&i.VarietyInternalIdentity,
		&i.VarietyBotanicalName,
		&i.VarietyFarmID,
		&i.VarietyCreationDate,
		&i.VarietyAcreage,
		&i.VarietyType,
		&i.VarietyItervalCode,
	)
	return i, err
}