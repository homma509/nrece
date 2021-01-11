package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter ...
func NewRouter() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	basePath := "nrece"

	appHandler := NewAppHandler()
	healthCheckHandler := NewHealthCheckHandler()
	helloHandler := NewHelloHandler()

	e.GET("/healthcheck", healthCheckHandler.HealthCheck())
	e.GET(basePath+"/v1/hello", helloHandler.SayHello())
	e.GET(basePath+"/v1/app", appHandler.GetApp())
	return e
}
