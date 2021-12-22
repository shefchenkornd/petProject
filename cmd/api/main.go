package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"context"
	"petProject/config"
	"petProject/controller"
	"petProject/logger"
	"petProject/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	// logger
	l := logger.Get()

	// Init repository store (with mysql/postgresql inside)
	//store, err := store.New(ctx)
	//if err != nil {
	//	return errors.New(err)
	//}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		errors.New("manager.New failed")
	}

	// Init controllers
	userController := controller.NewUsers(ctx, serviceManager, l)

	e := echo.New()

	// API V1
	v1 := e.Group("/v1")

	// User routes
	userRoutes := v1.Group("/users")
	userRoutes.GET("/users", userController.Get)
	//v1.GET("/users/:id", getUser)
	userRoutes.POST("/users", userController.Create)
	//v1.PUT("/users/:id", updateUser)
	//v1.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}
