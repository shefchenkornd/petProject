package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"petProject/lib/types"
	"petProject/logger"
	"petProject/model"
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
		services: service,
		logger: logger,
	}
}


func (c UserController) Get(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}

	user, err := c.services.User.GetUser(ctx.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
		return echo.NewHTTPError(http.StatusOK, user)
	}

	return ctx.JSON(http.StatusOK,  user)
}

func (c UserController) Create(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}

	createdUser, err := c.services.User.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
		}
	}

	c.logger.Debug().Msgf("Created user '%s'", createdUser.ID.String())

	return ctx.JSON(http.StatusOK, createdUser)
}

func (c UserController) Update(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}

	updatedUser, err := c.services.User.UpdateUser(ctx.Request().Context(), &user)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
		}
	}

	c.logger.Debug().Msgf("Updated user '%s'", updatedUser.ID.String())

	return ctx.JSON(http.StatusOK, updatedUser)
}

func (c UserController) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not delete user"))
	}

	err = c.services.User.DeleteUser(ctx.Request().Context(), id)
	if err != nil {

	}

	return ctx.JSON(http.StatusOK, model.OK)
}