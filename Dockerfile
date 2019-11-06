FROM alpine:3.10.1
LABEL maintainer="Haruaki Tamada" \
      ttt-version="1.0.0" \
      description="Verifying diploma of courses on ISE, KSU."

RUN    adduser -D ttt \
    && apk --no-cache add curl=7.66.0-r0 tar=1.32-r0 \
    && curl -s -L -O https://github.com/tamada/ttt/releases/download/v1.0.0/ttt-1.0.0_linux_amd64.tar.gz \
    && tar xfz ttt-1.0.0_linux_amd64.tar.gz  \
    && mv ttt-1.0.0 /opt                     \
    && ln -s /opt/ttt-1.0.0 /opt/ttt         \
    && ln -s /opt/ttt /usr/local/share/ttt   \
    && rm ttt-1.0.0_linux_amd64.tar.gz       \
    && ln -s /opt/ttt/ttt /usr/local/bin/ttt

ENV HOME="/home/ttt"

WORKDIR /home/ttt
USER    ttt

ENTRYPOINT [ "ttt" ]
