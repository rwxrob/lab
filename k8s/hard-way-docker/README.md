# Kubernetes the Hard Way in Docker (No Cloud)


THIS IS CURRENTLY WRONG AND BASED ON OLD `apt-key` INFO

This lab is inspired by the popular [Kubernetes the Hard Way] "tutorial"
from Kelsey Hightower, but without the Google Cloud Platform hassle and
dependency using *host containers* instead of virtual of real machines.
This lab also draws heavily from from the official Kubernetes
instructions for [Creating a cluster with kubeadm]. In other words, it
is more relevant and will get you certified faster.

> ⚠️
> This lab runs you into the wall on purpose to make a point of learning
> it "the hard way" (exactly like you would if doing this from scratch
> on your own). This will seem annoying since it deliberately forces you
> to repeat work you've already done, a lot. If you do not wish to redo
> work to learn it with repetition you can read through the entire lab
> and then decide how to navigate it.

[Kubernetes the Hard Way]: <https://github.com/kelseyhightower/kubernetes-the-hard-way>
[Creating a cluster with kubeadm]: <https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>

## First By Hand, Then Create a Dockerfile

The suggested way to work through these steps is to do all the
individual tasks one at a time without scripting them. This includes
typing out entries by hand so that you can internalize what is in them.
Later, you should go back and create your own Dockerfiles for the images
that optimize the steps you did. After all, this is how Docker was
originally designed to be used, each step creating an overlay.

## Prerequisites

This lab assumes you have a good handle on the following and that you
are working on a Linux workstation. 

> 🐧 ⚠️
> Seriously, just get a Linux machine. What are you waiting for? This
> stuff may work on a Mac or Windows machine, but many steps will be
> different.

* Linux Bash Command Line
* Git and GitHub for Source Management
* CommonMark Markdown
* Bash Scripting
* JSON / YAML
* Go Template Language
* Web Essentials (HTML, CSS, Plain JavaScript, DOM, HTTP)
* Networking and Network Troubleshooting
* Using and Creating Containers with `docker` (or `podman`) and `Dockerfile`
* (There is no need to learn Docker compose.)
* Have `podman` Installed and Working

## Overview

We'll take a phased, repetitive approach, adding complexity on each
repetition but starting with just the minimum, a single-node cluster:

Stage One:

1. Create a single node for both the control-plane and worker
1. Install control-plane components
1. Install the worker components
1. Join the worker to the control-plane (both on same node)
1. Test basic Kubernetes functionality

Stage Two:

1. Create two bridged networks: `control-net` and `worker-net`
1. Create `control-1` control node host on `control` network
1. Create `worker-1` worker node host on `workers` network
1. Install control-plane components on `control-1`
1. Install worker components on `worker-1`
1. Join `worker-1` to the cluster
1. Test basic Kubernetes functionality

## Stage One

**Create a host container using `docker`.**

Our hosts will be running Ubuntu (but you could be using any other OS
that you want). I personally recommend sticking with one of the major
enterprise-y distros (RedHat, Centos, Fedora, Suse, Ubuntu). 

Let's create a base host container:

```
docker run -itd -h single --name single ubuntu
```

Make sure you can connect to it.

```
docker exec -it single bash
```

> 💬
> Consider that the login step would normally be done with ssh if this
> Node where a machine. We use `docker exec -it` instead because our
> host is a container.

We will need to make this *host container* more like an actual host.
Let's install some tools we need. 

First let's update `apt` (just like you would with a host machine).


```
apt upgrade
apt update
```

> 💬
> Eventually, you will want to put these steps into a Dockerfile and
> create an image that we can run easily.

Now let's get `vim`, `curl`, and `gpg` (for adding package archives).

```
apt install vim-tiny curl gnupg
```

Now we are ready to install the three components of any Kubernetes Node:

1. Container Engine (Docker)
1. Kubelet
1. KubeProxy

These logical components are fulfilled by install the following packages:

1. `docker-ce`
1. `kubeadm`
1. `kubectl`
1. `kubelet`

Let's start with the container engine (Docker Community Edition). Even
though we will be cheating and using the docker engine running on our
workstation by mounting `/var/run/docker.sock` we still need the
`docker` command. There are many ways to get this, but let's use the
official docker package archive (`docker-ce`).

We'll authorize this package archive source by hand in a way that is
compatible with Debian and Ubuntu. 

First, let's get the public key from docker which is only available as
an ASCII-armored text file so we'll have to convert it before adding it.

```
curl -sSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor | apt-key add -
```

Now we can add the docker package archive to our list of approved `apt`
archive sources. Normally, you should add a `docker.list` file to the
`/etc/apt/sources.list.d` directory, but we'll go ahead and just add it
directly to `/etc/apt/sources.list` for now. 

Add the following line. You can use `vi` directly or `echo` with a redirect,
but careful not to blow it away:

```
deb https://download.docker.com/linux/ubuntu focal stable
```

Make sure `focal` and `stable` match your release. When in doubt you can
open the web page at <https://download.docker.com/linux/ubuntu> to get
hints about what is available. The `dpkg --print-architecture` command
can also help. 

> 💡
> Adding a source list directly in this way is 100% compatible with
> Debian as well as Ubuntu and saves you from the extra bloat and
> dependency on `apt-add-repo` (which is only support on Ubuntu). It
> pays to know the Debian way of doing things.

Now just update your apt repository cache and install the `docker-ce`
package for Docker Community Edition. 

```
apt update
apt -y install docker-ce
```

> ⚠️
> Note that the kubernetes.io documentation calls for `docker-ce-cli`
> and `containerd.io` packages, which is incorrect since they will also
> be installed with `docker-ce`.

Take your new `docker` command for spin to test it out.

```
docker version
```

Note the error is produces at the bottom.

```
docker version
Client: Docker Engine - Community
 Version:           20.10.9
 API version:       1.41
 Go version:        go1.16.8
 Git commit:        c2ea9bc
 Built:             Mon Oct  4 16:08:29 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
```

To get around this error we will exit, stop, and remove our container
creating a new one with a new `run` command and the volume mount
connecting the `/var/run/docker.sock` of the host machine. This is
a regular practice to allow accessing the same docker runtime engine
from containers as the host of the containers.

> 🛑
> This will completely discard all the work you have done so far. So
> you will have to repeat all the previous steps. This is
> intentional. You need the practice. This time try to do them all
> without looking at your notes too much. You did take notes, right? If
> you wish (and know how), you can begin to create your own Dockerfile
> image with notes and start using that instead.

While we are at it, let's mount the `/sys/fs/cgroup` file in read-only
mode so that we can install `systemd` (which we will need for `kubeadm`
to work).

```
exit
docker stop single
docker rm single
docker run -itd -h single --name single -v /sys/fs/cgroup:/sys/fs/group:ro /var/run/docker.sock:/var/run/docker.sock ubuntu
```

You can see from the `inspect` output that a new volume has been bound.

```
docker inspect single |jq '.[0].Mounts'
```

Now "login" to your node again.

```
docker exec -it single bash
```

Your test of `docker` should now work. Keep in mind the output will look
exactly like it did when you were not logged in because the `docker`
inside the container is using exactly the same `docker` as you installed
on your workstation. This is intentional.

```
docker version
docker ps
```

Your output should look something like this:

```
Client: Docker Engine - Community
 Version:           20.10.9
 API version:       1.41
 Go version:        go1.16.8
 Git commit:        c2ea9bc
 Built:             Mon Oct  4 16:08:29 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server:
 Engine:
  Version:          20.10.7
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.8
  Git commit:       20.10.7-0ubuntu1~21.04.1
  Built:            Wed Aug  4 12:24:19 2021
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.5.2-0ubuntu1~21.04.2
  GitCommit:        
 runc:
  Version:          1.0.0~rc95-0ubuntu1~21.04.2
  GitCommit:        
 docker-init:
  Version:          0.19.0
  GitCommit:        

CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS          PORTS     NAMES
08864d48d710   ubuntu    "bash"    54 minutes ago   Up 54 minutes             single
```

Huzzah! Docker is now installed on the node. 

> 💬
> Note that the [kubernetes.io][node-components] site does not include
> the container runtime engine in the diagram describing the essential
> components of a Kubernetes Node. This is probably because the
> container engine is technically not a part of Kubernetes itself, but a
> primary dependency for Kubernetes. In other words, you could always
> login to your Node and just use docker on it as if it were any other
> system with docker installed. 
> 
> In fact, even though it is a very bad practice, cloud-native admins
> will sometimes login to a Node machine just to use docker to
> trouble-shoot --- or worse --- to build images when they are too lazy
> to build them on their own workstations. Keep in mind, doing this is
> *very* risky since any mistake could leave Kubernetes and the Kubelet
> in a unusable state. 
>
> Hackers target the Docker scope issue to gain access to the underlying
> system by breaking out of Kubernetes and gaining access to the
> underlying host system, which is why the Docker engine should *never*
> run as root. Currently, this is the most popular attack vector against
> Kubernetes.

[node-components]: <https://kubernetes.io/docs/concepts/overview/components/>

Okay, now we need to install the Kubernetes specific stuff: `kubeadm`,
`kubectl`, and `kubelet`. We'll use similar steps to those to add the
Docker apt repo to our source list, but because the key is provided in
binary format we can use `apt-key` instead to add it.

Download and add the key to the system keyring. Then list the keys and
look for a new one from Google to verify.

```
curl -sSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
apt-key list
```

Now, for the entry in `/etc/apt/sources.list`. Add this however you want.

```
deb https://apt.kubernetes.io/ kubernetes-xenial main
```

Keep `kubernetes-xenial` for now. (I'm not quite sure on the naming
conventions, but this is from the standard docs.)

Don't forget to `apt update` after this. 

```
apt update
```

Then you can install everything you need in one command.

```
apt install kubeadm kubectl kubelet
```

> 💬
> Note that there is no `kube-proxy` here (which is an actual binary).
> The reason is because we are going to use `kubeadm` to set everything
> up and it will add the `kube-proxy` container instead as a DaemonSet
> Kubernetes resource bound to this host. Technically speaking, the
> `kubeadm` and `kubectl` commands are not needed on every single node,
> but it is convenient to include them --- especially when creating
> a base image for host containers.

Before we begin installing everything for the control plane let's first
check that we can get all the dependent images.

```
kubeadm config images pull
```

> ⚠️
> Kelsey's tutorial (as of this writing) has you `wget` this binaries
> directly and copy them into `/usr/local/bin`. I believe this is a bad
> practice, an anti-pattern that would fail most employability
> verifications, that you should never do. Not only does it completely
> ignore the fact that we have added the official Kubernetes package
> archive from which we should be doing `apt install` for things, but it
> completely ignores the fact that these binaries are now contained
> within containers instead of binaries.

Okay now for `kubeadm init`, finally.

```
kubeadm init
```

Which should produce something like this output.


## Troubleshooting

"What are all these errors? And what's a 'preflight'?"

The "preflight" phase is the first phase that `kubeadm` runs to check
everything and make sure the rest of the phases will work.

Most of the errors are because we are using a host container and not a
virtual or real machine. 


```
W1021 12:03:53.898229    7228 kubelet.go:209] cannot determine if systemd-resolved is active: no supported init system detected, skipping checking for services
[init] Using Kubernetes version: v1.22.2
[preflight] Running pre-flight checks
        [WARNING Firewalld]: no supported init system detected, skipping checking for services
        [WARNING Service-Docker]: no supported init system detected, skipping checking for services
[preflight] The system verification failed. Printing the output from the verification:
KERNEL_VERSION: 5.11.0-7633-generic
DOCKER_VERSION: 20.10.7
DOCKER_GRAPH_DRIVER: overlay2
OS: Linux
CGROUPS_CPU: enabled
CGROUPS_CPUACCT: enabled
CGROUPS_CPUSET: enabled
CGROUPS_DEVICES: enabled
CGROUPS_FREEZER: enabled
CGROUPS_MEMORY: enabled
CGROUPS_PIDS: enabled
CGROUPS_HUGETLB: enabled
        [WARNING Service-Kubelet]: no supported init system detected, skipping checking for services
error execution phase preflight: [preflight] Some fatal errors occurred:
        [ERROR Swap]: running with swap on is not supported. Please disable swap
        [ERROR SystemVerification]: failed to parse kernel config: unable to load kernel module: "configs", output: "modprobe: FATAL: Module configs not found in directory /lib/modules/5.11.0-7633-generic\n", err: exit status 1
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
To see the stack trace of this error execute with --v=5 or higher
```

```
W1021 12:03:53.898229    7228 kubelet.go:209] cannot determine if systemd-resolved is active: no supported init system detected, skipping checking for services
```

When we lookup that error in the source code
(`kubernetes/cmd/kubeadm/app/util/initsystem/initsystem_unix.go`) we see that
only `systemd` and `openrc` are supported "init systems" (even though this
isn't documented anywhere). That's right `kubeadm` doesn't care about our
`--init` and just ignores it. And since `openrc` is not supported at all on
Ubuntu we are left with the rather annoying task of installing `systemd` (or
using `podman` instead of `docker` which automatically installs `systemd` into
*any* running container). Either way, we have to stop and delete our running
container and start over again. At this point, we'll replace the `docker
run` with `podman run`, but first we have to install `podman`.

        [ERROR Swap]: running with swap on is not supported. Please disable swap

```

```




TODO figure out how to install `systemd` into the host container. This
is a standard part of the official documentation regarding the services
used by the control plane. Not having it makes any lab somewhat
worthless. Look to how Kind accomplishes this without risking the host
system. Unfortunately, this step is something that would only have to be
done with using Docker host containers as we are. This is one step
that learning approaches using virtual or real machines would not
require (but the other time-consuming steps they *do* require easily
exceed the time for this work, besides learning how to get
`systemd` in a container opens tons of systems engineering and
development and development use cases that don't have anything to do
with Kubernetes).

----

TODO run the Kubernetes conformance tests against our own cluster.







TODO Rethink where the following should go.


## Steps

**Create two Docker networks** one for the control plane (control nodes)
and one for the workers (worker nodes).

```bash
docker network create k8s-control
docker network create k8s-workers
```

* docker network create \| Docker Documentation  
  <https://docs.docker.com/engine/reference/commandline/network_create/>

**Create some control nodes for control plane.**

```bash
for i in {1..3};do
  docker run -itd \
    --network=k8s-control \
    --name control-$i \
    -h control-$i \
    ubuntu
done
```

**Create some workers.**

```bash
for i in {1..3};do
  docker run -itd \
    --network=k8s-workers \
    --name worker-$i \
    -h worker-$i \
    ubuntu
done
```

**Note the IP addresses of all the nodes (masters and workers)** which
will be needed when creating the certificates.

```bash
mapfile -t names < <(d ps --format '{{.Names}}')
for n in "${names[@]}"; do
  echo -n "$n "
  d inspect --format '{{json .}}' $n \
    | jq -r '.NetworkSettings.Networks[].IPAddress'
done
```

**Test access to masters and workers** using `docker exec -it` instead
of `ssh` to access these systems.

```bash
docker exec -it worker-1 bash
```

Should dump you onto a prompt:

```
root@worker-1:/#
```

**Install dependencies on all host containers, workers and .**[



> ⚠️
> At this point you can choose either to create certificates using your
> own Certificate Authority, or you can just use `kubeadm` and have your
> certificates managed with Kubernetes.

## Next Steps

**Simulate two master control planes in different data centers
(networks).**

## Alternate Paths

**Provision a Certificate Authority** using the original, standard,
ubiquitous `openssl` tool (there are others you can learn later).

**Generate CA key** which will be used to create and sign all the other
keys and create certificates.

```bash
openssl genrsa -o ca.key 2048
```

**Generate CA X.509 certificate** from the key. `req` is from the
term Certificate Signing Request (CSR). We'll use 10,000 days since we
don't need to worry about certificates expiring, but in real scenarios
you should setup a reminder immediately for before the cert expires to
be sure it gets replaced before then (or very bad things will happen).

```bash
openssl req -x509 
```
