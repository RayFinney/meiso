package sensors

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"meiso/influx"
	"meiso/models"
	"time"
)

type Repository struct {
	influxClient influxdb2.Client
}

func NewRepository(influxClient influxdb2.Client) Repository {
	return Repository{
		influxClient: influxClient,
	}
}

func (r *Repository) StoreStatistics(ctx context.Context, deviceId string, stats *models.SensorStats) error {
	writeAPI := r.influxClient.WriteAPIBlocking(influx.Organization, influx.Bucket)
	p := influxdb2.NewPoint(
		"sensors",
		map[string]string{
			"id": deviceId,
		},
		map[string]interface{}{
			"temperature": stats.Temp,
			"humidity":    stats.Humidity,
			"lux":         stats.Lux,
		},
		time.Now())
	// write asynchronously
	err := writeAPI.WritePoint(ctx, p)
	return err
}
