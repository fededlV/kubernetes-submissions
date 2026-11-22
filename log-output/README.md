# Log output

Split into two containers running in the same pod. Both mount a **PersistentVolumeClaim**
(`shared-claim`) that is also mounted by the **Ping-pong** app, so the two apps share data.

- **log-writer**: generates a random string (UUID) on startup and appends a line
  `timestamp: random-string` to `/usr/src/app/files/output.log` every 5 seconds.
- **log-reader**: on HTTP GET `/` (port 3000) returns the last line of the log plus the
  Ping-pong request count read from `/usr/src/app/files/pingpong.txt`:

```
2026-11-21T23:40:12.345Z: 8523ecb1-c716-4cb6-a044-b9e83bb98e43.
Ping / Pongs: 3
```

```
log-output/
├── log-writer/    # Go app that writes the log file
├── log-reader/    # Go app that serves last log line + ping-pong count
└── manifests/     # deployment (2 containers + PVC), service, ingress
```

The PersistentVolume and PersistentVolumeClaim live in [`../volumes`](../volumes), separate
from the app, since PVs are usually managed by cluster administrators.

## Run locally
All apps must use the same folder:
```
cd log-writer && LOG_FILE=../shared/output.log go run .
cd log-reader && LOG_FILE=../shared/output.log COUNT_FILE=../shared/pingpong.txt go run .
curl localhost:3000
```

## Images
```
docker build -t fededlv/log-output-writer:1.11 ./log-writer
docker build -t fededlv/log-output-reader:1.11 ./log-reader
docker push fededlv/log-output-writer:1.11
docker push fededlv/log-output-reader:1.11
```

## Deploy in Kubernetes (k3d)
```
docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube   # folder backing the PV
kubectl apply -f ../volumes/                             # PV + PVC first
kubectl get pv,pvc                                       # both Bound
kubectl apply -f ../ping-pong/manifests/
kubectl apply -f manifests/
kubectl get pods -o wide                                 # both pods on agent-0
curl http://localhost:8081/pingpong
curl http://localhost:8081
```

## Clean up
```
kubectl delete -f manifests/
kubectl delete -f ../ping-pong/manifests/
kubectl delete -f ../volumes/
```