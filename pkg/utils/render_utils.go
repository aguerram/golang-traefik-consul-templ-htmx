package utils

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(ctx *fiber.Ctx, component templ.Component) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return component.Render(ctx.Context(), ctx.Response().BodyWriter())
}
