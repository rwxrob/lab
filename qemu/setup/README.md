# Learning Lab: Setup QEMU Virtual Machine on Linux

> Qemu is short for "quick emulator"

Sometimes you just want a nice, safe place to try something really
crazy (like over-allocating a PV in Kubernetes) in order to see if it
will fail. This is an ideal situation for QEMU, which runs on your
favorite operating system without all the extra weight.

> ⚠️
> Don't forget to start `virt-manager` from the command line with `sudo`
(root) permissions if it does not prompt for a password.

## To Run as User ...

You can add yourself to the libvirt group. and uncomment
"unix_sock_group= libvirt" in /etc/libvirt/libvirtd.conf. to avoid
launching as root

*(Thanks to @malucious from Twitch community.)*

`sudo apt remove snapd` is the first thing I did (based on suggestions).

```
sudo virsh list --all
```

I added the IP to ssh and started using that from command line.

## From Command Line

1. Get the image (precisely the "cloud server" version).

```
curl -sSLO "https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img"
```

1. Allocate a disk volume with the image in it.

```
qemu-img create -F qcow2 -b focal-server-cloudimg-amd64.img -f qcow2 vm01.qcow2 10G
```

Related:

* Qemu tutorial  
  <https://linuxhint.com/qemu-tutorial/>
* Qemu/KVM Virtual Machines - Proxmox VE  
  <https://pve.proxmox.com/wiki/Qemu/KVM_Virtual_Machines>
* Can you run QEMU on Windows? -- Bridgitmendlermusic.com  
  <https://bridgitmendlermusic.com/can-you-run-qemu-on-windows/>
* <https://powersj.io/posts/ubuntu-qemu-cli/>
* <https://leftasexercise.com/2020/05/18/managing-kvm-virtual-machines-part-ii-the-libvirt-toolkit/>
* Setup Virtual Machine using QEMU in Ubuntu - TechPiezo  
  <https://techpiezo.com/linux/setup-virtual-machine-using-qemu-in-ubuntu/>
* QEMU - ArchWiki  
  <https://wiki.archlinux.org/title/QEMU>
* KVM - ArchWiki  
  <https://wiki.archlinux.org/title/KVM>
