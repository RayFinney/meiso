package sensors

import (
	"github.com/labstack/echo/v4"
	"meiso/models"
	"net/http"
)

type Delivery struct {
	sensorsService Service
}

func NewDelivery(sensorsService Service) Delivery {
	return Delivery{
		sensorsService: sensorsService,
	}
}

func (d *Delivery) StoreStatistics(ctx echo.Context) error {
	sensorStats := new(models.SensorStats)
	if err := ctx.Bind(sensorStats); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	if err := d.sensorsService.StoreStatistics(ctx.Request().Context(), sensorStats); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusCreated)
}
