# syntax=docker/dockerfile:1

FROM golang:1.21.4 as build

RUN mkdir -p /go/api
WORKDIR /go/api

COPY go.mod go.sum ./
RUN go mod download

COPY /api ./

RUN CGO_ENABLED=0 GOOS=linux go build -o pet-dex-api

EXPOSE 8080

FROM alpine:latest
WORKDIR /

COPY --from=build /go/api/pet-dex-api /

EXPOSE 8080

CMD ["./pet-dex-api"]
