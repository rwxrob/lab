# Setup NFS on Vagrant Virtual Machines

This learning lab covers setting up a shared network "drive" using the
very old (but still commonly used) NFS protocol and exporting a single,
secondary drive on a `control` machine that is mounted remotely from the
`worker` machines so that all machines share the same path to the file
systems mounted on that `control` volume.

## But First ...

* Install VirtualBox
* Be able to create Vagrant virtual machines

> ⚠️
> There is the *old* way of adding disks and the *new* "experimental"
way that is much easier. Make sure you know the difference (and don't
forget to export VAGRANT_EXPERIMENTAL if using the new way).

## Skills



## Tasks

1. **Create three Vagrant virtual machines ** using the `virtualbox`
   driver with 2 GB RAM and 1 CPU each. Disable swap on all machines (if
   you want, although not required). Name the machines `control`, and
   `worker-1` and `worker-2`. Create a 10 MB secondary volume and mount
   and export it as an NFS volume ***on the `control` machine only***. Mount the NFS volume on the two workers.

1. Setup NFS server on `control`

1. Setup NFS clients on `worker-1` and `worker-2`

1. Setup a mounted/shared `/s` directory and export

Related:

* Official Vagrant "disks" Documentation  
  <https://www.vagrantup.com/docs/disks/usage>
* Vagrant - Adding a second hard drive (the old way)  
  <https://everythingshouldbevirtual.com/virtualization/vagrant-adding-a-second-hard-drive/>
* Creating an NFS Server with Vagrant and Archlinux for Kubernetes Cluster  
  <https://dev.to/lordrahl90/creating-an-nfs-server-with-vagrant-1oke>
* Configuring NFS Server on RedHat Enterprise 7  
  <https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/storage_administration_guide/nfs-serverconfig>
* Quick NFS Server configuration on Redhat 7 Linux System - Linux Tutorials - Learn Linux Configuration  
  <https://linuxconfig.org/quick-nfs-server-configuration-on-redhat-7-linux>
* Mounting a disk and creating partitions in Linux \| Instruction manual  
  <https://serverspace.io/support/help/mounting-a-disk-and-creating-partitions-in-linux/>
* How to Setup NFS Server on CentOS 8 / RHEL 8  
  <https://www.linuxtechi.com/setup-nfs-server-on-centos-8-rhel-8/>
* RPC  
  <https://www.guru99.com/remote-procedure-call-rpc.html>
* How to Create Partitions in Linux  
  <https://phoenixnap.com/kb/linux-create-partition>
