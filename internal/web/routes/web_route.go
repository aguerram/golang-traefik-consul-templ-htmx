package route

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/web/handlers"
	"github.com/aguerram/gtcth/internal/web/middlewares"
	"github.com/aguerram/gtcth/internal/web/services"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(env *config.AppEnv, router fiber.Router, services *services.WebService) {
	//register middlewares
	router.Use(middleware.NewAppErrorHandler())

	//home
	homeHandler := handler.NewHomeHandler(env)
	router.Get("/", homeHandler.GetHome)
	router.Get("/p/{profileId}", homeHandler.GetUserProfile)

	//user
	userHandler := handler.NewUserHandler(services.UserService)
	usersRoute := router.Group("/users")
	usersRoute.Get("/", userHandler.ListUsers)
}
