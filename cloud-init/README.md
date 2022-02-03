# Create Variable VMware VMs with `cloud-init` Script

The goal of this lab is to arrive at a place where you can create a
simple script that will generate an ISO file that can be dropped into
any virtual machine image that supports the `cloud-init` standard. We'll
use Python so that the code coming from it can be integrated into
Ansible eventually.

> ⚠️
> When working on WSL2 (and using the `exe` equivalents, or aliases to
> such) note that stuff will be created in the default Windows home
> directory of the current user (usually `/mnt/c/Users/<user>`). This
> can be confusing because you might have specified the current working
> directory (which technically that is for `exe` commands).

## But First ...

* Install VMware Workstation Pro and add to your PATH2

## Tasks

1. **Write a script to create `cloudinit.iso` file**


1. Download the lastest Alma8 cloud qcow2 image (check cache)
1. Convert the qcow2 image to vmdk
1. Create a 10MB volume for NFS testing
1. Create VM: control
   1. Create the `cloudinit.iso`
      1. 2 CPU
      1. 2 GB RAM
      1. Static IP 
      1. SSH authorized_hosts
      1. Mount volume in `/s`
1. Create VMs: node1, node2
   1. Create the `cloudinit.iso`
      1. 1 CPU
      1. 2 GB RAM
      1. Static IP 
      1. SSH authorized_hosts

Related:

* <https://gist.github.com/lbogdan/11a3df688bb238698556127cdeea6034>
* <https://github.com/beeman/createvm.sh>
