package handler

import (
	"github.com/homma509/nrece/usecase"
	"github.com/labstack/echo"
)

// HealthCheckHandler ...
type HealthCheckHandler struct {
	*usecase.Healthcheck
}

// NewHealthCheckHandler ...
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{
		usecase.NewHealthcheck(),
	}
}

// HealthCheck ...
func (handler *HealthCheckHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		resJSON := handler.Get()
		return c.JSON(200, resJSON)
	}
}
