package core

import "github.com/gofiber/fiber/v2"

type SuccessResponse[T any] struct {
	ctx *fiber.Ctx
}

func (s SuccessResponse[T]) New(c *fiber.Ctx) SuccessResponse[T] {
	return SuccessResponse[T]{
		ctx: c,
	}
}

func (s SuccessResponse[T]) Default(result T) error {
	return s.ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  result,
	})
}

func (s SuccessResponse[T]) List(recs []T) error {
	return s.ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  recs,
	})
}

func (s SuccessResponse[T]) Create(rec T) error {
	return s.ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"result":  rec,
	})
}
