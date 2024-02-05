-- name: CreateUser :one
INSERT INTO customers (
    id, created_at, updated_at, name, api_key
) VALUES (
    $1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetCustomers :one
SELECT * FROM customers WHERE api_key = $1;
