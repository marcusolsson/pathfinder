package main

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
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

	var candidates []TransitPath

	for p := range generateCandidates(origin, destination, 3+rand.Intn(3)) {
		candidates = append(candidates, p)
	}

	return candidates
}

// generateCandidates generates new candidates and pushes them into a channel.
// Finding each candidate is potentially time-consuming, so we compute them in
// parallel.
func generateCandidates(origin, destination string, n int) chan TransitPath {
	ch := make(chan TransitPath)

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			ch <- findCandidate(origin, destination, nextDate(time.Now()))
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// findCandidate finds a random path between two locations starting from a
// given time.
func findCandidate(origin, destination string, start time.Time) TransitPath {
	v := allVertices(origin, destination)
	v = randChunk(v)

	var edges []TransitEdge

	edges, start = appendEdge(edges, origin, v[0], start)

	for j := 0; j < len(v)-1; j++ {
		edges, start = appendEdge(edges, v[j], v[j+1], start)
	}

	edges, _ = appendEdge(edges, v[len(v)-1], destination, start)

	return TransitPath{Edges: edges}
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

	shuffle(locations)

	t := len(locations)

	var c int
	if t > 4 {
		c = 1 + rand.Intn(5)
	} else {
		c = t
	}

	return locations[:c]
}

func shuffle(slc []string) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}

func nextDate(t time.Time) time.Time {
	n := time.Duration(rand.Intn(1000) - 500)
	return t.Add(24 * time.Hour).Add(n * time.Minute)
}

var port int

func main() {
	flag.IntVar(&port, "port", 8080, "the server port")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	r := render.New(render.Options{IndentJSON: true})
	router := mux.NewRouter()

	router.HandleFunc("/paths", func(w http.ResponseWriter, req *http.Request) {
		from := req.URL.Query().Get("from")
		to := req.URL.Query().Get("to")

		if len(from) == 0 || len(to) == 0 {
			r.JSON(w, http.StatusBadRequest, map[string]interface{}{"error": "missing parameters"})
			return
		}

		r.JSON(w, http.StatusOK, FindShortestPath(from, to))
	}).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
