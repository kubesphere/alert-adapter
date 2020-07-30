FROM golang:1.11-alpine3.7 as golang

ADD . /go/src/kubesphere.io/alert-adapter
WORKDIR /go/src/kubesphere.io/alert-adapter

RUN apk add --update --no-cache ca-certificates git

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN mkdir -p /adapter_bin
RUN go build -v -a -installsuffix cgo -ldflags '-w' -o /adapter_bin/adapter cmd/adapter/main.go


FROM alpine:3.7
RUN apk add --no-cache bash ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# modify pod (container) timezone
RUN apk add -U tzdata && ls /usr/share/zoneinfo && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata

COPY --from=golang /adapter_bin/adapter /alerting/adapter

RUN adduser -D -g alerter -u 1002 alerter && \
    chown -R alerter:alerter /alerting
USER alerter

EXPOSE 9200
EXPOSE 9201
EXPOSE 8080
