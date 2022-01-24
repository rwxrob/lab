# Zero to JupyterHub in Kubernetes, a Learning Lab

This lab repo contains my exploration, learning, and decisions related
to getting JupyterHub running at scale with hundreds of users and the
requirements documented below. It has its own repo (as opposed to being
included in my [zettelkasten][zk]) because it contains enough code and
artifacts to be managed independently. Many of the lessons-learned,
however, are captured as stand-alone [zettel][zk] entries. See
[Reference](#reference) below for external references to documentation
used. I've included a [log] will chronological notes and discoveries as
they happen. I've also recorded video of many colearning/exploration
sessions in [a YouTube playlist][videos]. 

## Index

* [Log][log]
* [YouTube Playlist][videos]
* [Helm 1.0.1 Official Chart](chart-1.0.1)

## Scenario and Objectives

You've been asked to get JupyterHub running in a Kubernetes cluster for
an organization with 1000s of users who use JupyterHub daily and would
prefer to just open their web browser, authenticate in the standard way,
and be presented with a web page that prompts to use either one of a few
standard container images, or an image that they have uploaded to a
private, internal registry (Portus, Harbour). Users already have a
shared network space (`/s` drive) which should be available. Users also
want to login in as their own user IDs when connecting to the console
via JupyterHub rather than the default (`jovian`) account so that files
created will have the same user ID number as that which matches other
systems.

## Reference

* [YouTube Playlist][videos]
* [JupyterHub --- JupyterHub 1.4.1 documentation](https://jupyterhub.readthedocs.io/en/stable/index.html)
* [Zero to JupyterHub with Kubernetes --- Zero to JupyterHub with Kubernetes documentation](https://zero-to-jupyterhub.readthedocs.io/en/latest/index.html)
* [KubeSpawner --- kubespawner 1.0.1.dev0
documentation](https://jupyterhub-kubespawner.readthedocs.io/en/latest/spawner.html)

[log]: log.md
[zk]: <https://github.com/rwxrob/zet>
[videos]: <https://youtube.com/playlist?list=PLrK9UeDMcQLoqHu7NEm6YR_OzpXwiJqau>
