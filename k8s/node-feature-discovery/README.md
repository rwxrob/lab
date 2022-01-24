# Node Feature Discovery in Kubernetes

This lab is an exploration of [node feature discovery](https://lite.duckduckgo.com/lite?kd=-1&kp=-1&q=node%20feature%20discovery), an official
project under [CNCF](https://lite.duckduckgo.com/lite?kd=-1&kp=-1&q=CNCF), which allows information about
[Kubernetes node machines](https://lite.duckduckgo.com/lite?kd=-1&kp=-1&q=Kubernetes%20node%20machines) to be queried easily from [Kubernetes
Pods]() and [containers](https://lite.duckduckgo.com/lite?kd=-1&kp=-1&q=containers).

## Observations

* Change `core.labelWhiteList` to regular expression to filter

## Steps

1.  Add the chart repo with helm
2.  Download the chart tarball
3.  Dump the default values.yaml file
4.  Do a dry-run install and save to dry.yaml
5.  Verify that dry.yaml is what is wanted
6.  Modify values.yaml as needed
7.  Repeat step 4
8.  Look at environment

