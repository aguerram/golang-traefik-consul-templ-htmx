package api

import (
	"github.com/aguerram/gtcth/config"
	route "github.com/aguerram/gtcth/internal/api/routes"
	"github.com/aguerram/gtcth/internal/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func InitializeApi(env *config.AppEnv, connection *pgx.Conn, group fiber.Router) {
	service := services.NewService(connection)
	route.InitializeRoutes(env, group, service)
}
