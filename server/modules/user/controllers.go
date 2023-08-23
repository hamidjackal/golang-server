package user

import (
	"server/modules/core"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService UserService
	successResp core.SuccessResponse[User]
	errorResp   core.ErrorResponse
}

func (u UserController) New() UserController {
	return UserController{
		userService: UserService{}.New(),
		successResp: core.SuccessResponse[User]{},
	}
}

func (u UserController) list(c *fiber.Ctx) error {
	return u.userService.list(c)
}

func (u UserController) create(c *fiber.Ctx) error {
	p := new(SignUp)
	err := c.BodyParser(p)

	if err != nil {
		resp := u.errorResp.New(c)
		return resp.InternalServerError()
	}

	p.Password = core.HashPassword(p.Password)

	user, err := u.userService.create(p)
	if err != nil {
		resp := core.ErrorResponse{}.New(c)
		return resp.InvalidRequest([]string{err.Error()})
	}

	resp := u.successResp.New(c)

	return resp.Create(user)
}

func (u UserController) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.userService.delete(id)
	if err != nil {
		resp := core.ErrorResponse{}.New(c)
		return resp.InvalidRequest([]string{err.Error()})
	}

	resp := core.SuccessResponse[string]{}.New(c)
	return resp.Default("Deleted")
}

func (u UserController) signin(c *fiber.Ctx) error {
	p := new(SignIn)
	err := c.BodyParser(p)

	if err != nil {
		resp := u.errorResp.New(c)
		return resp.InternalServerError()
	}

	user, err := u.userService.signin(p)

	if err != nil {
		resp := u.errorResp.New(c)
		return resp.InvalidRequest([]string{err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  user,
	})
}

func (u UserController) signup(c *fiber.Ctx) error {
	p := new(SignUp)
	err := c.BodyParser(p)

	if err != nil {
		resp := u.errorResp.New(c)
		return resp.InternalServerError()
	}

	p.Password = core.HashPassword(p.Password)

	user, err := u.userService.create(p)
	if err != nil {
		resp := core.ErrorResponse{}.New(c)
		return resp.InvalidRequest([]string{err.Error()})
	}

	resp := u.successResp.New(c)

	return resp.Create(user)
}

func (u UserController) getOne(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.userService.getOne(id)
	if err != nil {
		resp := core.ErrorResponse{}.New(c)
		return resp.InvalidRequest([]string{err.Error()})
	}

	resp := u.successResp.New(c)
	return resp.Default(user)
}
