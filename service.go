package main

import "github.com/marcusolsson/pathfinder/path"

// PathService provides the shortest path "algoritm".
type PathService interface {
	ShortestPath(origin, destination string) ([]path.TransitPath, error)
}

type pathService struct{}

func (pathService) ShortestPath(origin, destination string) ([]path.TransitPath, error) {
	if origin == "" || destination == "" {
		return nil, errInvalidArgument
	}
	return path.FindShortestPath(origin, destination), nil
}
