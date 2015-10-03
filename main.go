package main

import (
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

var defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	addr := ":" + port

	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stdout)

	var ps PathService
	ps = pathService{}
	ps = loggingMiddleware{logger, ps}

	shortestPathHandler := httptransport.NewServer(
		ctx,
		makeShortestPathEndpoint(ps),
		decodeShortestPathRequest,
		encodeResponse,
	)

	http.Handle("/paths", method("GET", shortestPathHandler))

	_ = logger.Log("msg", "HTTP", "addr", addr)
	_ = logger.Log("err", http.ListenAndServe(addr, nil))
}

func method(m string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != m {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		h.ServeHTTP(w, r)
	})
}
