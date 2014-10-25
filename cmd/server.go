package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/marcusolsson/pathfinder"
	"gopkg.in/unrolled/render.v1"
)

func main() {

	r := render.New(render.Options{IndentJSON: true})
	router := mux.NewRouter()

	router.HandleFunc("/paths", func(w http.ResponseWriter, req *http.Request) {
		from := req.URL.Query().Get("from")
		to := req.URL.Query().Get("to")

		if len(from) == 0 || len(to) == 0 {
			r.JSON(w, http.StatusBadRequest, map[string]interface{}{"error": "missing parameters"})
			return
		}

		r.JSON(w, http.StatusOK, pathfinder.FindShortestPath(from, to))
	}).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
