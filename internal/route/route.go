package route

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/handler"
	"github.com/aguerram/gtcth/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(env *config.AppEnv, app *fiber.App) {
	webRoutes(env, app.Group("/"))
}

func webRoutes(env *config.AppEnv, router fiber.Router) {
	//register middleware

	router.Use(middleware.NewAppErrorHandler())

	//health
	healthHandler := handler.NewHealthHandler(env)
	router.Get("/health", healthHandler.GetHealth)
	//home
	homeHandler := handler.NewHomeHandler(env)
	router.Get("/", homeHandler.GetHome)
	router.Get("/p/{profileId}", homeHandler.GetUserProfile)
}
