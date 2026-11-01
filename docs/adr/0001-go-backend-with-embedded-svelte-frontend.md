# 1. Go backend with an embedded Svelte frontend

Date: 2026-11-01

## Status

Accepted

## Context

Stashbook needs an API and a reactive UI. It runs on one small server, and it should be simple to build, ship, and run.

A separate frontend server would mean a second process to deploy, and the API and UI could get out of sync.

## Decision

The backend is written in Go.

The frontend is plain Svelte with TypeScript, built with Vite+ from the `svelte-ts` Vite template. It is not SvelteKit.

The Svelte build output in `web/dist` is embedded in the Go binary with `//go:embed`. The embed lives in `web/`, next to `dist/`, because Go can't embed files from a parent folder.

One Go binary serves both the API under `/api/...` and the frontend files.

## Consequences

- Deployment is a single binary. Node.js is only needed on the development machine to build the frontend, not at runtime.
- The API and UI always ship together, so the API needs no versioning.
- All reactivity runs in the browser. Go only serves the built files.
- The frontend has to be built before the Go binary, so the build order matters.
