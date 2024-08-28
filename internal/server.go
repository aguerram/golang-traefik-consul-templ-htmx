package internal

import (
	"context"
	"github.com/aguerram/gtcth/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GracefulShutdownHandler func(ctx context.Context)

func StartHttpServer(env *config.AppEnv) (*fiber.App, func(ctx context.Context)) {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/v1/health"
		},
	}))
	return app, func(ctx context.Context) {
		log.Info("Shutting down server")
		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Fatalf("Error shutting down server %v", err)
		} else {
			log.Info("Server successfully shutdown")
		}
	}
}

func HandleGracefulShutdowns(handlers ...GracefulShutdownHandler) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		log.Info("Received shutdown signal")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for _, handler := range handlers {
			handler(timeoutCtx) // You can pass shutdownCtx to handlers that require context for graceful shutdown
		}

		log.Info("Service successfully shutdown")
		os.Exit(0)
	}()
}
