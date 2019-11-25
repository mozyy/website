FROM docker_go_builder AS builder

FROM alpine:latest

ARG APP_NAME=datamanage
ARG EXPOSE_PORT=30001

WORKDIR /
COPY ./docker ./docker
COPY --from=builder /go-apps/$APP_NAME ./app
EXPOSE $EXPOSE_PORT
CMD ["./app"]
