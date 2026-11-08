# Todo app

Servidor HTTP en Go que sirve una página HTML en `/`. El puerto se toma de la variable de entorno `PORT` (por defecto 8080).

## Run locally
```
PORT=3000 go run .
```

## Image
```
docker build -t fededlv/todo-app:1.6 .
docker push fededlv/todo-app:1.6
```

## Deploy in Kubernetes (k3d)
Cluster con el puerto 80 del load balancer expuesto:
```
k3d cluster create --port 8081:80@loadbalancer --agents 2
```

Aplicar los manifiestos (Deployment + Service ClusterIP + Ingress):
```
kubectl apply -f manifests/
kubectl get pods,svc,ing
```

Abrir http://localhost:8081 en el navegador.

## Clean up
```
kubectl delete -f manifests/
```