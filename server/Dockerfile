FROM golang:1.16

COPY . $GOPATH/src/test-tls
COPY ./tls $GOPATH/bin/tls

WORKDIR $GOPATH/src/test-tls
RUN GO111MODULE=on GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o $GOPATH/bin/main

WORKDIR $GOPATH/bin
ENTRYPOINT ["./main"]
