FROM golang:alpine AS builder
ADD ./ /data/app/
WORKDIR /data/app/cmd

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scheduler_linux_amd64 GoApplication.go

CMD ["./scheduler_linux_amd64", "-configs=../configs/application.yaml"]