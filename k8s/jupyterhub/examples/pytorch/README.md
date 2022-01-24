# `base-notebook` from `docker-stack` + `pytorch`

This is the `base-notebook` as is with only the `FROM` line changed to
be built from `nvcr.io/nvidia/pytorch`.

----

Here's an overview of how to merge the `base-notebook` image with the
latest `nvcr.io/nvidia/pytorch` image. You can adapt this approach to
other NCG images.

Pull `nvcr.io/nvidia/pytorch` locally and tag for convenience (releases
will vary).

```
docker pull nvcr.io/nvidia/pytorch:21.04-py3
docker tag nvcr.io/nvidia/pytorch:21.04-py3 pytorch
```

Copy out the example `Dockerfile` and put someplace under your version
management.

```
docker run pytorch cat docker-examples/Dockerfile.custompytorch > Dockerfile
```

This file has a bunch of stuff you should read through and fully
understand. Take particular note, of the legal information ---
especially the part about "derivative works" which may mean that you
should not publish your extended image publicly (but ask your own legal
department).

```Dockerfile

=============
== PyTorch ==
=============

NVIDIA Release 21.04 (build 22382700)
PyTorch Version 1.9.0a0+2ecb2c7

Container image Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.

Copyright (c) 2014-2021 Facebook Inc.
Copyright (c) 2011-2014 Idiap Research Institute (Ronan Collobert)
Copyright (c) 2012-2014 Deepmind Technologies    (Koray Kavukcuoglu)
Copyright (c) 2011-2012 NEC Laboratories America (Koray Kavukcuoglu)
Copyright (c) 2011-2013 NYU                      (Clement Farabet)
Copyright (c) 2006-2010 NEC Laboratories America (Ronan Collobert, Leon Bottou, Iain Melvin, Jason Weston)
Copyright (c) 2006      Idiap Research Institute (Samy Bengio)
Copyright (c) 2001-2004 Idiap Research Institute (Ronan Collobert, Samy Bengio, Johnny Mariethoz)
Copyright (c) 2015      Google Inc.
Copyright (c) 2015      Yangqing Jia
Copyright (c) 2013-2016 The Caffe contributors
All rights reserved.

NVIDIA Deep Learning Profiler (dlprof) Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.

Various files include modifications (c) NVIDIA CORPORATION.  All rights reserved.

This container image and its contents are governed by the NVIDIA Deep Learning Container License.
By pulling and using the container, you accept the terms and conditions of this license:
https://developer.nvidia.com/ngc/nvidia-deep-learning-container-license

WARNING: The NVIDIA Driver was not detected.  GPU functionality will not be available.
   Use 'nvidia-docker run' to start this container; see
   https://github.com/NVIDIA/nvidia-docker/wiki/nvidia-docker .

NOTE: MOFED driver for multi-node communication was not detected.
      Multi-node communication performance may be reduced.

NOTE: The SHMEM allocation limit is set to the default of 64MB.  This may be
   insufficient for PyTorch.  NVIDIA recommends the use of the following flags:
   nvidia-docker run --ipc=host ...

#
# This example Dockerfile illustrates a method to apply
# patches to the source code in NVIDIA's PyTorch
# container image and to rebuild PyTorch.  The RUN command
# included below will rebuild PyTorch in the same way as
# it was built in the original image.
#
# By applying customizations through a Dockerfile and
# `docker build` in this manner rather than modifying the
# container interactively, it will be straightforward to
# apply the same changes to later versions of the PyTorch
# container image.
#
# https://docs.docker.com/engine/reference/builder/
#
FROM nvcr.io/nvidia/pytorch:21.04-py3

# Bring in changes from outside container to /tmp
# (assumes my-pytorch-modifications.patch is in same directory as Dockerfile)
COPY my-pytorch-modifications.patch /tmp

# Change working directory to PyTorch source path
WORKDIR /opt/pytorch

# Apply modifications
RUN patch -p1 < /tmp/my-pytorch-modifications.patch

# Rebuild PyTorch
RUN cd pytorch && \
    CUDA_HOME="/usr/local/cuda" \
    CMAKE_PREFIX_PATH="$(dirname $(which conda))/../" \
    NCCL_INCLUDE_DIR="/usr/include/" \
    NCCL_LIB_DIR="/usr/lib/" \
    USE_SYSTEM_NCCL=1 \
    USE_OPENCV=1 \
    pip install --no-cache-dir -v .

# Reset default working directory
WORKDIR /workspace
```

The only reason we grabbed a copy of that file is to be sure we are
extending something that is supported by the original image creators. We
only actually need a few lines from it.

```Dockerfile
FROM nvcr.io/nvidia/pytorch:21.04-py3

WORKDIR /workspace
```
