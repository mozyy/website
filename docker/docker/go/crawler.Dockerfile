FROM golang:1.13.0 AS builder

ARG APP_NAME=crawler
ARG EXPOSE_PORT
ENV GOPROXY https://goproxy.cn

WORKDIR /
COPY ./go/common ./build/common
COPY ./go/go.mod ./build/go.mod
# 复制当前应用
COPY ./go/$APP_NAME ./build/$APP_NAME
# 打包应用
RUN cd ./build && \
    CGO_ENABLED=0 GOOS=linux go build -a -o /app ./$APP_NAME/main.go

FROM alpine:latest
WORKDIR /
COPY ./docker ./docker
COPY --from=builder /app ./app
EXPOSE $EXPOSE_PORT
CMD ["./app"]
