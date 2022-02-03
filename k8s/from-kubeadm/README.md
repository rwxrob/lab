# Setup Kubernetes from `kubeadm`

*Note: initially this lab used Vagrant but I've decided to use VMware
base images with snapshots since they are faster to bring up and use
Ansible for configuration where needed for additional configurations.*

## But first ...

This particular lab makes several base assumptions about the approach
that are consistent with the tools most used by working professionals.

* Setup a Linux workstation (bare metal, WSL2 is fine)
* Install VMware Workstation Pro (might work with VirtualBox)
* Learn `cloud-init` for creating VMs
* Learn how to setup 3 VMs with NFS enabled

## Tasks

1. **Get a cloud-enabled version of enterprise Linux** (preferably
   RedHat variation like Alma, Rocky, CentOS).

1. **Create three virtual machines with `cloud-init`** each with 2
   cores, 2 GB RAM, 10 GB storage named `control`, `node1` and `node2`
   (with matching host names).

1. **Create a small NFS volume** at the (`/s`) mount point. It should be
   able to be mounted from the `node` workers (but don't mount them,
   StorageClass provider will cover remote access).

1. **Install the Kubernetes components**
   Provision each with the standard Kubernetes components
   and any other administrative tools you wish.

1. **Initialize control-plane with `kubeadm init ...`**

1. **Join the worker nodes** 

1. **Deploy Calico CNI** so that you have networking between your pods.
   Warning: this step requires some care when working with Vagrant
   (which has horrible networking and takes over the default `eth0`
   interface).

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
* <https://levelup.gitconnected.com/how-to-use-nfs-in-kubernetes-cluster-storage-class-ed1179a83817>

* <https://github.com/vagrant-libvirt/vagrant-libvirt#additional-disks>
