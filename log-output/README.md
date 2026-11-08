# Log output

Genera un UUID al iniciar (guardado en memoria) y lo imprime cada 5 segundos junto a un timestamp.
Además expone un endpoint HTTP que devuelve el estado actual en JSON:

```
GET /status   ->  {"timestamp":"2026-11-08T13:30:00.000Z","id":"<uuid>"}
```

El puerto se toma de la variable de entorno `PORT` (por defecto 8080).

## Run locally
```
PORT=3000 go run .
curl http://localhost:3000/status
```

## Image
```
docker build -t fededlv/log-output:1.7 .
docker push fededlv/log-output:1.7
```

## Deploy in Kubernetes (k3d)
Crear el cluster exponiendo el puerto 80 del load balancer (si no se hizo antes):
```
k3d cluster create --port 8081:80@loadbalancer --agents 2
```

Aplicar los manifiestos:
```
kubectl apply -f manifests/
kubectl get pods,svc,ing
kubectl logs -f deployment/log-output-dep
```

Abrir http://localhost:8081/status en el navegador.

## Clean up
```
kubectl delete -f manifests/
```