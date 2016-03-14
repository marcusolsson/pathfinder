package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/marcusolsson/pathfinder/path"
)

type loggingService struct {
	logger log.Logger
	PathService
}

func (s loggingService) ShortestPath(origin, destination string) (paths []path.TransitPath, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "shortest_path",
			"origin", origin,
			"destination", destination,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.PathService.ShortestPath(origin, destination)
}
