package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

const defaultPort = "8080"

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")

		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	var ps PathService
	ps = pathService{}
	ps = loggingService{log.NewContext(logger).With("component", "path"), ps}

	httpLogger := log.NewContext(logger).With("component", "http")

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(httpLogger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	shortestPathHandler := httptransport.NewServer(
		ctx,
		makeShortestPathEndpoint(ps),
		decodeShortestPathRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/paths", shortestPathHandler).Methods("GET")
	r.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	http.Handle("/", r)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
