package pathfinder

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/marcusolsson/pathfinder/path"
)

type loggingMiddleware struct {
	logger log.Logger
	next   PathService
}

// NewLoggingMiddleware creates a new logging middleware.
func NewLoggingMiddleware(logger log.Logger, next PathService) PathService {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (s *loggingMiddleware) ShortestPath(origin, destination string) (paths []path.TransitPath, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "shortest_path",
			"origin", origin,
			"destination", destination,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.ShortestPath(origin, destination)
}
