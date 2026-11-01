# Stashbook

A personal finance ledger and envelope budget. Hand-entered debit and credit entries per account, reconciled monthly, with receipts stored in Paperless-ngx.

## Features

- A ledger for each account: checking, savings, credit cards, loans, brokerage, and retirement
- Monthly reconciliation against bank statements
- Envelope budgeting by category
- Recurring subscriptions, monthly or annual
- Receipts linked from Paperless-ngx
- Reports for medical expenses, W-2 business expenses (reimbursed and non-reimbursed), and side-hustle expenses

## Stack

- **Backend:** Go, chi, pgx, sqlc
- **Frontend:** Svelte on Vite+, embedded in the Go binary
- **Database:** PostgreSQL 18, with migrations by goose
- **Receipts:** Paperless-ngx
- **Secrets:** Doppler
- **Deployment:** Podman pods, with K3s planned

## Status

Early development.
