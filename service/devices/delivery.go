package devices

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Delivery struct {
	devicesService Service
}

func NewDelivery(devicesService Service) Delivery {
	return Delivery{
		devicesService: devicesService,
	}
}

func (d *Delivery) StoreStartupLogs(ctx echo.Context) error {
	if err := d.devicesService.StoreStartupLogs(ctx.Request().Context()); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusCreated)
}

func (d *Delivery) GetLogs(ctx echo.Context) error {
	logs, err := d.devicesService.GetLogs()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, logs)
}
