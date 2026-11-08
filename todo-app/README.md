# Todo app

Servidor HTTP en Go que sirve una página HTML en `/`. El puerto se toma de la variable de entorno `PORT` (por defecto 8080).

## Run locally
```
PORT=3000 go run .
```

## Image
```
docker build -t TU_USUARIO_DOCKERHUB/todo-app:1.5 .
docker push TU_USUARIO_DOCKERHUB/todo-app:1.5
```

## Deploy in Kubernetes
```
kubectl apply -f manifests/deployment.yaml
kubectl get pods
kubectl logs -f deployment/todo-app-dep
```

## Test with port-forward
```
kubectl port-forward deployment/todo-app-dep 3003:3000
```
Abrir http://localhost:3003 en el navegador.