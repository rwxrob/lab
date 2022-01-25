# Setup Kubernetes from `kubeadm` (Vagrant)

## But first ...

* Setup a true Linux workstation (on bare metal)
* Install VirtualBox
* Install Vagrant
* Learn how to setup 3 Vagrant VMs with NFS enabled

## Tasks

1. **Create three virtual machines** with enterprise Linux (preferably
   RedHat variants: CentOS, ALMA), each with 2 cores, 2 GB RAM, 10 GB
   storage named `control`, `node1` and `node2` (with matching
   host names). 

1. **Create a small NFS volume** at the (`/s`) mount point. It should be
   able to be mounted from the `node` workers (but don't mount them,
   StorageClass provider will cover remote access).

1. **Install the Kubernetes components**
   Provision each with the standard Kubernetes components
   and any other administrative tools you wish.

1. **Install Kubernetes from `kubeadm`**

1. **Run `kubeadm init` on `control`** to create a control plane and
   create the first node (local to the control plane).

1. **Deploy a Pod CNI** so that you have networking between your pods
   (preferably Calico).

1. **Join the worker nodes** 

1. **Create a StorageClass for the NFS Volume**

Bonus:

1. **Setup ssh login from `control` to `node*`** using pre-generated
   certs. Disable password login to all nodes in favor of certificates.
   (Remember, this is for a local test environment.)

Related:

* Machine Settings     
  <https://www.vagrantup.com/docs/vagrantfile/machine_settings>
* Installing kubeadm \| Kubernetes  
  <https://v1-21.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/>
* Creating a cluster with kubeadm \| Kubernetes  
  <https://v1-21.docs.kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>
* kubernetes - Kubelet service is not running. It seems like the kubelet isn\'t running or healthy - Stack Overflow  
  <https://stackoverflow.com/questions/70229371/kubelet-service-is-not-running-it-seems-like-the-kubelet-isnt-running-or-healt>
* <https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner>
* <https://kubernetes.io/docs/reference/setup-tools/kubeadm/implementation-details/>
