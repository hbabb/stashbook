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
