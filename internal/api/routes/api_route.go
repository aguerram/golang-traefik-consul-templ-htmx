package routes

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/api/handlers"
	"github.com/aguerram/gtcth/internal/api/services"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(env *config.AppEnv, router fiber.Router, services *services.ApiService) {
	//routes
	healthHandler := handlers.NewHealthHandler(env, services.HealthService)
	router.Get("/health", healthHandler.GetHealth)
}
