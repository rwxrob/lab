terraform {
  required_providers {
    libvirt = {
      source = "dmacvicar/libvirt"
    }
  }
}

provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_volume" "myvol" {
  name   = "ubuntu2010"
  source = "https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img"
  format = "qcow2"
}

resource "libvirt_domain" "myvm" {
  name   = "vmthing"
  memory = "2048"
  vcpu   = 2
  network_interface {
    network_name = "default"
  }
  disk {
    volume_id = libvirt_volume.myvol.id
  }
  console {
    type        = "pty"
    target_type = "serial"
    target_port = 0
  }
  graphics {
    type        = "spice" # just for drag-n-drop support
    listen_type = "address"
    autoport    = true
  }
}

output "ip" {
  value = libvirt_domain.myvm.network_interface.0.addresses.0
}
