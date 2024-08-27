package handlers

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/api/dto/response"
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
	check, err := h.healthService.HealthCheck()
	var status response.HealthCheckResponse
	if err != nil {
		log.Errorf("Error checking health: %v", err)
		status = response.HealthCheckResponse{
			Status: "DOWN",
		}
	}
	if !check {
		log.Errorf("Health check failed, some components are down")
		status = response.HealthCheckResponse{
			Status: "DOWN",
		}
	}
	status = response.HealthCheckResponse{
		Status: "UP",
	}
	return ctx.JSON(status)
}
