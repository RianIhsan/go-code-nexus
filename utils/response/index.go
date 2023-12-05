package response

import (
	"github.com/gofiber/fiber/v2"
)

type GeneralMessage struct {
	Message string `json:"message"`
}

func SendStatusOkResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(GeneralMessage{
		Message: message,
	})
}

func SendStatusInternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(GeneralMessage{
		Message: message,
	})
}
