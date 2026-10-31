# Log output

Genera un UUID al iniciar y lo imprime cada 5 segundos junto a un timestamp.

## Run locally
```
go run .
```

## Image
```
docker build -t TU_USUARIO_DOCKERHUB/log-output:1.1 .
docker push TU_USUARIO_DOCKERHUB/log-output:1.1
```

## Deploy in Kubernetes
```
kubectl apply -f manifests/deployment.yaml
kubectl logs -f deployment/log-output-dep
```

## Clean up
```
kubectl delete -f manifests/deployment.yaml
```