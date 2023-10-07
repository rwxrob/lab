# Setup Kubernetes from `kubeadm`

Home grown clusters need the following to be added:

* Which nodes be virtualised and if so how?
    Yes, KVM
    One or two GPU physical nodes

* What host OS will be used for nodes (including server)?
    Redhat (Rocky)

* What container engine will be used on nodes?
    containerd

* Which CNI provider?
    Calico (maybe Flannel or Cillium later)

* Which CSI provider(s)?
    Generic NFS (since most things have compat layer)
    Ceph or Luster later?

* Which ingress?
    Istio (overkill, but what we use at work)

* How will external DNS be managed?

* Which hardware load balancer?
    MetalLb

* How will certificates be created and managed?
    Vault

* How will security be managed?
    Keycloak OAuth2 User Flow

* How will out-of-cluster secrets be managed?
    Vault

* Will node need GPU support?
    Yes, install GPU Feature Discovery

* How will cluster creation be automated?
    Ansible (drawing on Kubespray recipes)

## Know why we are building an on-prem Kubernetes cluster

* Cost especially for high-GPU utilization
* Lack of GPU support in the cloud
* Full training in all aspects of cluster setup (CNI, Ingress, StorageClass, etc.)

## Know why we are not using a turn-key Kubernetes solution (like OpenShift)

* Cost
* Availability
* Proprietary lock-in
* Lack of customization, drivers, etc.

## What type of hardware nodes are we going to use?

* KVM virtual machine nodes (no GPU)
* One or more hardware nodes (with GPU)

## What host operating system should we use?

* Do I really need to understand the host OS well?
* RedHat is more common in the enterprise
* Ubuntu is easier
* RedHat is what my company uses currently
* RedHat is more likely to play nice with Ansible
* RedHat has selinux, etc. making more enterprise-y
* RedHat Marketplace is a thing

## What container engine are we going to use?

* `CRI-O` does *not* allow `podman`/`docker` to work "locally"
* `docker-ce` is Docker (and I hate Docker)
* `containerd` is Docker (but maintained differently)
* Using `containerd` primarily because work chose to use it

## How are we going to deploy all of this automatically?

* Ansible + scripts

## What authentication strategy are we using for the cluster?

* User Auth Flow (no refresh token) for "users"
* Client-Credentials for applications
* Keycloak provider

## How will enterprise secrets be managed?

* Vault and vault Kubernetes operator
* Kubernetes certs managed by Kubernetes

## Are we going to setup own RootCA?

* Yes, hopefully using Vault

## How much "high availability" are we going to emulate?

## How will storage be managed?

* One NFS server as StorageClass using existing machine
* Decide on other CSI options later
* Eventually, Rook/Ceph solution with 3-6 machines
* Eventually, perhaps a Lustre-based CSI
* <https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner>

--------------------------------------------------------------------

***OLD STUFF***

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
