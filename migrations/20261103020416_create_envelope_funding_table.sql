-- +goose Up
CREATE TABLE envelope_funding (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_id integer NOT NULL REFERENCES categories (id),
    month date NOT NULL,
    amount_cents bigint NOT NULL
);

-- +goose Down
DROP TABLE envelope_funding;
