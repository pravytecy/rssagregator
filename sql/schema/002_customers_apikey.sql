--+goose Up
ALTER TABLE customers ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')
);

--+goose Down
ALTER TABLE customers DROP COLUMN api_key;
