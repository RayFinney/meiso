package sensors

import "github.com/labstack/echo/v4"

func Register(g *echo.Group, d Delivery) {
	sensorsG := g.Group("/sensors")
	sensorsG.POST("/statistics", d.StoreStatistics)
}
