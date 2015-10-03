package main

import "github.com/marcusolsson/pathfinder/path"

// PathService provides the shortest path "algoritm".
type PathService interface {
	ShortestPath(origin, destination string) []path.TransitPath
}

type pathService struct{}

func (pathService) ShortestPath(origin, destination string) []path.TransitPath {
	return path.FindShortestPath(origin, destination)
}
