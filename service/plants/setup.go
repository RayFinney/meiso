package plants

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func Setup(g *echo.Group, pgx *pgxpool.Pool) (Delivery, Service, Repository) {
	r := NewRepository(pgx)
	s := NewService(r)
	d := NewDelivery(s)
	Register(g, d)
	return d, s, r
}
