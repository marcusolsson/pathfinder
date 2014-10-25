package pathfinder

import (
	"math/rand"
	"time"
)

// TransitEdge is a direct transportation between two locations.
type TransitEdge struct {
	VoyageNumber string    `json:"voyage"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
}

// TransitPath is a series of transit edges.
type TransitPath struct {
	Edges []TransitEdge `json:"edges"`
}

// FindShortestPath computes the shortest paths between two locations.
func FindShortestPath(origin, destination string) []TransitPath {
	start := time.Now()
	date := nextDate(start)

	v := allVertices(origin, destination)

	candidates := make([]TransitPath, 3+rand.Intn(3))
	for i := range candidates {
		v = randChunk(v)

		var edges []TransitEdge

		edges, date = appendEdge(edges, origin, v[0], date)

		for j := 0; j < len(v)-1; j++ {
			edges, date = appendEdge(edges, v[j], v[j+1], date)
		}

		edges, _ = appendEdge(edges, v[len(v)-1], destination, date)

		candidates[i] = TransitPath{Edges: edges}
	}

	return candidates
}

func appendEdge(edges []TransitEdge, curr, next string, date time.Time) ([]TransitEdge, time.Time) {
	var (
		from = nextDate(date)
		to   = nextDate(from)
	)

	edges = append(edges, TransitEdge{
		VoyageNumber: randVoyageNumber(),
		Origin:       curr,
		Destination:  next,
		Departure:    from,
		Arrival:      to,
	})

	return edges, nextDate(to)
}

func allVertices(origin, destination string) []string {
	locations := []string{"CNHKG", "AUMEL", "SESTO", "FIHEL", "USCHI", "JNTKO", "DEHAM", "CNSHA", "NLRTM", "SEGOT", "CNHGH", "USNYC", "USDAL"}
	for i, l := range locations {
		if l == origin {
			locations = append(locations[:i], locations[i+1:]...)
		}
	}
	for i, l := range locations {
		if l == destination {
			locations = append(locations[:i], locations[i+1:]...)
		}
	}
	return locations
}

func randVoyageNumber() string {
	switch rand.Intn(5) {
	case 0:
		return "0100S"
	case 1:
		return "0200T"
	case 2:
		return "0300A"
	case 3:
		return "0301S"
	}
	return "0400S"
}

func randChunk(locations []string) []string {
	// TODO: Shuffle the locations first for even more randomness.

	t := len(locations)

	var c int
	if t > 4 {
		c = 1 + rand.Intn(5)
	} else {
		c = t
	}

	return locations[:c]
}

func nextDate(t time.Time) time.Time {
	n := time.Duration(rand.Intn(1000) - 500)
	return t.Add(24 * time.Hour).Add(n * time.Minute)
}
