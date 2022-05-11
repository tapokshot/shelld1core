package httpServer

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tapokshot/shelld1core/log"
	"github.com/tapokshot/shelld1core/middleware"
)

type Server struct {
	Router *echo.Echo
	logger *log.Logger
}

func New(logger *log.Logger) (*Server, error) {
	s := &Server{
		Router: configureEcho(),
		logger: logger,
	}
	return s, nil
}

// Run http server
func (s *Server) Run(addr string) error {
	s.logger.Info(fmt.Sprintf("http server starting at port %s", addr))
	err := s.Router.Start(addr)
	if err != nil {
		return errors.Wrap(err, "error run httpServer")
	}
	return nil
}

// configureEcho create echo router
func configureEcho() *echo.Echo {
	e := echo.New()
	e.Debug = false
	e.HTTPErrorHandler = middleware.ErrorHandler()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.SetTraceId())
	e.Use(middleware.RequestLogger([]string{}))
	return e
}

// AddEndpoints add endpoints to the router with middleware
func (s *Server) AddEndpoints(endpoints []*Endpoint, middleware ...echo.MiddlewareFunc) {
	for _, endpoint := range endpoints {
		s.Router.Add(endpoint.Method, endpoint.Path, wrapHandler(endpoint.Handle), middleware...)
	}
}

// AddEndpointsGroup add endpoints group with prefix to the router
func (s *Server) AddEndpointsGroup(prefix string, endpoints []*Endpoint, middleware ...echo.MiddlewareFunc) {
	g := s.Router.Group(prefix, middleware...)
	for _, endpoint := range endpoints {
		g.Add(endpoint.Method, endpoint.Path, wrapHandler(endpoint.Handle), middleware...)
	}
}
