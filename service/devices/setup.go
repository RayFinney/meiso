package devices

import (
	"github.com/labstack/echo/v4"
)

func Setup(g *echo.Group) (Delivery, Service, Repository) {
	r := NewRepository()
	s := NewService(r)
	d := NewDelivery(s)
	Register(g, d)
	return d, s, r
}
