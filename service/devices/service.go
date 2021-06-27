package devices

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Service struct {
	devicesRepo Repository
}

func NewService(devicesRepo Repository) Service {
	return Service{
		devicesRepo: devicesRepo,
	}
}

func (s *Service) StoreStartupLogs(ctx context.Context) error {
	deviceId := ctx.Value("device-id").(string)
	if err := s.devicesRepo.StoreStartupLogs(deviceId); err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "devices",
			"function": "StoreStartupLogs",
		}).Warningf("unable to store startup in logs: %v", err)
		return err
	}
	return nil
}

func (s *Service) GetLogs() (string, error) {
	logs, err := s.devicesRepo.GetLogs()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":  "devices",
			"function": "GetLogs",
		}).Warningf("unable to read logs: %v", err)
		return "", err
	}
	return logs, nil
}
