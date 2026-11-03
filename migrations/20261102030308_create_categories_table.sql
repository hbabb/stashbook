-- +goose Up
CREATE TABLE categories (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE categories;
