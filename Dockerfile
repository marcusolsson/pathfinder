# build stage
FROM billyteves/alpine-golang-glide:latest AS build-env
ADD . /go/src/github.com/marcusolsson/pathfinder
RUN cd /go/src/github.com/marcusolsson/pathfinder && glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pathfinder .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/marcusolsson/pathfinder/pathfinder /app/
ADD docs /docs
EXPOSE 8080
CMD ["/app/pathfinder"]

