# Learning Lab: Kubernetes PV/PVC Storage Monitor

The goal of this lab is to get the exact percent available of a given
PVC volume for a specific K8SAPP application.

> ⚠️
> You may want to do all of this on a virtual machine rather than risk
> your main workstation.

1. Create a minimal Go program and encapsulating FROM SCRATCH container
   that purposefully consumes disk in a file at a specific path so that
   it can be specifically mounted as a Docker volume and later as a PVC
   to test storage consumption monitoring.

1. Create a local Kubernetes test environment (preferably Kind) and
   setup a minimal application with a small amount of assigned storage
   (PV/PVC) using the default storage class.

1. ~~Familiarize yourself with the Kubernetes API, specifically the
   queries for current storage utilization of a specified PVC.~~ (There
   is no API for this) Create a monitoring Pod (or sidecar container)
   that calls `du`/`df` and outputs to logs.

1. Experiment with ways to have monitor send events and status
   information to any number of established monitoring solutions
   (perhaps Icinga).

1. Package the Go PVC monitor as a FROM SCRATCH container and write the
   simplest deployment possible (`kompose` or just bash script). This
   could also just be done with BusyBox to get started.

Related:

* <https://github.com/rwxrob/filldisk>
* Volume Health Monitoring \| Kubernetes  
  <https://kubernetes.io/docs/concepts/storage/volume-health-monitoring/>
* <https://www.robustperception.io/filesystem-metrics-from-the-node-exporter>
* <https://github.com/kubernetes-csi/csi-driver-nfs>
* <https://github.com/NickrenREN/kubemonitor/tree/master/build/kube_storage_monitor/local_monitor>
* Sidecar Containers - Kubernetes CSI Developer Documentation  
  <https://kubernetes-csi.github.io/docs/sidecar-containers.html>
* Container Storage Interface (CSI)  
  <https://github.com/container-storage-interface/spec/blob/master/spec.md#nodevolumestats>
* <https://github.com/kubernetes/kubernetes/blob/v2.21.0-alpha.2/pkg/volume/csi/csi_metrics.go#L71>
* How to create PV and PVC in Kubernetes - Knoldus Blogs  
  <https://blog.knoldus.com/how-to-create-pv-and-pvc-in-kubernetes/>
* Reclaiming Persistent Volumes in Kubernetes \| by DataStax \| Building the Open Data Stack \| Medium  
  <https://medium.com/building-the-open-data-stack/reclaiming-persistent-volumes-in-kubernetes-5e035ba8c770>
