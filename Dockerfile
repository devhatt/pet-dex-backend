FROM golang:1.21.4 as build
WORKDIR /usr/src/app/go/api
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o pet-dex-api ./api/

FROM alpine:3.15.11 as api
WORKDIR /usr/src/app/go/api
COPY --from=build /usr/src/app/go/api .
EXPOSE 3000
CMD ["./pet-dex-api"]
