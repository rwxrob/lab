# Setup Local Kubernetes Development with Kind, A Learning Lab

[Kind] is a full Kubernetes setup created originally to test Kubernetes
itself from a single machine. It's no surprise it is the favorite among
cloud-native developers and engineers as a sort of local cloud-native
lab in which to experiment and develop new things.

One of the great things about learning to use Kubernetes with Kind is
that, because it is entirely based on Docker, you get even more
experience using Docker skills that you may not have learned before,
such as `docker network`.

## Prerequisites

You will most likely waste your time with this lab unless you have
*already* mastered the following:

* Create and publish container images from Linux bash command line

Since `kind` is often the first thing those learning Kubernetes will
want to install, understanding the following Kubernetes Objects is not
required for this lab before starting it, but after finishing it you
will have at least a basic initial understanding of them:

* Deployment
* Pod
* Node
* Service
* ClusterRole
* ClusterRoleBinding

## Learning Objectives

You will know you have mastered this learning lab when you can
confidently do the following within an hour or so with only minimal
reference to anything in this repo or on the Internet:

* Install and configure local `kind` clusters
* Install and access a Kubernetes web dashboard
* Code a bash script to capture common tasks
* Setup a local private registry in docker
* Deploy a simple `httphey` service locally

## Self Assessment

Even though you will usually be looking up much of these configurations
on the job (no one expects you to memorize all this stuff), it will
serve you well now to memorize it at least once and be able to do it
without looking at anything. 

Repeat the steps as many times as required to commit everything to
memory. Pretend you are giving yourself these steps to perform during a
coding interview with a time limit. Be able to explain each step out
loud as you are doing it. Maybe even stream your final assessment and
make a video showing others how to do it explaining as you go. Helping
someone else learn to do something is one of the best measures of
mastery.

## Steps

1. Create `build` script
   1. Add tab completion support to script
   1. Add `install` command to install `kind` to script
   1. Add `usage` command to script
   1. Add `delete` command to script (optional name)
   1. Add `create` command to build new cluster (optional name)
   1. Add `admin` command to grant admin to all (optional name)
   1. Add `init` command to delete and create new (optional name)
1. Customize `kind.yaml` file
   1. Set the Kubernetes version
   1. Mount `/mnt/kind` to `/files`
   1. Forward port 8000 (will use for `httphey` later)
   1. Use local Docker as private registry
1. Update `build` script
   1. Update `create` to detect `<name>.yaml` file
   1. Add `reg` command to create private registry
   1. Add `regcheck` command to check private registry is up
1. Create and deploy a test deployment (`httphey`)
   1. Pull `rwxrob/httphey` to local docker 
   1. Create `deploy.yaml` Kubernetes Deployment
      1. Add `httphey` Deployment
      1. Add `httphey` Service
      1. Add `login` Deployment (`debian`) to login and check in-cluster
   1. Add `deploy` command to `build`

[Kind]: <https://kind.sigs.k8s.io>

## Also See

* [Quick Start](https://kind.sigs.k8s.io/docs/user/quick-start)
* [httphey](https://github.com/rwxrob/httphey)

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Code released under the [Apache 2.0](LICENSE).
Written material under [Creative Commons Attribution-NonCommercial (CC-BY-NC)](https://creativecommons.org/licenses/by-nc-sa/2.0/).
Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).
