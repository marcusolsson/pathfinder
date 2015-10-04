package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/marcusolsson/pathfinder/path"
)

type loggingMiddleware struct {
	logger log.Logger
	PathService
}

func (mw loggingMiddleware) ShortestPath(origin, destination string) (paths []path.TransitPath) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "shortest_path",
			"origin", origin,
			"destination", destination,
			"took", time.Since(begin),
		)
	}(time.Now())

	paths = mw.PathService.ShortestPath(origin, destination)
	return
}
