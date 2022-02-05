package httpinterface

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"

	"github.com/TudorHulban/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Config Concentrates attributes for starting a Gin server.
type Config struct {
	// for busy servers graceful shutdown period
	GracefulSecs uint8
	// IPV4 to run
	IPV4Address string
	// port to run
	Port uint16
	// binary version
	BinaryVersion string
	// logger to use with Gin
	GLogger *log.LogInfo
}

// CreateConfig Package Helper for creating Gin configuration.
// Holds also validators.
// socket like "0.0.0.0:8080"
func CreateConfig(socket, ldVersion string, logger *log.LogInfo, graceSeconds uint8) (*Config, error) {
	// input validation
	ipv4 := socket[:strings.Index(socket, ":")]

	// check is IPV4
	errParseIP := isIpv4(ipv4)
	if errParseIP != nil {
		return nil, errors.WithMessage(errParseIP, "provided Gin listening port could not be parsed")
	}

	port, errParsePort := strconv.Atoi(socket[strings.Index(socket, ":")+1:])
	if errParsePort != nil {
		return nil, errors.WithMessage(errParsePort, "provided Gin listening port could not be parsed")
	}

	return &Config{
		GracefulSecs:  graceSeconds,
		IPV4Address:   ipv4,
		Port:          uint16(port),
		BinaryVersion: ldVersion,
		GLogger:       logger,
	}, nil
}

// HTTPServer is HTTP server wrapper.
type HTTPServer struct {
	cfg       *Config
	isReady   func() bool // for probes
	engine    *gin.Engine
	crudLogic persistence.IUserCRUD // access CRUD persistance layer
}

// NewGinServer creates a new HTTP server.
// No cfg validation, use create config helper for that.
func NewGinServer(config *Config) (*HTTPServer, error) {
	s := new(HTTPServer)
	s.SetAsNotReady()
	s.cfg = config

	if s.cfg.GLogger.GetLogLevel() != log.DEBUG {
		go s.cfg.GLogger.Info("Setting Gin Log Level to Release Mode")
		gin.SetMode(gin.ReleaseMode) // to be called before initializing the router!
	}
	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.RedirectTrailingSlash = true
	s.engine.HandleMethodNotAllowed = false

	// adding routes
	if errRoInfra := s.registerRoutes(s.prepareInfraRoutes()); errRoInfra != nil {
		return nil, errors.WithMessage(errRoInfra, "could not prepare infrastructure routes")
	}

	if errRoLog := s.registerRoutes(s.prepareLoginRoute()); errRoLog != nil {
		return nil, errors.WithMessage(errRoLog, "could not prepare login routes")
	}

	if errRoCreateUser := s.registerRoutes(s.prepareCreateUserRoute()); errRoCreateUser != nil {
		return nil, errors.WithMessage(errRoCreateUser, "could not prepare create user routes")
	}

	// adding CRUD logic
	s.crudLogic = cgorm.NewUser(s.cfg.GLogger)

	return s, nil
}

// Run Method for starting HTTP Server.
func (s *HTTPServer) Run(ctx context.Context) error {
	gracefull := &http.Server{
		Handler: s.engine,
		Addr:    fmt.Sprintf("%s:%d", s.cfg.IPV4Address, s.cfg.Port),
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(ctx, time.Duration(int64(s.cfg.GracefulSecs))*time.Second)
		defer cancel()

		s.cfg.GLogger.Info("shutting down HTTP Server")
		if errShutdown := gracefull.Shutdown(ctx); errShutdown != nil {
			s.cfg.GLogger.Info(errShutdown, " - ", "could not gracefully stop HTTP server")
		}
	}()

	go s.cfg.GLogger.Infof("starting HTTP Server on port: %v", s.cfg.Port)
	return gracefull.ListenAndServe()
}

// SetAsReady Setter method for switching to readiness state READY.
func (s *HTTPServer) SetAsReady() {
	s.isReady = func() bool { return true }
}

// SetAsNotReady Setter method for switching to readiness state NOTREADY.
func (s *HTTPServer) SetAsNotReady() {
	s.isReady = func() bool { return false }
}

// GetIfReady Getter method for current readiness state.
func (s *HTTPServer) GetIfReady() bool {
	return s.isReady()
}
