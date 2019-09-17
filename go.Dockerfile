FROM golang:1.13.0 AS builder

WORKDIR /
COPY ./go ./build
RUN cd ./build && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app ./app.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY ./docker ./docker
COPY --from=builder app .
EXPOSE 6503
CMD ["./app"]