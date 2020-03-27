package main

import (
	"github.com/labstack/echo"
	"hexample.com/src/user/infrastructure"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	c := infrastructure.NewGetUserByIDEchoController()

	e.GET("/user/:id", c.Invoke)

	e.POST("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.PATCH("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.DELETE("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
