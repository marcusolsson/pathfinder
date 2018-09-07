package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/marcusolsson/pathfinder"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Paths pathfinder.PathService

	Logger log.Logger

	router chi.Router
}

// New returns a new HTTP server.
func New(ps pathfinder.PathService, logger log.Logger) *Server {
	s := &Server{
		Paths:  ps,
		Logger: logger,
	}

	r := chi.NewRouter()

	r.Use(accessControl)

	r.Get("/paths", s.shortestPaths)
	r.Method("GET", "/docs", http.StripPrefix("/docs/", http.FileServer(http.Dir("api"))))

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) shortestPaths(w http.ResponseWriter, r *http.Request) {
	var (
		from = r.URL.Query().Get("from")
		to   = r.URL.Query().Get("to")
	)

	paths, err := s.Paths.ShortestPath(from, to)
	if err != nil {
		if err == pathfinder.ErrInvalidArgument {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Paths interface{} `json:"paths"`
	}{
		Paths: paths,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
