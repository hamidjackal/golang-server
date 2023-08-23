package healthCheck

import (
	"github.com/gofiber/fiber/v2"
)

func Router(api fiber.Router) {
	router := api.Group("/health-check")
	router.Get("/", healthy)
}
