package web

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/web/routes"
	"github.com/aguerram/gtcth/internal/web/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func InitializeWebApp(env *config.AppEnv, connection *pgx.Conn, group fiber.Router) {
	//initialize services
	webServices := services.NewService()
	//register routes
	route.InitializeRoutes(env, group, webServices)
}
