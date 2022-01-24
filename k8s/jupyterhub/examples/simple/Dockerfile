FROM ubuntu:focal

# Step One: Install all dependencies

USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y && apt-get install -y --no-install-recommends \
      tini \
      ca-certificates \
      curl \
      locales \
      python3-pip \
      npm \
      && \
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    echo "en_US.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen

# Step Two: Install JupyterHub and JupyterLab and Proxy

RUN python3 -m pip install notebook jupyterhub jupyterlab
RUN npm install -g configurable-http-proxy

# Step Three: Configure for Use with KubeSpawner / K8S

EXPOSE 8888

# Step Four: Setup Startup

#ENTRYPOINT ["tini", "-g", "--"]
#CMD ["start-
# ARGS passed by default:
#   jupyterhub
#   --config
#   /usr/local/etc/jupyterhub/jupyterhub_config.py
#   --upgrade-db
