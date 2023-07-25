// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: operator.sql

package models

import (
	"context"
)

const createOperator = `-- name: CreateOperator :one
INSERT INTo operators (
        operator_id,
        operator_name,
        operator_headquater,
        operator_number_pilots_deployed,
        opertor_contact
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, operator_id, operator_name, operator_headquater, operator_number_pilots_deployed, opertor_contact
`

type CreateOperatorParams struct {
	OperatorID                   string
	OperatorName                 string
	OperatorHeadquater           string
	OperatorNumberPilotsDeployed int32
	OpertorContact               int64
}

func (q *Queries) CreateOperator(ctx context.Context, arg CreateOperatorParams) (Operator, error) {
	row := q.db.QueryRowContext(ctx, createOperator,
		arg.OperatorID,
		arg.OperatorName,
		arg.OperatorHeadquater,
		arg.OperatorNumberPilotsDeployed,
		arg.OpertorContact,
	)
	var i Operator
	err := row.Scan(
		&i.ID,
		&i.OperatorID,
		&i.OperatorName,
		&i.OperatorHeadquater,
		&i.OperatorNumberPilotsDeployed,
		&i.OpertorContact,
	)
	return i, err
}
