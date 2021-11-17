package httpserver

import (
	"eventer/config"
	"eventer/core/logger"

	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/fx"
	"net/http"
)

type server struct {
	server *http.Server
	log    logger.Logger

	port     int
	host     string
	handlers API
}

func Make(
	cfg *config.Config,
	log logger.Logger,
	handlers API,
) Server {
	return &server{
		log:      log,
		port:     cfg.HTTPServer.Port,
		host:     cfg.HTTPServer.Host,
		handlers: handlers,
	}
}

var Module = fx.Options(
	fx.Provide(Make),
	fx.Invoke(Register),
)

func Register(s Server) {
	go s.Register()
}

func (s *server) Register() {
	ctx := context.Background()

	router := chi.NewRouter()

	// TODO: define custom middleware.Logger with logger.LoggerInterface
	router.Use(middleware.Logger)

	// cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           1000, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/api/health-check", MakeHandlesChain(AppendTracingContext, s.handlers.HealthCheck))

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:           router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	s.log.Info(ctx, "HTTP server up and running on: ", map[string]interface{}{"server_port": s.port, "server_host": s.host})

	if err := s.server.ListenAndServe(); err != nil {
		s.log.Error(ctx, "Register http server error", map[string]interface{}{"error": err.Error()})
	}
}
