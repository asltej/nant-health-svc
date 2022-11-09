package handlers

import (
	"nant-health-svc/config"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	// initialize router.
	r := mux.NewRouter()

	// register handlers.
	r.HandleFunc("/v1/metrics", PostMetrics).Methods(http.MethodPost)
	r.HandleFunc("/v1/metrics", GetMetrics).Methods(http.MethodGet)

	// listen and serve on port.
	http.ListenAndServe(":"+config.HTTPPort, r)
}
