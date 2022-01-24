# Kubernetes the Hard Way in Docker (No Cloud)

This lab is inspired by the popular [Kubernetes the Hard Way] "tutorial"
from Kelsey Hightower, but without the Google Cloud Platform hassle and
dependency using Docker instead to create what we call *host containers*
(or just *hosts*). This lab also draws heavily from from the official
Kubernetes instructions for [Creating a cluster with kubeadm]. In other
words, it is more relevant and will get you certified faster.

[Kubernetes the Hard Way]: <https://github.com/kelseyhightower/kubernetes-the-hard-way>
[Creating a cluster with kubeadm]: <https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>

Host containers are easier and more flexible to setup than host
machines, virtual or real --- especially since we can just mount
`/var/run/docker.sock` into each and everything ends up using a single
Docker runtime engine. This is faster and easier to get through and
focuses on the core learning content. We use the built-in networking
that Docker provides to simulate different bridged networks and even
data centers. This approach is not novel or niche. In fact, this is
*exactly* what both Kind and Minikube do. You'll just be doing the magic
(the hard stuff) yourself.

## First By Hand, Then Create Dockerfile

The suggested way to work through these steps is to do all the
individual tasks one at a time without scripting them at all. This
includes typing out things like `/etc/apt/source.list.d` entries by hand
so that you can internalize what is in them. Later, you should go back
and create your own Dockerfiles for the images that optimize the steps
that you did by hand. After all, this is how Docker was designed to be
used, each step creates an overlay.

## Overview

We take a phased, repetitive approach, adding complexity on each
repetition but starting with just the minimum, a single-node cluster:

Stage One:

1. Create a single node to contain control-plane and worker
1. Install control-plane components
1. Install the worker components
1. Joint the worker to the control-plan (both on same node)
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

**Create a single node to contain control-plane and worker.**

Our hosts will be running Ubuntu (but you could be using any other OS
that you want). Let's create a base host container:

```
docker run -itd -h single --name single ubuntu
```

Make sure you can connect to it.

```
docker exec -it single bash
```

> 💬
> Consider that the login step would normally be done with ssh if this
> Node where a real machine. We use
`docker exec -it` instead because our host is a container (and not
a machine).

We will need to turn this into more of an actual host than is available
from just a container. Let's install some tools we will need. 

First let's update `apt` and do the usual things when dealing with a
container. You might consider putting this into a Dockerfile as you go
so you can save the work later.

```
apt upgrade
apt update
```

Now let's get `vim` and the package necessary to do installation:

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

Let's start with the container engine (Docker). Even though we will be
cheating and using the docker engine running on our workstation by
mounting `/var/run/docker.sock` we still need the `docker` command.
There are many ways to get this, but let's use the official docker
package archive (`docker-ce`).

We'll authorize this package archive source by hand in a way that is compatible with Debian and Ubuntu. 

First, let's get the public key from docker (which is only available as
an ASCII-armored text file).

```
curl -sSL -o /tmp/docker.pub https://download.docker.com/linux/ubuntu/gpg
```

Now we can convert the text file into a binary key file.

```
gpg --dearmor /tmp/docker.pub
```

For `apt` to see this key we need to add it to `/usr/share/keyrings` and
we'll keep with the naming conventions used for other "keyrings" even
though this is really just a single key.

```
mv /tmp/docker.pub.gpg /usr/share/keyrings/docker-archive-keyring.gpg
```

Now we can add the docker package archive to our list of approved `apt`
archive sources. Normally, you should add a `docker.list` file to the
`/etc/apt/sources.list.d` directory, but we'll go ahead and just add it
directly to `/etc/apt/sources.list` for now. Add the following line. You
can use `vi` directly or `echo` with a redirect, but careful not to blow
it away:

```
deb [signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu focal stable
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

```
exit
docker stop single
docker rm single
docker run -itd -h single --name single -v /var/run/docker.sock:/var/run/docker.sock ubuntu
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
Docker apt repo to our source list.

Change into the keyrings directory.

```
cd /usr/share/keyrings
```

Now we `curl` down the binary directly from `https:
//packages.cloud.google.com` (which is organized quite differently than
Docker). The `-O` keeps the original name

```
curl -sSL https://packages.cloud.google.com/apt/doc/apt-key.gpg -O
mv apt-key kubernetes-archive-keyring.gpg
```

Now for then entry in `/etc/apt/sources.list`. Add this however you want.

```
deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main
```

Keep `kubernetes-xenial` for now. (I'm not quite sure on the naming
conventions, but this is from the standard docs.)

Don't forget to `apt update` after this. Then you can install everything
you need in one command.

```
apt update
apt install kubeadm kubectl kubelet
```

Before we begin installing everything for the control plane let's first
check that we can get all the dependent images.

```
kubeadm config images pull
```

> ⚠️
> Kelsey's tutorial (as of this writing) has you `wget` this binaries
> directly and copy them into `/usr/local/bin`. I believe this is a bad
> practice (even an anti-pattern) that you should never do. Not only
> does it completely ignore the fact that we have added the official
> Kubernetes package archive from which we should be doing `apt install`
> for things, but it completely ignores the fact that these binaries are
> now contained within containers instead of binaries. This shows how
> out of date that "tutorial" has become.

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
