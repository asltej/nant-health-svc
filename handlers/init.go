package handlers

import (
	"nant-health-svc/config"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/metrics", PostMetrics).Methods(http.MethodPost)
	r.HandleFunc("/v1/metrics", GetMetrics).Methods(http.MethodGet)

	http.ListenAndServe(":"+config.HTTPPort, r)
}
