# Learning Lab: Create Kubernetes Services (Solution)

> ⚠️
> This is a *solution* to the learning lab that is kept separate in the
> "solution" branch on purpose. Be careful never to merge it into the
> main trunk. Other solutions, of course, are expected to be posted by
> others in their own ways. This is just the solution created at the
> time the lab was created.

## Create a Kind cluster with three worker nodes

Create a `kind.yaml` file.

```yaml
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
name: lab-k8s-services
nodes:
  - role: control-plane
  - role: worker
  - role: worker
  - role: worker
```

Run `kind` to create the cluster pointing to `kind.yaml` file.

```bash
kind create cluster --config kind.yaml
```

```out
Creating cluster "lab-k8s-services" ...
 • Ensuring node image (kindest/node:v1.21.1) 🖼  ...
 ✓ Ensuring node image (kindest/node:v1.21.1) 🖼
 • Preparing nodes 📦 📦 📦 📦   ...
 ✓ Preparing nodes 📦 📦 📦 📦 
 • Writing configuration 📜  ...
 ✓ Writing configuration 📜
 • Starting control-plane 🕹️  ...
 ✓ Starting control-plane 🕹️
 • Installing CNI 🔌  ...
 ✓ Installing CNI 🔌
 • Installing StorageClass 💾  ...
 ✓ Installing StorageClass 💾
 • Joining worker nodes 🚜  ...
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-lab-k8s-services"
You can now use your cluster with:

kubectl cluster-info --context kind-lab-k8s-services

Have a nice day! 👋
```

## Create a TCP server for use in creating Services

<https://github.com/rwxrob/gotcptest>

## Create a namespace

Make sure you have the correct `context` set.

```bash
kubectl config get-contexts
```

Because we need the namespace to be created before we can even do
`--dry-run=server` creations of other resource, let's just do `kubectl
create ns lab-k8s-services` and then just dump that into a file.

```bash
kubectl create ns lab-k8s-services -o yaml >> k8s.yaml
```

I plan on explicitly listing the `namespace` in every resource manifest
so I capture it into the beginning of a file. We put the `namespace`
first so that the other resources in the same YAML file can use that
namespace. That way, if the namespace is set to something else, it still
does what is expected (at the cost of remembering to update the
`namespace` properties anytime one is changed.

This produces a `k8s.yaml` file that needs some clean up. 

```yaml
apiVersion: v1
kind: Namespace
metadata:
  creationTimestamp: "2021-12-23T21:55:16Z"
  labels:
    kubernetes.io/metadata.name: lab-k8s-services
  name: lab-k8s-services
  uid: d75ca20a-36da-41a7-97ec-28731e48ad8e
spec:
  finalizers:
  - kubernetes
status:
  phase: Active
```

The `spec`, and `status`,`metadata.{creationTimestamp,uuid}` can be
safely removed. The `spec.finalizers` is there to help avoid removing
something that you want to keep, but would block in "Terminating" status
so you usually would remove it.

We'll leave the labels since they don't hurt and are sort of a standard
Kubernetes thing these days --- especially for non-namespaced resources.

```yaml
apiVersion: v1
kind: Namespace
metadata:
  labels:
    kubernetes.io/metadata.name: lab-k8s-services
  name: lab-k8s-services
---
```

The `---` at the end is so that the next resource appended will be
separated.

Now we need to set the default namespace for everything to make it
easier to create the rest.

```bash
kubectl config set-context --current --namespace lab-k8s-services
kubectl config get-contexts
```

```out
Context "kind-lab-k8s-services" modified.
CURRENT   NAME                    CLUSTER                 AUTHINFO                NAMESPACE
*         kind-lab-k8s-services   kind-lab-k8s-services   kind-lab-k8s-services   lab-k8s-services
```

## Create four different Deployments with three replicas

All deployments will use the `rwxrob/gotcptest` image with a different
name for each (`NAME=other`).

## Create a ClusterIP Service for one of Deployments

```bash
kubectl create deployment testclusterip --image=rwxrob/gotcptest --dry-run=server -o yaml >> k8s.yaml
```

Note that this one requires an `--image` or it won't work.

We do not want to create an actual deployment yet. That is what
`--dry-run=server` will do. It creates YAML output and validates the
creation against the current Kubernetes server API. (The
`--dry-run=client` will not.) Then, we will making changes to the
default deployment YAML (such as increasing the number of replicas).

The output, again, contains stuff that can be removed (and altered).

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: "2021-12-23T22:13:47Z"
  generation: 1
  labels:
    app: testclusterip
  name: testclusterip
  namespace: default
  uid: b7b3688c-da0a-43f2-b3ad-8f53c1478949
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: testclusterip
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: testclusterip
    spec:
      containers:
      - image: rwxrob/gotcptest
        imagePullPolicy: Always
        name: gotcptest
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status: {}
```

Remove all the `uuid`, `createTimestamp`, `generation` properties and the `status` section.

Change `spec.replicas` to 3.

## Create a NodePost Service for one of Deployments

## Create a LoadBalancer Service for one of Deployments

## Create a ExternalName Service for one of Deployments

## Create a DaemonSet with exposure as a service 

## Related

* Service \| Kubernetes  
  <https://kubernetes.io/docs/concepts/services-networking/service/>
* Learning Lab: Create TCP Service in Go  
  <https://github.com/rwxrob/lab-go-tcp-server>
* Learning Lab: Setup Local Kubernetes with Kind  
  <https://github.com/rwxrob/lab-k8s-kind>
