package sensors

import (
	"context"
	"github.com/sirupsen/logrus"
	"meiso/models"
)

type Service struct {
	sensorsRepo Repository
}

func NewService(sensorsRepo Repository) Service {
	return Service{
		sensorsRepo: sensorsRepo,
	}
}

func (s *Service) StoreStatistics(ctx context.Context, sensorStats *models.SensorStats) error {
	deviceId := ctx.Value("device-id").(string)
	if deviceId == "" {
		return DEVICE_ID_REQUIRED
	}
	if err := s.sensorsRepo.StoreStatistics(ctx, deviceId, sensorStats); err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "sensors",
			"function": "StoreStatistics",
		}).Warningf("unable to store statistics: %v", err)
		return err
	}
	return nil
}
