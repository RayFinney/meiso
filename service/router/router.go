package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"meiso/devices"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(devices.InjectDevice)

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	return e
}
