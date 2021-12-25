package main

import (
	"context"
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"petProject/config"
	"petProject/controller"
	"petProject/logger"
	"petProject/service"
	"petProject/store"
	"time"
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
	store, err := store.New(ctx)
	if err != nil {
		return errors.New(err.Error())
	}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		errors.New("manager.New failed")
	}

	// Init controllers
	userController := controller.NewUsers(ctx, serviceManager, l)

	// Initialize Echo instance
	e := echo.New()

	// Disable Echo JSON logger in debug mode
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	// API V1
	v1 := e.Group("/v1")

	// User routes
	userRoutes := v1.Group("/users")
	userRoutes.GET("/users/:id", userController.Get)
	userRoutes.POST("/users", userController.Create)
	userRoutes.PUT("/users/:id", userController.Update)
	userRoutes.DELETE("/users/:id", userController.Delete)

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
