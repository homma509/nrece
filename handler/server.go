package handler

import (
	"fmt"
	"log"

	"github.com/homma509/nrece/domain/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server ...
type Server struct {
	*echo.Echo
}

// NewServer ...
func NewServer(repo repository.AppRepository) (*Server, error) {
	s := &Server{
		echo.New(),
	}

	// Middleware
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	basePath := "nrece"

	appHandler, err := NewAppHandler(repo)
	if err != nil {
		log.Println("[error] new", err)
		return nil, err
	}
	healthCheckHandler := NewHealthCheckHandler()
	helloHandler := NewHelloHandler()

	s.GET("/healthcheck", healthCheckHandler.HealthCheck())
	s.GET(basePath+"/v1/hello", helloHandler.SayHello())
	s.GET(basePath+"/v1/app", appHandler.GetApp())
	return s, nil
}

// Run ...
func (s *Server) Run(port int) int {
	s.Logger.Fatal(s.Start(fmt.Sprintf(":%d", port)))
	return 0
}
