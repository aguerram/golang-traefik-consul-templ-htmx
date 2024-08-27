package middleware

import (
	"errors"
	"github.com/aguerram/gtcth/internal/web/views/page"
	"github.com/aguerram/gtcth/pkg/utils"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func NewAppErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Panic: %v", r)
				c.Status(fiber.StatusInternalServerError)
				utils.Render(c, page.Error("Internal Server Error"))
			}
		}()
		err := c.Next()
		if err != nil {
			//check if error is fiber error
			var e *fiber.Error
			var errorMessage string
			if errors.As(err, &e) {
				errorMessage = e.Message
				c.Status(e.Code)
			} else {
				errorMessage = "Internal Server Error"
				c.Status(fiber.StatusInternalServerError)
			}
			log.Errorf("Handler web error %s", errorMessage)
			return utils.Render(c, page.Error(errorMessage))
		}
		return err
	}
}
