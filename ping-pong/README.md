# Ping-pong

Servidor HTTP en Go que responde `pong N` a cada `GET /pingpong`, donde `N` es un contador en memoria que empieza en 0 (se reinicia si el pod se reinicia).
El puerto se toma de la variable de entorno `PORT` (por defecto 8080).

## Run locally
```
PORT=3000 go run .
curl http://localhost:3000/pingpong   # pong 0
curl http://localhost:3000/pingpong   # pong 1
```

## Image
```
docker build -t fededlv/ping-pong:1.9 .
docker push fededlv/ping-pong:1.9
```

## Deploy in Kubernetes (k3d)
Comparte el Ingress con "Log output": `/` va a log-output y `/pingpong` a ping-pong.
El Ingress está en `log-output/manifests/ingress.yaml`.
```
kubectl apply -f manifests/
kubectl apply -f ../log-output/manifests/
kubectl get pods,svc,ing
curl http://localhost:8081/pingpong
```

## Clean up
```
kubectl delete -f manifests/
```