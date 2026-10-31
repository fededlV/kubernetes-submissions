# Todo app

Servidor web en Go. El puerto se configura con la variable de entorno `PORT` (por defecto 3000).

## Run locally
```
PORT=8080 go run .
```

## Image
```
docker build -t TU_USUARIO/todo-app:1.2 .
docker push TU_USUARIO/todo-app:1.2
```

## Deploy in Kubernetes
```
kubectl create deployment todo-app --image=TU_USUARIO/todo-app:1.2
kubectl set env deployment/todo-app PORT=8080
kubectl logs -f deployment/todo-app
```