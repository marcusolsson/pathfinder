# build stage
FROM billyteves/alpine-golang-glide:latest AS build-env
ADD . /go/src/github.com/marcusolsson/pathfinder
RUN cd /go/src/github.com/marcusolsson/pathfinder && glide install && go build -o pathfinder cmd/pathfinder/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/marcusolsson/pathfinder/pathfinder /app/
ADD docs /docs
EXPOSE 8080
CMD ["/app/pathfinder"]

