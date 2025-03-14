# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN pwd && ls && go build -o /ramock ./cmd/ramock/main.go

FROM alpine:3.14 as production-stage

COPY --from=build-stage /ramock /ramock
COPY data/endpoints-example.yaml /config/endpoints.yaml

ENV RAMOCK_LISTEN_PORT="8008"
ENV RAMOCK_STRICT_PATHS=true
ENV RAMOCK_ENDPOINTS_FILE=/config/endpoints.yaml
ENV LOGGING_APP_TAG="ramock-docker"

EXPOSE 8008

VOLUME /config

ENTRYPOINT ["/ramock"]