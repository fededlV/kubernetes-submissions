# Volumes

Storage shared by **Ping-pong** and **Log output**. Kept apart from the apps because
PersistentVolumes are cluster-level resources, usually managed by cluster administrators.

- `persistentvolume.yaml`: `shared-pv`, 1Gi `local` volume at `/tmp/kube` on node
  `k3d-k3s-default-agent-0` (`nodeAffinity` → pods using it are scheduled on that node).
- `persistentvolumeclaim.yaml`: `shared-claim`, bound to `shared-pv` through
  `storageClassName: shared-pv-class`. Both deployments reference this claim.

`ReadWriteOnce` works for two pods because it is per **node**, and both pods run on the same node.

## Setup
```
docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube
kubectl apply -f .
kubectl get pv,pvc      # STATUS: Bound
```

## Clean up
Delete the deployments first, then:
```
kubectl delete -f .
```
A `Released` PV does not rebind automatically: delete and re-create it.