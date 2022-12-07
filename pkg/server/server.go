package server

import (
	"context"
	"github.com/cyberwo1f/go-example-api/pkg/handler"
	"github.com/cyberwo1f/go-example-api/pkg/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
)

type Config struct {
	Log *zap.Logger
}

type Server struct {
	Mux        *mux.Router
	server     *http.Server
	handler    *handler.Handler
	middleware *middleware.Middleware
	log        *zap.Logger
}

func NewServer(registry *handler.Handler, middleware *middleware.Middleware, cfg *Config) *Server {
	if cfg == nil {
		log.Fatalf("config is nil")
	}

	s := &Server{
		Mux:        mux.NewRouter(),
		log:        cfg.Log,
		handler:    registry,
		middleware: middleware,
	}

	s.registerHandler()
	return s
}

func (s *Server) Serve(listener net.Listener) error {
	server := &http.Server{
		Handler: s.Mux,
	}
	s.server = server
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandler() {
	// global middleware
	s.Mux.Use(s.middleware.CorsMiddleware)

	// common
	s.Mux.Methods(http.MethodOptions).Handler(s.options())
	s.Mux.Methods(http.MethodGet).Path("/version").Handler(s.handler.Version.GetVersion())
	s.Mux.Methods(http.MethodGet).Path("/healthz").Handler(s.healthCheckHandler())

	// v1
	s.Mux.Methods(http.MethodGet).Path("/user/list").Handler(s.handler.V1.GetUsers())
	s.Mux.Methods(http.MethodGet).Path("/message/list/{id:[0-9]+}").Handler(s.handler.V1.GetMessages())
}

func (s *Server) healthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) options() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("requested option")
		w.WriteHeader(http.StatusOK)
	})
}
