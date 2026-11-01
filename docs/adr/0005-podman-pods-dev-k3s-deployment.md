# 5. Podman pods now, K3s for deployment later

Date: 2026-11-01

## Status

Accepted

## Context

Stashbook runs several services together: the Go app, its PostgreSQL database, Paperless-ngx with its own PostgreSQL database and Redis, and Newt for the Pangolin tunnel.

They need to run together on one machine now. Kubernetes is also something to learn on this project, even though it's more than the app strictly needs.

## Decision

Run all services as containers in Podman pods for now.

Move to K3s, a lightweight Kubernetes distribution, for the final deployment on a single small server.

## Consequences

- PostgreSQL, Paperless-ngx, and Redis are each just another container, so they add little overhead.
- Podman runs without a root daemon.
- K3s stores pod data on the server's own disk by default, so backups must still copy that data somewhere else.
- Pangolin and Newt handle incoming traffic, so the Traefik ingress that ships with K3s isn't needed.
- The move to K3s means writing Kubernetes manifests for every service.
