pathfinder 
==========

[![Build Status](https://travis-ci.org/marcusolsson/pathfinder.svg?branch=master)](https://travis-ci.org/marcusolsson/pathfinder)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

The routing context from the original DDD Sample Application, written in Go.

## Running the application

Start the application on port 8080 (or whatever the `PORT` variable is set to).

```
go run cmd/pathfinder/main.go
```

### Docker

You can also run the application using Docker:

```
docker run --name some-pathfinder -p 8080:8080 marcusolsson/pathfinder
```

## Try it!

```
curl 'localhost:8080/paths?from=SESTO&to=FIHEL'
```
