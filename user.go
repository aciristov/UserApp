package main

import (
	"goa-user/app"
	"goa-user/db"

	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	Repository db.UserRepository
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, repo db.UserRepository) *UserController {
	return &UserController{Controller: service.NewController("UserController"), Repository: repo}
}

// CreateUser runs the CreateUser action.
func (c *UserController) CreateUser(ctx *app.CreateUserUserContext) error {
	// UserController_CreateUser: start_implement

	// Put your logic here

	dbUser := c.Repository.CreateUser(*ctx.Payload.Username, *ctx.Payload.Password, *ctx.Payload.Email)

	// UserController_CreateUser: end_implement

	res := &app.ExampleUser{
		Email:	&dbUser.Email,
		Username: &dbUser.Username,
		ID:	&dbUser.ID,
	}
	return ctx.OK(res)
}

// GetUser runs the GetUser action.
func (c *UserController) GetUser(ctx *app.GetUserUserContext) error {
	// UserController_GetUser: start_implement

	dbUser := c.Repository.GetUserById(ctx.UserID)

	if dbUser == nil{
		return ctx.NotFound()
	}

	// Put your logic here

	// UserController_GetUser: end_implement

	res := &app.ExampleUser{
		Username: &dbUser.Username,
		Email : &dbUser.Email,
		ID: &dbUser.ID,
	}
	return ctx.OK(res)
}
