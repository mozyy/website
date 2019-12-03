FROM golang:1.13.0

ENV GOPROXY https://goproxy.cn

WORKDIR /

COPY ./docker/docker/go/mod ./go.yyue.dev

RUN cd ./go.yyue.dev && \
    go mod download

COPY ./go.yyue.dev ./go.yyue.dev

# 构建应用
RUN mkdir /go-apps

# 一起构建对内存要求较高, 为保证小内存机器构建, 所以单个应用构建
# RUN cd ./go.yyue.dev && \
#     CGO_ENABLED=0 GOOS=linux go build -a -o /go-apps ./...

RUN cd ./go.yyue.dev && \
    CGO_ENABLED=0 GOOS=linux go build -a -o /go-apps/datamanage ./datamanage/main.go

RUN cd ./go.yyue.dev && \
    CGO_ENABLED=0 GOOS=linux go build -a -o /go-apps/crawler ./crawler/main.go