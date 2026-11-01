# 4. Doppler for secrets

Date: 2026-11-01

## Status

Accepted

## Context

Stashbook needs secrets: database passwords, connection strings, the Paperless-ngx API token, and the Pangolin tunnel credentials.

Secrets must never be committed to the repo, and `.env` files are easy to commit by mistake or leave lying around.

## Decision

Keep all secrets in Doppler, in a `stashbook` project with a `dev` config and a `prd` config.

Secrets are added and changed in the Doppler dashboard.

Commands that need secrets run under `doppler run`, which injects them as environment variables. The Go app reads its settings from environment variables with `caarlos0/env`.

## Consequences

- No `.env` files exist. `.gitignore` still excludes `.env` and `.env.*` as a safety net.
- Development and production use the same variable names, with different values per config.
- Running the app or migrations needs the Doppler CLI and a `doppler setup` in the project folder.
