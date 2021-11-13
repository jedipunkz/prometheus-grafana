FROM golang:1.17.3
MAINTAINER jedipunkz

WORKDIR /go/src/

ADD . /go/src/

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/sandbox-exporter

ENTRYPOINT ["/go/bin/sandbox-exporter"]

