# Todo app

Servidor web en Go. El puerto se configura con la variable de entorno `PORT` (por defecto 3000).

## Run locally
```
PORT=8080 go run .
```

## Image
```
docker build -t TU_USUARIO/todo-app:1.2 .# Todo app

Servidor web del proyecto del curso. Al iniciar imprime el puerto en el que escucha (configurable con la variable de entorno `PORT`).

## Run locally
```
PORT=3000 go run .
```

## Image
```
docker build -t TU_USUARIO_DOCKERHUB/todo-app:1.2 .
docker push TU_USUARIO_DOCKERHUB/todo-app:1.2
```

## Deploy in Kubernetes
```
kubectl apply -f manifests/deployment.yaml
kubectl logs -f deployment/todo-app-dep
```

## Clean up
```
kubectl delete -f manifests/deployment.yaml
```
docker push TU_USUARIO/todo-app:1.2
```

## Deploy in Kubernetes
```
kubectl create deployment todo-app --image=TU_USUARIO/todo-app:1.2
kubectl set env deployment/todo-app PORT=8080
kubectl logs -f deployment/todo-app
```