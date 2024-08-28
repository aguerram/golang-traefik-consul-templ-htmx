package handlers

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/api/services"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type HealthHandler struct {
	env           *config.AppEnv
	healthService services.HealthService
}

func NewHealthHandler(env *config.AppEnv, healthService services.HealthService) *HealthHandler {
	return &HealthHandler{
		env:           env,
		healthService: healthService,
	}
}

func (h *HealthHandler) GetHealth(ctx *fiber.Ctx) error {
	check, err := h.healthService.HealthCheck(ctx.UserContext())
	if err != nil {
		log.Errorf("Error checking health: %v", err)
		return err
	}
	if check.Status == "DOWN" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(check)
	}
	return ctx.JSON(check)
}
