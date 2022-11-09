package handlers

import (
	"bytes"
	"context"
	"nant-health-svc/db"
	"nant-health-svc/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPostMetrics_StatusCreated(t *testing.T) {
	body := []byte(`{"machineId":12345,"stats":{"cpuTemp":90,"fanSpeed":400,"HDDSpace":800},"lastLoggedIn":"","sysTime":"0001-01-01T00:00:00Z"}`)

	r, err := http.NewRequest(http.MethodPost, "/v1/metrics", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(PostMetrics)
	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"machineId":12345,"stats":{"cpuTemp":90,"fanSpeed":400,"HDDSpace":800},"lastLoggedIn":"","sysTime":"0001-01-01T00:00:00Z"}]`

	if reflect.DeepEqual(w.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestPostMetrics_StatusBadRequest(t *testing.T) {
	body := []byte(`{"machineId":,"stats":{"cpuTemp":90,"fanSpeed":400,"HDDSpace":800},"lastLoggedIn":"","sysTime":"0001-01-01T00:00:00Z"}`)

	r, err := http.NewRequest(http.MethodPost, "/v1/metrics", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(PostMetrics)
	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostMetrics_StatusBadRequest_ServiceLayer(t *testing.T) {
	body := []byte(`{"machineId":0,"stats":{"cpuTemp":90,"fanSpeed":400,"HDDSpace":800},"lastLoggedIn":"","sysTime":"0001-01-01T00:00:00Z"}`)

	r, err := http.NewRequest(http.MethodPost, "/v1/metrics", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(PostMetrics)
	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetMetrics(t *testing.T) {
	// initialize db.
	db.NewClient()

	// save dummy record in db.
	doc := models.Metrics{
		MachineID: 12345,
		Stats: models.Stats{
			CPUTemp:  90,
			FanSpeed: 400,
			HDDSpace: 800,
		},
	}
	db.SaveMetrics(context.TODO(), doc)

	r, err := http.NewRequest(http.MethodGet, "/v1/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMetrics)
	handler.ServeHTTP(w, r)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"machineId":12345,"stats":{"cpuTemp":90,"fanSpeed":400,"HDDSpace":800},"lastLoggedIn":"","sysTime":"0001-01-01T00:00:00Z"}]`

	if reflect.DeepEqual(w.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
