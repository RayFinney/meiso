package sensors

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
)

func Setup(g *echo.Group, influxClient influxdb2.Client) (Delivery, Service, Repository) {
	r := NewRepository(influxClient)
	s := NewService(r)
	d := NewDelivery(s)
	Register(g, d)
	return d, s, r
}
