# Learn Multicall Binaries (BusyBox, CmdBox, etc.)

Multicall is the method of causing a program to do different things
based solely on its name and creating copies and links with those
different names. This method is garnering more attention since the
popular [BusyBox] uses this method for the ultralight base [base
container] of the same name. BusyBox was originally created to fit
almost all the main shell tools from Linux onto a "single floppy disk." 

[CmdBox] also uses multicall, but allows developers to
write their commands in Go (instead of C) and compose them together.

[base container]: <https://hub.docker.com/_/busybox>
[BusyBox]: <https://www.busybox.net>
[CmdBox]:  <https://github.com/rwxrob/cmdbox>

## Learning Objective

## Requirements

Create a program that behaves differently depending on how it is named,
specifically it should include the following:

1. Create a program named `greet` that prints "Greetings!"
1. Write a `build` POSIX shell script to compile (if necessary)
1. Copy the `greet` to `hi` and have it print "Hi there!"
1. Symlink `hello` to local path to `greet` and print "Hello."
1. Symlink `salut` to absolute path to `greet` and print "Salut!"
1. Hard link `privet` to `greet` and print "Привет!"
1. Put all of this in `build` function in `build` script
1. Add a `run` function to `build` script that runs each
1. Add a `clean` function to `build` script (be careful)
1. Have `build` script take one arg matching function
1. Have `build` script default to `build` with no args

## Examples

One possible solution (implemented in Go) is contained in the
[examples/greet](examples/greet) subdirectory. Study the
[`build`](examples/greet/build) POSIX shell script specifically for ways
to implement the multicall approach.

## Getting Started

### Git

To use Git simply fork this repo and clone it to your local system. This
option requires you to have everything needed to code this from your own
terminal:

```
gh repo fork rwxrob/lab-multicall
gh repo clone fork lab-multicall
```

Note: This will require you to check your code to see that it meets the
requirements yourself (which you should learn to do anyway). If you
wanted more automated evaluation use the [Docker](#docker) container
instead.

### Docker Hub

This lab is hosted on [hub.docker.io] enabling you to use it directly


[fork this repo](#git) and change the [Dockerfile](Dockerfile) to build
`FROM <yourbase>`.

```
docker run -it -d --name multicall rwxrob/lab-multicall
```

[hub.docker.io]: <https://hub.docker.com/rwxrob/lab-multicall>
