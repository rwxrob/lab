# Learning Lab: Use Terraform for Local Development

> 🤬
> FAIL: You can't have Terraform unless you have a cloud to form! I gave
> up on this learning lab because it is literally impossible to use
> Terraform unless you have more than one actual machine and they are
> managed by some soft of cloud management software. I moved to setting
> up a Proxmox cluster and then managing that with Terraform instead.
> When working with stuff on the same system Vagrant was the option, but
> it is also really, really bad when using QEMU and is based on ancient
> tech. Therefore, I will stick with bash and `qemu` commands for
> bringing up local environments for testing and using Docker workspace
> containers when possible.

The goal is to be able to emulate complete IT architectures quickly and
easily on a single computer with 8 or more cores and 16GB of RAM or
more, to create an "IT shop in box" to test things like multi-cluster
Kubernetes deployments, cluster + stuff, etc. The sort of things that
you might want to *really* break during experimentation.

We are essentially using Terraform instead of the (old and busted)
Vagrant (from the same company).

1. Install Terraform using the most modern methods. (Don't forget to
   `exec bash -l` to refresh your shell and get the tag completion
   working or use the super-secret, completely undocumented `terraform
   -install-autocomplete`)

1. Provision a single VM using `libvirt` on the local machine with nothing
   more than Terraform.

1. Test using Docker as a *provider*.

1. Become aware of what is sensitive and what is not. Be sure not to
   save files with sensitive data and know what that means (`.tf` is
   okay, `.tfstate` and `.tfplan` not so much). Consider `.gitignore`.

1. Become familiar withe the essentials of the Terraform
   Configuration Language.

1. Create a `.tf` file that provisions three VMs for a Kubernetes
   control plane and five worker nodes. (But leave out all the
   configuration. We'll do that in the Ansible lab.)

1. Have Terraform output the Ansible inventory/host file after applying.

1. Create a reusable Terraform module.

## Gotchas

Watch out for the permissions on the `default` pool. You'll probably
want to create an *additional* pool that you have full permissions so
that when `tf apply` is called it will download the image to the
directory that can later be properly accessed. Such is not the case with
`/var/lib/libvirt/images` (by default).

Related:

* <https://github.com/rwxrob/lab-qemu-setup>
* Introduction \| Terraform by HashiCorp  
  <https://www.terraform.io/intro>
* Use Ubuntu Server Cloud Image to Create a KVM Virtual Machine with Fixed Network Properties \| by Yu Ping \| Medium  
  <https://yping88.medium.com/use-ubuntu-server-20-04-cloud-image-to-create-a-kvm-virtual-machine-with-fixed-network-properties-62ecae025f6c>
* Overview - Configuration Language \| Terraform by HashiCorp  
  <https://www.terraform.io/language>
* How To Provision VMs on KVM with Terraform \| ComputingForGeeks  
  <https://computingforgeeks.com/how-to-provision-vms-on-kvm-with-terraform/>
* Providers - Configuration Language \| Terraform by HashiCorp  
  <https://www.terraform.io/language/providers>
* <https://registry.terraform.io/browse/providers>
* <https://github.com/dmacvicar/terraform-provider-libvirt>
* <https://learn.hashicorp.com/terraform>
* <https://github.com/Telmate/terraform-provider-proxmox>
* <https://cloudinit.readthedocs.io/en/latest/> 

Next Steps:

* Learning Lab: Install Local Kubernetes with `kubeadm`
