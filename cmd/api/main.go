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
	importDataController := controller.NewImportDataController(ctx, serviceManager, l)

	// Initialize Echo instance
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Disable Echo JSON logger in debug mode
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	// API V1
	v1 := e.Group("/api/v1")

	// User routes
	userRoutes := v1.Group("/users")
	userRoutes.GET("/:id", userController.Get)
	userRoutes.POST("/", userController.Create)
	userRoutes.PUT("/:id", userController.Update)
	userRoutes.DELETE("/:id", userController.Delete)

	// import Json API
	v1.GET("/importData", importDataController.Import)

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
