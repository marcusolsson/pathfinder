package pathfinder

import (
	"errors"

	"github.com/marcusolsson/pathfinder/path"
)

// ErrInvalidArgument is used if the user provides invalid arguments.
var ErrInvalidArgument = errors.New("invalid argument")

// PathService provides the shortest path "algorithm".
type PathService interface {
	ShortestPath(origin, destination string) ([]path.TransitPath, error)
}

type pathService struct{}

// NewPathService creates a new path service.
func NewPathService() PathService {
	return &pathService{}
}

func (*pathService) ShortestPath(origin, destination string) ([]path.TransitPath, error) {
	if origin == "" || destination == "" {
		return nil, ErrInvalidArgument
	}
	return path.FindShortestPath(origin, destination), nil
}
