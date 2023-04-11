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