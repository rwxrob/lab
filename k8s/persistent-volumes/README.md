# Lab: Kubernetes PersistentVolumes

This is just a place to experiment with Kubernetes `PersistentVolumes` and
`PersistentVolumeClaims` and how to assign them. I wanted to see if I
could just get a local simple `hostPath` to work. Once I got all these
pieces working together, and was able to watch them wait for first bind
and stuff I feel confident using other types of PV/PVC that more
standard. I'm looking forward, for example, to getting GlusterFS working
at home eventually.

## Commands Used

```
kind create cluster --name pv
k apply -f .
watch k get pods
k get pv
k get pvc
k describe pv test-volume
k describe pvc test-volume-claim
k explain pv
k explain pv.spec
k explain pvc
k explain pvc.spec
```

## Gotchas

* Make sure the names match when associating volumes with claims.

## Observations

* References to "manual" for `storageClassName` are either wrong or
  outdated. Should be "standard" for locally provisioned mounts.

* Kind keeps them all as Docker volumes, which after some digging I found
to be saved in `/var/lib/docker/volumes` when using `standard` and
`hostPath` (for learning).

## Related Resource

* <https://youtu.be/uSxlgK1bCuA>
