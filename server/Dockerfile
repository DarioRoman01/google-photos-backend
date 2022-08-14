ARG GO_VERSION=1.18.4

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src/

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY bucket bucket
COPY command-service command-service
COPY database database
COPY mail-service mail-service
COPY mailpb mailpb
COPY middlewares middlewares
COPY models models
COPY upload-service upload-service
COPY uploadpb uploadpb
COPY utils utils

RUN go install ./...

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=builder /go/bin .