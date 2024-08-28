package handler

import (
	"github.com/aguerram/gtcth/internal/web/services"
	"github.com/aguerram/gtcth/internal/web/views/page/user"
	"github.com/aguerram/gtcth/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) ListUsers(ctx *fiber.Ctx) error {
	users, err := h.userService.ListUsers(ctx.Context())
	if err != nil {
		return err
	}
	return utils.Render(ctx, user.UserIndex(users))
}
