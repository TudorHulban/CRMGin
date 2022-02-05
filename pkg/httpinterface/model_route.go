package httpinterface

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Route Concentrates information to define a Gin route.
type route struct {
	Group    string
	Endpoint string
	Method   string          // HTTP Method
	Handler  gin.HandlerFunc // could maybe be another type from another framework
}

// registerRoute Method adds one route to Gin internal router.
func (s *HTTPServer) registerRoute(r route) error {
	r.Method = strings.ToTitle(r.Method)

	go s.cfg.GLogger.Debugf("Adding Route: %v, Method: %v", r.Group+r.Endpoint, r.Method)

	switch r.Method {
	case http.MethodGet:
		s.engine.GET(r.Group+r.Endpoint, r.Handler)
	case http.MethodPost:
		s.engine.POST(r.Group+r.Endpoint, r.Handler)
	case http.MethodPut:
		s.engine.PUT(r.Group+r.Endpoint, r.Handler)
	case http.MethodPatch:
		s.engine.PATCH(r.Group+r.Endpoint, r.Handler)
	case http.MethodDelete:
		s.engine.DELETE(r.Group+r.Endpoint, r.Handler)
	// could be used for grpc
	case "ANY":
		s.engine.Any(r.Group+r.Endpoint, r.Handler)
	default:
		return errors.New("unsupported method: " + r.Method)
	}

	return nil
}

// RegisterRoutes Registers the routes passed.
func (s *HTTPServer) registerRoutes(routes []route) error {
	if len(routes) == 0 {
		return errors.New("no routes to add")
	}

	go s.cfg.GLogger.Debugf("Routes to add: %v", routes)

	for _, route := range routes {
		if errReg := s.registerRoute(route); errReg != nil {
			return errReg
		}
	}

	return nil
}
