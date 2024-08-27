package handler

import (
	"github.com/aguerram/gtcth/config"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	env *config.AppEnv
}

func NewHealthHandler(env *config.AppEnv) *HealthHandler {
	return &HealthHandler{
		env: env,
	}
}

func (h *HealthHandler) GetHealth(ctx *fiber.Ctx) error {
	return ctx.SendString("UP")
}
