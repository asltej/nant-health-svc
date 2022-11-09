package db

import (
	"context"
	"errors"
	"nant-health-svc/models"
)

var (
	// Create Inmemory location
	metricsCollection []models.Metrics
)

func NewClient() {
	metricsCollection = make([]models.Metrics, 0)
}

// SaveMetrics will save the metrics record.
func SaveMetrics(ctx context.Context, doc models.Metrics) error {
	// validate machine ID
	if doc.MachineID == 0 {
		return errors.New("MachineID(i.e., Primary key) is required")
	}

	// save record to DB.
	metricsCollection = append(metricsCollection, doc)

	return nil
}

// GetMetrics will return the metrics records.
func GetMetrics(ctx context.Context) []models.Metrics {
	// return all metrics records.
	return metricsCollection
}
