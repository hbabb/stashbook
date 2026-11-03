-- +goose Up
CREATE TYPE frequency_type AS ENUM (
    'monthly',
    'annual'
);

CREATE TABLE subscriptions (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payee varchar(100) NOT NULL,
    amount_cents bigint NOT NULL,
    account_id integer NOT NULL REFERENCES accounts (id),
    category_id integer REFERENCES categories (id),
    frequency frequency_type NOT NULL,
    next_due date NOT NULL
);

-- +goose Down
DROP TABLE subscriptions;
DROP TYPE frequency_type;
