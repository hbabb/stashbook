# 5. Kind for development, K3s for deployment

Date: 2026-11-01

## Status

Accepted

## Context

Stashbook runs several services together: the Go app, its PostgreSQL database, Paperless-ngx with its own PostgreSQL database and Redis, and Newt for the Pangolin tunnel.

Kubernetes is something to learn on this project, even though it's more than the app strictly needs. Development and deployment should run the same Kubernetes manifests, the YAML files that describe each service.

## Decision

Use Kind for development. Kind runs a local Kubernetes cluster with each node as a container, using Podman.

Use K3s, a lightweight Kubernetes distribution, for deployment on a single small server.

## Consequences

- The same manifests run in development and in deployment.
- Kind clusters start and tear down quickly, which fits rebuilding often during development.
- Kind's Podman support is officially marked experimental.
- Deleting a Kind cluster deletes the data inside it, including the development database.
- K3s stores pod data on the server's own disk by default, so backups must still copy that data somewhere else.
- Pangolin and Newt handle incoming traffic, so the Traefik ingress that ships with K3s isn't needed.
- `kubectl`, the Kubernetes command-line tool, is needed in both environments.
