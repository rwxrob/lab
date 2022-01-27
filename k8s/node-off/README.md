# Offline-ing a Kubernetes Node Safely

## But first ...

* **Setup a test Kubernetes cluster of virtual machines** that can be
  with Vagrant, Minikube (just make sure not using Docker emulation).

## Tasks

1. **Create a fair number of running pods on both worker Nodes.** The
   pods can be from any kind of Kubernetes resource.

1. **Clear one of the Nodes from all running pods** so that taking it
   offline --- even if it were powered off --- won't harm anything that
   was running.

1. **Restart the cleared Node and bring it back into the cluster.**
   Monitor that everything gets distributed properly over time.
