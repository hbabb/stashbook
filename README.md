# Stashbook

[![Status](https://img.shields.io/badge/status-early_development-orange)](#status)
[![License](https://img.shields.io/github/license/hbabb/stashbook)](LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/hbabb/stashbook)](go.mod)

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Svelte](https://img.shields.io/badge/Svelte-FF3E00?logo=svelte&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-9135FF?logo=vite&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![Paperless-ngx](https://img.shields.io/badge/Paperless--ngx-17541F?logo=paperlessngx&logoColor=white)
![Doppler](https://img.shields.io/badge/Doppler-555555)
![Kind](https://img.shields.io/badge/Kind-326CE5?logo=kubernetes&logoColor=white)
![K3s](https://img.shields.io/badge/K3s-FFC61C?logo=k3s&logoColor=black)
![Podman](https://img.shields.io/badge/Podman-892CA0?logo=podman&logoColor=white)

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
- **Development:** Kind on rootless Podman
- **Deployment:** K3s
- **Tooling:** just, golangci-lint, sqlfluff, Jujutsu

## Development

### Requirements

- Go
- Podman
- Kind and kubectl
- Doppler CLI
- just
- golangci-lint

### Running locally

1. Link the project to Doppler's `dev` config: `doppler setup`
2. Create the Kind cluster and start the database: `just up`
3. In a separate terminal, connect to the database: `just db-forward`
4. Apply the migrations: `just migrate-up`
5. Start the server: `just dev`

The API listens on `http://localhost:8080`.

### Before each commit

Run `just pre-commit`. It regenerates the sqlc code, tidies `go.mod`, formats, and lints.

### Docs

- Architecture decisions: [`docs/adr/`](docs/adr/)
- Database schema: [`docs/schema.dbml`](docs/schema.dbml)

## Status

Early development. The API serves `GET /api/accounts`.

## License

[MIT](LICENSE)
