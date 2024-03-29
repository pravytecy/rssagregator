// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: customers.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO customers (
    id, created_at, updated_at, name, api_key
) VALUES (
    $1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING id, created_at, updated_at, name, api_key
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getCustomers = `-- name: GetCustomers :one
SELECT id, created_at, updated_at, name, api_key FROM customers WHERE api_key = $1
`

func (q *Queries) GetCustomers(ctx context.Context, apiKey string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomers, apiKey)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
