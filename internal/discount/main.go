package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/noodlensk/discount-demo/internal/common/logs"
	"github.com/noodlensk/discount-demo/internal/common/server"
	"github.com/noodlensk/discount-demo/internal/discount/ports"
	"github.com/noodlensk/discount-demo/internal/discount/service"
)

func main() {
	ctx := context.Background()

	logger := logs.NewLogger()

	app := service.NewApplication(ctx, logger)

	if err := server.RunHTTPServer(":8080", logger, func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHTTPServer(app), router)
	}); err != nil {
		logger.Fatal(err)
	}
}
