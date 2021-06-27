package devices

import (
	"context"
	"github.com/labstack/echo/v4"
)

func InjectDevice(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		newRequest := ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), "device-id", ctx.Request().Header.Get("device-id")))
		*ctx.Request() = *newRequest
		return next(ctx)
	}
}
