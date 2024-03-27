FROM golang:1.21.4

WORKDIR /app
ENV INTEGRATION=true

COPY go.mod go.sum ./ 
RUN go mod download

COPY . ./

CMD [ "go", "test", "./integration-tests/..." ]
