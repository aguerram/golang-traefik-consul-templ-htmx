package handler

import (
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal/view"
	"github.com/aguerram/gtcth/internal/view/page"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	env *config.AppEnv
}

func NewHomeHandler(env *config.AppEnv) *HomeHandler {
	return &HomeHandler{
		env: env,
	}
}

func (h *HomeHandler) GetHome(ctx *fiber.Ctx) error {
	return view.Render(ctx, page.Home())
}
func (h *HomeHandler) GetUserProfile(ctx *fiber.Ctx) error {
	profileId := ctx.Params("profileId")
	if profileId == "0" {
		return fiber.NewError(fiber.StatusBadRequest, "profileId is required")
	}
	return view.Render(ctx, page.UserProfile(profileId))
}
