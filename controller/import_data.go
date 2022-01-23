package controller

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"petProject/logger"
	"petProject/service"
)


const ImportUrl = "https://jsonplaceholder.typicode.com/todos/1"

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

// ImportDataController struct to import json file
type ImportDataController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

func NewImportDataController(
	ctx context.Context,
	services *service.Manager,
	logger *logger.Logger,
) *ImportDataController {
	return &ImportDataController{ctx: ctx, services: services, logger: logger}
}

func (c *ImportDataController) Import(ctx echo.Context) error {
	resp, err := http.Get(ImportUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	defer resp.Body.Close()

	data := make([]byte, 1024)
	n, err := resp.Body.Read(data)
	if err != nil {
		if err != io.EOF {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	// fmt.Println("number of bytes read", n)
	// fmt.Println("string(data)", string(data))

	var todo Todo
	if err = json.Unmarshal(data[:n], &todo); err != nil {
		// fmt.Println("Unmarshal err", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	// fmt.Println("todo", todo)

	return ctx.JSON(http.StatusOK, todo)
}
