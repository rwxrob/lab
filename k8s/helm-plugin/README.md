# Learning Lab: Create a Basic Helm Plugin

## Ways to Install a Helm Plugin

1. `cp -r myplugin $(helm env HELM_PLUGINS)`
1. `helm plugin install https://github.com/you/yourplugin`
1. `helm plugin install yourplugin`
1. `helm plugin install https://yoursite/yourplugin.{tar,tar.gz}`
1. `helm plugin install yourplugin.{tar,tar.gz}`

## Related

* Helm Plugins  
  <https://helm.sh/docs/topics/plugins/>
