export KIND_EXPERIMENTAL_PROVIDER := "podman"

# Create the Kind cluster and start the database
up:
    kind create cluster --name stashbook --wait 2m
    doppler secrets substitute deploy/db-secret.yaml | kubectl create -f -
    kubectl apply -f deploy/pod.yaml
    kubectl wait --for=condition=Ready pod/stashbook --timeout=2m

# Delete the Kind cluster, including the database data
down:
    kind delete cluster --name stashbook

# Connect localhost:5432 to the database (leave running)
db-forward:
    kubectl port-forward pod/stashbook 5432:5432

# Create a new migration file: just migration <name>
migration name:
    go tool goose -dir migrations create {{name}} sql

# Apply all pending migrations
migrate-up:
    doppler run --command='go tool goose -dir migrations postgres "$DATABASE_URL" up'

# Roll back the most recent migration
migrate-down:
    doppler run --command='go tool goose -dir migrations postgres "$DATABASE_URL" down'

# Show which migrations have been applied
migrate-status:
    doppler run --command='go tool goose -dir migrations postgres "$DATABASE_URL" status'

# Generate Go code from sqlc
generate:
    go tool sqlc generate
