FROM golang:1.22.2-bookworm
COPY go.* /go/src/github.com/flowerinthenight/oomkill-trace/
COPY *.go /go/src/github.com/flowerinthenight/oomkill-trace/
WORKDIR /go/src/github.com/flowerinthenight/oomkill-trace/
RUN GOFLAGS=-mod=vendor GOOS=linux go build -v -trimpath -o oomkill-trace .

FROM alpine:3.19.1
RUN apk --no-cache add ca-certificates bcc-tools && ls -laF /usr/share/bcc/tools/
WORKDIR /app/
COPY --from=0 /go/src/github.com/flowerinthenight/oomkill-trace/oomkill-trace .
ENTRYPOINT ["/app/oomkill-trace"]
CMD ["-local=false"]
