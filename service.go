package pathfinder

import (
	"errors"

	"github.com/marcusolsson/pathfinder/path"
)

var ErrInvalidArgument = errors.New("invalid argument")

// PathService provides the shortest path "algoritm".
type PathService interface {
	ShortestPath(origin, destination string) ([]path.TransitPath, error)
}

type pathService struct{}

func NewPathService() PathService {
	return pathService{}
}

func (pathService) ShortestPath(origin, destination string) ([]path.TransitPath, error) {
	if origin == "" || destination == "" {
		return nil, ErrInvalidArgument
	}
	return path.FindShortestPath(origin, destination), nil
}
