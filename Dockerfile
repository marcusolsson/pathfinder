FROM golang:1.10.3-alpine AS build-env
ADD . /go/src/github.com/marcusolsson/pathfinder
WORKDIR /go/src/github.com/marcusolsson/pathfinder
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o goapp ./cmd/pathfinder/main.go

FROM alpine:3.7
WORKDIR /app
COPY --from=build-env /go/src/github.com/marcusolsson/pathfinder/goapp /app/
ADD docs /docs
EXPOSE 8080
CMD ["./goapp"]

