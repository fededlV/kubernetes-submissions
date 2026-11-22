# Ping-pong

Servidor HTTP en Go que responde `pong N` a cada `GET /pingpong`.
El contador se guarda en un archivo (`COUNT_FILE`, por defecto `/usr/src/app/files/pingpong.txt`)
dentro de un **PersistentVolume** compartido con "Log output", que lo muestra como `Ping / Pongs: N`.
Al arrancar lee el valor guardado, así que el conteo sobrevive a reinicios del pod.
El puerto se toma de la variable de entorno `PORT` (por defecto 8080).

## Run locally
```
PORT=3000 COUNT_FILE=../log-output/shared/pingpong.txt go run .
curl http://localhost:3000/pingpong   # pong 0
curl http://localhost:3000/pingpong   # pong 1
```

## Image
```
docker build -t fededlv/ping-pong:1.11 .
docker push fededlv/ping-pong:1.11
```

## Deploy in Kubernetes (k3d)
Usa el PVC `shared-claim` definido en [`../volumes`](../volumes), que hay que aplicar antes.
Comparte el Ingress con "Log output": `/` va a log-output y `/pingpong` a ping-pong.
El Ingress está en `log-output/manifests/ingress.yaml`.
```
kubectl apply -f ../volumes/
kubectl apply -f manifests/
kubectl apply -f ../log-output/manifests/
kubectl get pods,svc,ing
curl http://localhost:8081/pingpong
```

## Clean up
```
kubectl delete -f manifests/
```