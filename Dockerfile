FROM golang:1.22.2-bookworm
COPY go.* /go/src/github.com/flowerinthenight/oomkill-trace/
COPY *.go /go/src/github.com/flowerinthenight/oomkill-trace/
WORKDIR /go/src/github.com/flowerinthenight/oomkill-trace/
RUN GOFLAGS=-mod=vendor GOOS=linux go build -v -trimpath -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o oomkill-trace .

# FROM debian:stable-slim
# RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && echo deb http://cloudfront.debian.net/debian sid main >> /etc/apt/sources.list && apt-get install -y bpfcc-tools libbpfcc libbpfcc-dev linux-headers-$(uname -r) && rm -rf /var/lib/apt/lists/*

FROM ubuntu:24.04
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates bpfcc-tools && rm -rf /var/lib/apt/lists/*
WORKDIR /app/
COPY --from=0 /go/src/github.com/flowerinthenight/oomkill-trace/oomkill-trace .
ENTRYPOINT ["/app/oomkill-trace"]
CMD [""]
