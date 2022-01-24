# Learning Lab: Vagrant Local Kubernetes from `kubeadm`

## But first ...

* Setup a true Linux workstation (on bare metal)
* Install `libvirt`/QEMU/KVM
* Install Vagrant and configure for `libvirt`

## Tasks

1. **Create three virtual machines** with enterprise Linux (preferably
   RedHat variants: CentOS, ALMA), each with 2 cores, 2 GB RAM, 10 GB
   storage named `control`, `node-1` and `node-2` (with matching
   host names). Provision each with the standard Kubernetes components
   and any other administrative tools you wish.

1. **Create an NFS volume** and mount on all three systems at the same
   mount point location.

1. **Setup ssh login from `control` to `node-*`** using pre-generated
   certs. Disable password login to all nodes in favor of certificates.
   (Remember, this is for a local test environment.)

1. **Run `kubeadm init` on `control`** to create a control plane and
   create the first node (local to the control plane).

1. **Deploy a Pod CNI** so that you have networking between your pods
   (preferably Calico).

1. **Join the worker nodes** 

1. **Create a StorageClass for the NFS Volume**

Related:

* Machine Settings     
  <https://www.vagrantup.com/docs/vagrantfile/machine_settings>
* Installing kubeadm \| Kubernetes  
  <https://v1-21.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/>
* Using kubeadm to Create a Cluster  
  <https://v1-21.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>
* Creating a cluster with kubeadm \| Kubernetes  
  <https://v1-21.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>
* kubernetes - Kubelet service is not running. It seems like the kubelet isn\'t running or healthy - Stack Overflow  
  <https://stackoverflow.com/questions/70229371/kubelet-service-is-not-running-it-seems-like-the-kubelet-isnt-running-or-healt>
