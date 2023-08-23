package healthCheck

import (
	"server/modules/core"

	"github.com/gofiber/fiber/v2"
)

func healthy(c *fiber.Ctx) error {
	newResp := core.SuccessResponse[string]{}.New(c)
	return newResp.Default("Healthy!")
}
