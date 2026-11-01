# 2. PostgreSQL over SQLite

Date: 2026-11-01

## Status

Accepted

## Context

The ledger's data must be validated by the database itself: types, string lengths, required fields, and fixed lists of allowed values.

By default, SQLite ignores the length in `varchar(n)` and doesn't enforce column types. Enforcing them needs `STRICT` tables plus a separate `CHECK` constraint for every length and allowed-value rule. That is extra work for something that should be built in.

## Decision

Use PostgreSQL 18. Schema changes are managed as SQL migrations with goose.

## Consequences

- `varchar(n)`, `NOT NULL`, and column types are enforced with no extra rules.
- Native `date` and `boolean` types, and real enums for fixed lists like account type.
- Adding a value to an enum later needs a migration (`ALTER TYPE`), not just a new row.
- PostgreSQL runs as its own container, which is one more container in the pod.
- Backups use `pg_dump` instead of copying a single file.
