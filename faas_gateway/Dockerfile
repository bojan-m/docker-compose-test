FROM golang:alpine
COPY /src /go/src/docker-compose-test/faas_gateway/src
WORKDIR /go/src/docker-compose-test/faas_gateway/src
CMD go run gateway.go
