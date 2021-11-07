package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/noodlensk/discount-demo/internal/common/auth"
	"github.com/noodlensk/discount-demo/internal/common/logs"
)

func RunHTTPServer(addr string, logger *zap.SugaredLogger, createHandler func(router chi.Router) http.Handler) error {
	return RunHTTPServerOnAddr(addr, logger, createHandler)
}

func RunHTTPServerOnAddr(addr string, logger *zap.SugaredLogger, createHandler func(router chi.Router) http.Handler) error {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter, logger)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	logger.Infof("Starting HTTP server on %q", addr)

	return http.ListenAndServe(addr, rootRouter)
}

func setMiddlewares(router *chi.Mux, logger *zap.SugaredLogger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
		middleware.NoCache,
	)

	router.Use(logs.NewHTTPMiddleware(logger))

	router.Use(auth.NewStaticTokenHTTPMiddleware().Middleware)
}
