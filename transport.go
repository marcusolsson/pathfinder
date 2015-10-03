package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/marcusolsson/pathfinder/path"
	"golang.org/x/net/context"
)

type shortestPathRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type shortestPathResponse struct {
	Paths []path.TransitPath `json:"paths"`
}

func makeShortestPathEndpoint(ps PathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(shortestPathRequest)
		paths := ps.ShortestPath(req.From, req.To)
		return shortestPathResponse{Paths: paths}, nil
	}
}

func decodeShortestPathRequest(r *http.Request) (interface{}, error) {
	var (
		from = r.URL.Query().Get("from")
		to   = r.URL.Query().Get("to")
	)

	if from == "" || to == "" {
		return nil, errors.New("missing parameters")
	}

	return shortestPathRequest{From: from, To: to}, nil
}

func encodeResponse(w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}
