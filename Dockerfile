FROM golang:1.18.0-alpine3.15 AS builder
WORKDIR /build
RUN apk add --no-cache git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM alpine:3.15
WORKDIR /root
ENV TZ Asia/Shanghai
# 修改时区为东八区
RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata
RUN apk add gcompat
COPY --from=builder /build/app /root/app
EXPOSE 8080
ENTRYPOINT ./app