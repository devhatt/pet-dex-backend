# syntax=docker/dockerfile:1

FROM golang:1.21.4 as build

RUN mkdir -p /go/api
WORKDIR /go/api

COPY go.mod go.sum ./
COPY /api ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o pet-dex-api

FROM alpine:latest
WORKDIR /

COPY --from=build /go/api/pet-dex-api /

EXPOSE 8080

CMD ["./pet-dex-api"]
