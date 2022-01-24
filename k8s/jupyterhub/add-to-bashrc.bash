#!/bin/bash

_have minikube && . <(minikube completion bash)
_have mk && complete -o default -F __start_minikube mk

