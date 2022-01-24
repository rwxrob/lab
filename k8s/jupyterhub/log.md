## Tuesday, September 14, 2021, 11:48:10AM EDT

* Reworked GitHub repo, `README.md`, and began first video log of work.

## Tuesday, September 14, 2021, 12:51:32PM EDT

* Wrote the scenario and objectives paragraph.
* Removed old `build` script
* Add `minikube` completion to `~/.bashrc`
* Copy `template-bash-command` at `build`
* Ripped out `config` from `build`
* Add the chart repo to to `helm`
* Create a `values.yaml` empty file
* Install the Helm 1.0.1 chart as is.
* Disable *all* `prePullers`
* Updated `install` to `apply`
* Use `netshoot` to look around
* Disable `hub.networkPolicy` (just to keep similar to work)
* Get access to the web interface

## Wednesday, September 15, 2021, 12:21:06PM EDT

* Login using local login (`jovyan:jupyter`)
* Change default from 'tree' to 'lab' view
* Change to run in `jhub` namespace (instead of default)
* Remembered user pods are not cleaned up with `helm uninstall`
* Discovered `helm get manifest | k get -f -`
* Added `extraConfig` Python for spawning into user's namespace
* Fought with `hub` ServiceAccount missing permissions

## Thursday, September 16, 2021, 11:11:52AM EDT

* Discovered how to use `mk -p` properly
* Cleaned up errant minikube clusters (`mk delete -p ...`)
* Create `jupyterhub` ClusterRole
* Add `namespaces:create` and `persistentvolumeclaim:create` to role
* Create `jupyterhub` ClusterRoleBinding for `hub`
* Create a mock `imauser` and `jovyan` user namespaces
* Add `nstype=user` label to namespaces
* Add `uid=USERUNIXID` to annotations (not labels)
* Get `.singleuser.profileList` working
* Realized only one user server at a time is allowed
* Tests of plain "Ubuntu" container failed
* Realized containers have specific JupyterHub requirements
* Images *must* include "OCI runtime"
* Images *must* include `jupyterhub-singleuser` executable
* Realized failed images cause zombie state, can't stop server
* Consider REST API to kill zombie images for developers
* Running in user namespace allows user to recover from zombie
* Hung zombie images will eventually clear out, but delay
* Confirmed home directory persistence across image types

Error created by plain "Ubuntu" image:

```
2021-09-16T18:55:44Z [Warning] Error: failed to start container "notebook": Error response from daemon: OCI runtime create failed: container_linux.go:370: starting container process caused: exec: "jupyterhub-singleuser": executable file not found in $PATH: unknown
```

## Friday, September 17, 2021, 8:37:47AM EDT

* User images must be derived from standard images
* Running as user ID will require `postStart.command`
* Use `curl` `postStart.command`, most sustainable
* Helm chart unaffected by `postStart` script changes (if `curl`) 
* Division of work:
  1. Helm chart
  1. template customization
  1. `postStart` script (git repo, etc.)
* Keep `joyvan` user, just add if custom user detected
* Detect custom user if `/s/USER/jupyterhub` exists
* Images are tightly bound to profiles
* Must add profile for additional image and bind to it
* Cannot alter template to allow image names
* Template customization blow out `profileList`
* Template changes amount to simply other ways of `profileList`
* Recommend sticking to isolated number of images

## Monday, September 20, 2021, 9:44:28AM EDT

* ELK = Elastic + Logstash + Kibana
* EFK = Elastic + Fluentd + Kibana
* EFG = Elastic + Fluentd + Graylog
* ECK = Elastic Cloud on Kubernetes (Operator)
* Labels can only have 63 characters
* Decided to focus on pulling Git images for notebooks
* Reminder: always `helm ... --namespace` 
* Reminder: always `k config set-namespace --namespace NS`
* Enable registry plugin in Minikube
* Minikube plugins are specifically assigned to profiles
* Decide who creates and updates images (`.singleuser.profileList`)?
* Decided to mirror all standard JupyterHub image *only*
* Decided sufficient reason to keep user namespace pods
* Decided mounting NAS not needed (or perhaps even possible)
* Decided focus on public (internal) git pulls
* How do we manage reclaiming PVs after users stop using?
* Learned Minikube uses <https://gcr.io/k8s-minikube/kicbase> image
* Minikube hard-coded to use specific ISO VM image  
  <https://storage.googleapis.com/minikube/iso/minikube-v1.2.0.iso>  
  (`~/.config/profiles`)

## Tuesday, September 21, 2021, 12:40:03PM EDT

* Compile a list of standard JupyterHub ('golden') images
* Derive image from <https://github.com/jupyter/docker-stacks>
* Documentation at <https://jupyter-docker-stacks.readthedocs.io>
* Discovered that only `vim-tiny` is on images
* Adjusted my `.vimrc` to detect with `if has("eval")`

## Wednesday, September 22, 2021, 9:01:46AM EDT

* Add full `.singleuser.profileList` entries for each standard image
* Add environment variables to profiles without affecting them
* Tried `http_proxy` for user but it interferes
* Add `c.JupyterHub.allow_named_servers = True`

## Thursday, September 23, 2021, 11:28:45AM EDT

* How to change default namespace on minikube cluster?  
  (`~/.minikube/profiles/jhub/config`)
* Got `c.Kubespawner.profile_list` function ('Callable') to work
* Discussed future of everything-is-a-controller

## Tuesday, September 28, 2021, 1:06:28PM EDT

* Setting `c.KubeSpawner.image` changes image
* Setting `c.KubeSpawner.options_form` overrides form input
* Use `c.KubeSpawner.options_from_form` to handle form data
* ***Only*** the `profile` field name is allowed (propagated)
* Use `c.KubeSpawner.pre_spawn_hook` to set `image` from `profile`
* Horrible hack that needs to be reported
* <https://github.com/jupyterhub/jupyterhub/blob/HEAD/examples/spawn-form/jupyterhub_config.py>
* <https://github.com/jupyterhub/kubespawner/pull/389/commits/5a9651c4f30d9ae96676937af1ac0e8082e6c4ed> (broke `options_form`)

The following error will show up in the logs for hub, but it is okay
since we are using the JavaScript to stuff everything into the `profile`
approved field. (This is a very annoying bug in the JupyterHub code
itself.)

```
* [W 2021-10-14 16:36:39.344 JupyterHub spawner:2811] Ignoring unrecognized KubeSpawner user_options: cpu_guarantee, image, mem_guarantee
```

## Wednesday, September 29, 2021, 1:25:23PM EDT

* Expanding the HTML form for "profile"

## Thursday, September 30, 2021, 12:15:13PM EDT

* Add JS function to combine all form data into `profile` string value
* Add Python function to unpack the `profile` string on back end 

## Monday, October 4, 2021, 11:16:54AM EDT

* Research KubeSpawner values for CPU, RAM constraints.
* Added `cpu_guarantee` and `mem_guarantee` to form fields

## Tuesday, October 12, 2021, 11:05:12PM EDT

* Decided to wait on testing existence of image in registry
* Attempt to get Pytorch JupyterHub NVIDIA image working
* Pytorch container says "no derivative works" but ...
* Pytorch image has examples of how to extend meaning ...
* You cannot publish your customizations which explains why ...
* There is no Pytorch image in the docker-stacks collection
* Minikube lab env bork from lack of disk
* Decided to port all of lab into Kind instead
* Annoyed that Minikube uses *additional* container engine

## Thursday, October 14, 2021, 11:49:31AM EDT

* Hit wall with Kind port being dumb about service mappings
* Wrote 20211014151743 about Minikube v.s. Kind re networking
* Decided to pretty much never use Kind again
* Did `docker pull nvcr.io/nvidia/pytorch:21.04-py3`
* Discovered that `mk --driver=docker` (the default) fixes disk issue
* Started work on a PyTorch for JupyterHub Dockerfile and README.md

Here's the errors from a completely unmodified PyTorch image:

```
2021-10-14T16:36:40Z [Normal] Container image "jupyterhub/k8s-network-tools:1.0.1" already present on machine
2021-10-14T16:36:40Z [Normal] Created container block-cloud-metadata
2021-10-14T16:36:40Z [Normal] Started container block-cloud-metadata
2021-10-14T16:36:40Z [Normal] Container image "pytorch" already present on machine
2021-10-14T16:36:40Z [Normal] Created container notebook
2021-10-14T16:36:40Z [Normal] Started container notebook
2021-10-14T16:36:42Z [Warning] Back-off restarting failed container
Spawn failed: Server at http://172.17.0.7:8888/user/jovyan/ didn't respond in 30 seconds
```

It appears there is no JupyterHub installed and the initial
`start-singleuser.sh` is not there as is required by the `base-notebook`
image from `docker-stacks` official images:

```
# Configure container startup
ENTRYPOINT ["tini", "-g", "--"]
CMD ["start-notebook.sh"]

# Copy local files as late as possible to avoid cache busting
COPY start.sh start-notebook.sh start-singleuser.sh /usr/local/bin/

# Currently need to have both jupyter_notebook_config and
# jupyter_server_config to support classic and lab COPY
jupyter_notebook_config.py /etc/jupyter/
```

Here's everything that seems broken:

1. `start-singleuser.sh` is missing so pod startup fails
1. nothing is talking on 8888 (no JupyterHub)

Also, the `base` uses `tini` to run the `jupyterhub` command as a
service, which is why the `ENTRYPOINT` is different and fails with the
default PyTorch image.

## Thursday, October 14, 2021, 10:35:43PM EDT

* Really annoyed with merging `pytorch` and `base-notebook`:
  * Conflicting versions of fucking Anaconda
  * Unresolved dependencies all over the place
  * Going to require fresh image creation almost from scratch
  * base requires python 3.9.7, PyTorch on Python 3.8.8
* Struggling with decision about how to proceed
  * If totally new image value of "tutorial" blown
  * Wasted time on image creation might distract from base usage
