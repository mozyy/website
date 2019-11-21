FROM golang:1.13.0

ENV GOPROXY https://goproxy.cn

WORKDIR /
COPY ./go.yyue.dev ./go.yyue.dev
# 打包应用
RUN mkdir /go-apps && \
    cd ./go.yyue.dev && \
    CGO_ENABLED=0 GOOS=linux go build -a -o /go-apps ./...