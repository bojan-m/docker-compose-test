FROM golang:alpine
COPY /src /go/src/docker-compose-test/faas_gateway/src
WORKDIR /go/src/docker-compose-test/faas_gateway/src
RUN apk add git
RUN go get github.com/docker/docker/client
CMD go run gateway.go
