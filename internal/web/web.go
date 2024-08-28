package web

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/web/routes"
	"github.com/aguerram/gtcth/internal/web/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func InitializeWebApp(env *config.AppEnv, db *pgx.Conn, group fiber.Router) {
	webServices := services.NewService(db)
	//initialize services

	//register routes
	route.InitializeRoutes(env, group, webServices)
}
