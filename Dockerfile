FROM golang:1.21.4 as build
WORKDIR /go/api
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o pet-dex-api ./api/

FROM alpine:latest as api
COPY --from=build /go/api .
EXPOSE 3000
CMD ["./pet-dex-api"]
