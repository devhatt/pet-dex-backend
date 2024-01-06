# syntax=docker/dockerfile:1

FROM golang:1.21.4 as build

# Set destination for COPY
RUN mkdir -p /go/api
WORKDIR /go/api

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY /api ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o pet-dex-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

FROM alpine:latest

WORKDIR /

COPY --from=build /go/api/pet-dex-api /

EXPOSE 8080

# Run
CMD ["./pet-dex-api"]