FROM golang:1.9
MAINTAINER Jordi Miguel <neoandroid@kaledoniah.net>

# Install VLC libs
RUN apt-get update && apt-get install -y --no-install-recommends vlc

WORKDIR /go/bin/app
COPY blackbird blackbird

RUN mkdir conf
COPY conf/app.prod.conf conf/app.conf

ENTRYPOINT ["/go/bin/app/blackbird"]
