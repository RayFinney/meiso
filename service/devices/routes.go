package devices

import "github.com/labstack/echo/v4"

func Register(g *echo.Group, d Delivery) {
	devicesG := g.Group("/devices")
	devicesG.POST("/startup", d.StoreStartupLogs)
	devicesG.GET("/logs", d.GetLogs)
}
