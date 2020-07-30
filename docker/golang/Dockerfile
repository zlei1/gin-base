FROM golang:1.14-alpine AS builder

RUN apk add --update --no-cache \
      --repository http://mirrors.aliyun.com/alpine/v3.12/community \
      --repository http://mirrors.aliyun.com/alpine/v3.12/main

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o gin-base .

FROM debian:stretch-slim
WORKDIR /deploy
COPY --from=builder /app/gin-base .
COPY config ./config
