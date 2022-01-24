FROM rwxrob/base
LABEL maintainer "Rob Muhlestein <rob@rwx.gg>"
USER root
COPY . .
RUN chown -R you:you .
USER you
CMD ["bash"]
