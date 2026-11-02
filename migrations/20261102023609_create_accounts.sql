-- +goose Up
CREATE TYPE account_type AS ENUM (
    'checking',
    'savings',
    'loan',
    'credit_card',
    'cash',
    'brokerage',
    'retirement'
);

CREATE TYPE account_owner AS ENUM (
    'personal',
    'business'
);

CREATE TABLE accounts (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE,
    type account_type NOT NULL,
    owner account_owner NOT NULL DEFAULT 'personal',
    opening_balance_cents bigint NOT NULL DEFAULT 0
);
-- +goose Down
DROP TABLE accounts;
DROP TYPE account_owner;
DROP TYPE account_type;
