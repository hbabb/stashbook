-- name: CreateEntry :one
INSERT INTO entries (
    account_id, date, payee, debit_cents, credit_cents, category_id,
    business, reimbursement, cleared, paperless_doc_id, memo
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: ListEntriesByAccount :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY date, id;

-- name: UpdateEntry :one
UPDATE entries
SET
    account_id = $2,
    date = $3,
    payee = $4,
    debit_cents = $5,
    credit_cents = $6,
    category_id = $7,
    business = $8,
    reimbursement = $9,
    cleared = $10,
    paperless_doc_id = $11,
    memo = $12
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: ListUnclearedEntries :many
SELECT * FROM entries
WHERE account_id = $1 AND NOT cleared
ORDER BY date, id;

-- name: SetEntryCleared :exec
UPDATE entries
SET cleared = $2
WHERE id = $1;

-- name: GetAccountBalance :one
SELECT (
    a.opening_balance_cents
    + COALESCE(SUM(e.credit_cents - e.debit_cents), 0)
)::bigint AS balance_cents
FROM accounts AS a
LEFT JOIN entries AS e ON a.id = e.account_id
WHERE a.id = $1
GROUP BY a.id;

-- name: ListEntriesByCategoryAndDates :many
SELECT * FROM entries
WHERE
    category_id = sqlc.arg(category_id)
    AND date >= sqlc.arg(start_date)
    AND date <= sqlc.arg(end_date)
ORDER BY date, id;

-- name: ListBusinessEntries :many
SELECT * FROM entries
WHERE business = $1
ORDER BY date, id;

-- name: ListEntriesByReimbursement :many
SELECT * FROM entries
WHERE reimbursement = $1
ORDER BY date, id;

-- name: SetReimbursementStatus :exec
UPDATE entries
SET reimbursement = $2
WHERE id = $1;
