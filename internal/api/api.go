package api

import (
	"github.com/aguerram/gtcth/config"
	route "github.com/aguerram/gtcth/internal/api/routes"
	"github.com/aguerram/gtcth/internal/api/services"
	"github.com/gofiber/fiber/v2"
)

func InitializeApi(env *config.AppEnv, group fiber.Router) {
	service := services.NewService()
	route.InitializeRoutes(env, group, service)
}
