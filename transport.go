package pathfinder

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

var errInvalidArgument = errors.New("invalid argument")

func MakeHTTPHandler(ctx context.Context, ps PathService, logger log.Logger) http.Handler {
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	shortestPathHandler := httptransport.NewServer(
		makeShortestPathEndpoint(ps),
		decodeShortestPathRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/paths", shortestPathHandler).Methods("GET")
	r.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))

	return r
}

func decodeShortestPathRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		from = r.URL.Query().Get("from")
		to   = r.URL.Query().Get("to")
	)
	return shortestPathRequest{From: from, To: to}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case errInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
