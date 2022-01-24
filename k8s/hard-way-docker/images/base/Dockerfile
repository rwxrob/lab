ARG BASE_IMAGE=ubuntu
FROM $BASE_IMAGE

LABEL MAINTAINER "Rob Muhlestein <rob@rwx.gg>"
LABEL SOURCE "https://github.com/rwxrob/lab-k8s-hard-way-docker"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get -y --no-install-recommends upgrade && \
    apt-get update -y && \
    apt-get install -y --no-install-recommends \
        ca-certificates vim net-tools ssh openssh-server \
        apt-transport-https curl bash-completion \
        gnupg lsb-release

# ------------------ container runtime (containerd) ------------------

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt-get update -y && \
    apt-get -y install docker-ce

# ------------------ install kubelet kubeadm kubectl -----------------

RUN curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg 

RUN echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list

RUN apt-get update -y && \
    apt-get install -y kubelet kubeadm kubectl && \
    apt-mark hold kubelet kubeadm kubectl

    #apt-get clean && \
    #rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* /var/log/dmesg.* && \
    #cat /dev/null > /var/log/dmesg

# TODO install container runtime
# TODO install kube-proxy

