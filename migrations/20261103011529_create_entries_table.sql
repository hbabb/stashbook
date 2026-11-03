-- +goose Up
CREATE TYPE business_type AS ENUM (
    'none',
    'w2',
    'side_hustle'
);

CREATE TYPE reimbursement_status AS ENUM (
    'pending',
    'reimbursed',
    'not_reimbursed'
);

CREATE TABLE entries (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_id integer NOT NULL REFERENCES accounts (id),
    date date NOT NULL,
    payee varchar(100) NOT NULL,
    debit_cents bigint NOT NULL DEFAULT 0,
    credit_cents bigint NOT NULL DEFAULT 0,
    category_id integer REFERENCES categories (id),
    business business_type NOT NULL DEFAULT 'none',
    reimbursement reimbursement_status,
    cleared boolean NOT NULL DEFAULT false,
    paperless_doc_id integer,
    memo varchar(255)
);
-- +goose Down
DROP TABLE entries;
DROP TYPE reimbursement_status;
DROP TYPE business_type;
