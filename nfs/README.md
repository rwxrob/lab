# Setup NFS on Virtual Machines

*Note: Original lab used vagrant. That work is now in
[vagrant](vagrant), but going forward I'm using VMware and Ansible for
everything. Also, originally I used VirtualBox, which will likely still
work, but I've decided to use the professional tool of choice, VMware.*

This learning lab covers setting up a shared network "drive" using the
very old (but still commonly used) NFS protocol and exporting a single,
secondary drive on a `control` machine that is mounted remotely from the
`worker` machines so that all machines share the same path to the file
systems mounted on that `control` volume.

## But First ...

* Install VMware
* Be able to create VMware virtual machines

## Skills

## Tasks

1. **Create three virtual machines** with 2 GB RAM and 1 CPU each.
   Disable swap on all machines (if you want, although not required).
   Name the machines `control`, and `node1` and `node2`. Create a
   10 MB secondary volume and mount and export it as an NFS volume ***on
   the `control` machine only***. Mount the NFS volume on the two
   worker nodes.

1. Setup NFS server on `control`

1. Setup NFS clients on `node1` and `node2`

1. Setup a mounted/shared `/s` directory and export

Related:

* <https://repo.almalinux.org/almalinux/8/cloud/x86_64/images/>
* Configuring NFS Server on RedHat Enterprise 7  
  <https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/storage_administration_guide/nfs-serverconfig>
* Quick NFS Server configuration on Redhat 7 Linux System - Linux Tutorials - Learn Linux Configuration  
  <https://linuxconfig.org/quick-nfs-server-configuration-on-redhat-7-linux>
* Mounting a disk and creating partitions in Linux \| Instruction manual  
  <https://serverspace.io/support/help/mounting-a-disk-and-creating-partitions-in-linux/>
* How to Setup NFS Server on CentOS 8 / RHEL 8  
  <https://www.linuxtechi.com/setup-nfs-server-on-centos-8-rhel-8/>
* RPC <https://www.guru99.com/remote-procedure-call-rpc.html>
* How to Create Partitions in Linux  
  <https://phoenixnap.com/kb/linux-create-partition>
