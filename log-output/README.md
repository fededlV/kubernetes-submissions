# Log output

Split into two containers running in the same pod, sharing an `emptyDir` volume:

- **log-writer**: generates a random string (UUID) on startup and appends a line
  `timestamp: random-string` to `/usr/src/app/files/output.log` every 5 seconds.
- **log-reader**: serves the contents of that file on HTTP GET `/` (port 3000).

```
log-output/
├── log-writer/    # Go app that writes the file
├── log-reader/    # Go app that serves the file over HTTP
└── manifests/     # deployment (2 containers + shared volume), service, ingress
```

## Run locally
Both apps must use the same file path:
```
cd log-writer && LOG_FILE=../shared/output.log go run .
cd log-reader && LOG_FILE=../shared/output.log go run .
curl localhost:3000
```

## Images
```
docker build -t fededlv/log-output-writer:1.10 ./log-writer
docker build -t fededlv/log-output-reader:1.10 ./log-reader
docker push fededlv/log-output-writer:1.10
docker push fededlv/log-output-reader:1.10
```

## Deploy in Kubernetes
```
kubectl apply -f manifests/
kubectl get pods                                              # READY 2/2
kubectl logs deploy/log-output-dep -c log-writer              # writer logs
kubectl logs deploy/log-output-dep --all-containers --prefix  # both containers
curl http://localhost:8081
```

## Clean up
```
kubectl delete -f manifests/
```