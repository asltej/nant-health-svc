package service

import (
	"context"
	"nant-health-svc/db"
	"nant-health-svc/models"
)

// CreateMetrics will create the metrics record and save it in db.
func CreateMetrics(ctx context.Context, doc models.Metrics) error {
	// save metrics in DB.
	return db.SaveMetrics(ctx, doc)
}

// GetMetrics will get the metrics records from db.
func GetMetrics(ctx context.Context) []models.Metrics {
	// get metrics from DB.
	return db.GetMetrics(ctx)
}
