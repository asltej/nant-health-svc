package db

import (
	"context"
	"nant-health-svc/models"
	"reflect"
	"testing"
)

func TestSaveMetrics(t *testing.T) {
	type args struct {
		ctx context.Context
		doc models.Metrics
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - Save Metrics Record",
			args: args{
				ctx: context.TODO(),
				doc: models.Metrics{
					MachineID: 12345,
					Stats: models.Stats{
						CPUTemp:  90,
						FanSpeed: 400,
						HDDSpace: 800,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Fail - Unknow Metrics Record",
			args: args{
				ctx: context.TODO(),
				doc: models.Metrics{
					Stats: models.Stats{
						CPUTemp:  90,
						FanSpeed: 400,
						HDDSpace: 800,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveMetrics(tt.args.ctx, tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("SaveMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMetrics(t *testing.T) {
	// initialize db.
	NewClient()

	// save dummy record in db.
	doc := models.Metrics{
		MachineID: 12345,
		Stats: models.Stats{
			CPUTemp:  90,
			FanSpeed: 400,
			HDDSpace: 800,
		},
	}
	SaveMetrics(context.TODO(), doc)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []models.Metrics
	}{
		{
			name: "Test - Get Metrics Record",
			args: args{
				ctx: context.TODO(),
			},
			want: []models.Metrics{
				{
					MachineID: 12345,
					Stats: models.Stats{
						CPUTemp:  90,
						FanSpeed: 400,
						HDDSpace: 800,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMetrics(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
