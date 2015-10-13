FROM golang:1.5.1

WORKDIR /go/src/github.com/marcusolsson/pathfinder
ADD . /go/src/github.com/marcusolsson/pathfinder

RUN go get github.com/tools/godep

RUN godep go install github.com/marcusolsson/pathfinder

ENTRYPOINT /go/bin/pathfinder

EXPOSE 8080

