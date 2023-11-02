FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN  make build

FROM debian:stable-slim

COPY --from=builder /src/docker/sources.list /tmp

RUN  mv /tmp/sources.list /etc/apt/  \
     && apt-get update && apt-get install -y --no-install-recommends \
        apt-transport-https\
        ca-certificates  \
        netbase \
     && rm -rf /var/lib/apt/lists/ \
     && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin/app /app/server

WORKDIR /app

EXPOSE 8090
VOLUME /data/conf

CMD ["./server", "-config", "/data/conf.yaml"]
