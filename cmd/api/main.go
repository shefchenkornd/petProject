package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	user "petProject/model"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// Routing
	g := e.Group("/api")
	g.GET("/users", getUsers)
	//g.GET("/users/:id", getUser)
	g.POST("/users", saveUser)
	//g.PUT("/users/:id", updateUser)
	//g.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func run() error {

	return nil
}

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "all users!")
}

func saveUser(c echo.Context) error {
	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
