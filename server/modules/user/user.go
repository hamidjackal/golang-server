package user

import (
	"server/modules/core"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Router(api fiber.Router) {
	controller := InitUserController()

	router := api.Group("/users")
	router.Post("/signin", controller.signin)
	router.Post("/signup", adaptor.HTTPMiddleware(core.Validate[SignUp]), controller.signup)

	router.Get("/:id", adaptor.HTTPMiddleware(core.AuthMiddleware), controller.getOne)
	router.Get("/", adaptor.HTTPMiddleware(core.AuthMiddleware), controller.list)
	router.Delete("/:id", adaptor.HTTPMiddleware(core.AuthMiddleware), controller.delete)
	router.Patch("/:id", adaptor.HTTPMiddleware(core.AuthMiddleware), controller.patch)
	router.Post("/",
		adaptor.HTTPMiddleware(core.AuthMiddleware),
		adaptor.HTTPMiddleware(core.Validate[SignUp]),
		controller.create,
	)

}
