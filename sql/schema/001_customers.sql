--+goose Up

CREATE TABLE customers (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

--+goose Down
DROP TABLE customers;
