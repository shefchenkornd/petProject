package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"petProject/logger"
	"petProject/service"
)

type UserController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewUsers creates a new user controller.
func NewUsers(ctx context.Context, service *service.Manager, logger *logger.Logger) *UserController {
	return &UserController{
		ctx: ctx,
		service: service,
		logger: logger,
	}
}


func (c UserController) Get(ctx echo.Context) error {
	c.services.GetUser()

	return ctx.JSON(http.StatusOK, "all users!")
}

func (c UserController) Create(ctx echo.Context) error {
	c.services.CreateUser()


	return ctx.JSON(http.StatusOK, "all users!")
}