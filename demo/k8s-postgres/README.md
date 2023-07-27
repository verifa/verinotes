# Kubernetes + Postgres demo

This setups a postgres and connects VeriNotes to it, it assumes you have a cluster (such as kind or k3d) running and uses your default kubectl context.

## Usage

```bash
kubectl create namespace verinotes
kubectl apply -f postgres.yaml -n verinotes
```

I've hardcoded the credentials for postgres into the manifest, so just apply it:

```bash
kubectl apply -f verinotes-deployment.yaml -n verinotes
```
