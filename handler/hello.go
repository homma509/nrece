package handler

import "github.com/labstack/echo"

// HelloHandler ...
type HelloHandler struct {
}

// NewHelloHandler ...
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// SayHello ...
func (handler *HelloHandler) SayHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "Hello World!")
	}
}
