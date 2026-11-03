-- name: CreateAccount :one
INSERT INTO accounts (name, type, owner, opening_balance_cents)
VALUES ($1, $2, $3, $4)
RETURNING id, name, type, owner, opening_balance_cents;

-- name: GetAccount :one
SELECT
    id,
    name,
    type,
    owner,
    opening_balance_cents
FROM accounts
WHERE id = $1;

-- name: ListAccounts :many
SELECT
    id,
    name,
    type,
    owner,
    opening_balance_cents
FROM accounts
ORDER BY name;

-- name: UpdateAccount :one
UPDATE accounts
SET
    name = $2,
    type = $3,
    owner = $4,
    opening_balance_cents = $5
WHERE id = $1
RETURNING id, name, type, owner, opening_balance_cents;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
