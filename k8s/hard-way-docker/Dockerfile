FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get -y update && apt-get -y upgrade
RUN apt-get -y install --no-install-recommends \
  vim-tiny curl gnupg ca-certificates

RUN curl -sSL \
  https://download.docker.com/linux/ubuntu/gpg \
  -o /etc/apt/trusted.gpg.d/docker.asc && \
  echo 'deb https://download.docker.com/linux/ubuntu focal stable' \
    > /etc/apt/sources.list.d/docker.list && \
  apt-get -y update && \
  apt-get -y install --no-install-recommends docker-ce  

RUN apt-get -y install systemd
RUN ln -s /usr/bin/systemd /bin/init

ENTRYPOINT ["/bin/init"]

# ------------------------------ cleanup -----------------------------

# RUN apt-get clean -y
# RUN rm -rf \
#    /var/cache/debconf/* \
#    /var/lib/apt/lists/* \
#    /var/log/* \
#    /tmp/* \
#    /var/tmp/* \
#    /usr/share/doc/* \
#    /usr/share/doc-base/* \
#    /usr/share/man/* \
#    /usr/share/local/*
