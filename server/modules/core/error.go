package core

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	ctx *fiber.Ctx
}

func (e ErrorResponse) New(c *fiber.Ctx) ErrorResponse {
	return ErrorResponse{
		ctx: c,
	}
}

func (e ErrorResponse) InternalServerError() error {
	return e.ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "Internal server error",
	})
}

func (e ErrorResponse) InvalidRequest(errs []string) error {
	return e.ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": errs,
	})
}
