package handlers

import (
	"context"
	"encoding/json"
	"nant-health-svc/models"
	"nant-health-svc/service"
	"net/http"
)

// PostMetrics will handle post requests for metrics.
func PostMetrics(w http.ResponseWriter, r *http.Request) {
	// location to decode.
	var doc models.Metrics

	// NewDecoder will decode the metrics to doc location.
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	// make a call to service layer
	if err := service.CreateMetrics(context.TODO(), doc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	// send 201 created as response code.
	w.WriteHeader(http.StatusCreated)
}

// GetMetrics will handle Get requests for metrics.
func GetMetrics(w http.ResponseWriter, r *http.Request) {
	// make a cal to service layer to retrieve the data.
	metrics := service.GetMetrics(context.TODO())

	// Encode the data to Json
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}
}
