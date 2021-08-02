package plants

import "github.com/labstack/echo/v4"

func Register(g *echo.Group, d Delivery) {
	plantsG := g.Group("/plants")
	plantsG.POST("")
	plantsG.PUT("/:id")
	plantsG.GET("")
	plantsG.GET("/:id")
	plantsG.DELETE("/:id")
}
