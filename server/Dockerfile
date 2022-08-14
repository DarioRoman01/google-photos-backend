ARG GO_VERSION=1.18.4

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src/

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY bucket bucket
COPY models models

RUN go install

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=builder /go/bin .