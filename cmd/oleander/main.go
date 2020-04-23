package main

import (
	"github.com/labstack/echo"
	infraDs "hexample.com/src/oleander/datastore/infrastructure"
	"hexample.com/src/oleander/user/application/reset_password"
	"hexample.com/src/oleander/user/infrastructure"
	"hexample.com/src/shared/shared_infrastructure/shared_infra_event_bus"
	"net/http"
)

func main() {
	e := echo.New()

	ur := infrastructure.NewDummyUserRepository()

	eventBus, err := shared_infra_event_bus.NewMemoryEventBus(nil)
	if err != nil {
		panic(err)
	}

	// Register User endpoints
	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	c := infrastructure.NewGetUserByIDEchoController()
	e.GET("/user/:id", c.Invoke)
	cCreate := infrastructure.NewCreateUserEchoController(eventBus, ur)
	e.POST("/user/:id", cCreate.Invoke)
	e.PATCH("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.DELETE("/user/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	cResetPassword := infrastructure.NewResetUserPasswordEchoController(ur)
	e.POST("/user/:id/password/reset", cResetPassword.Invoke)

	// Register datastore endpoints
	registerDatastoreEndpoints(e)

	// Register subscribers
	rpoucd := reset_password.NewResetPasswordUCOnUserCreatedDomainEvent(ur)
	err = eventBus.RegisterOnEventBus(rpoucd)
	if err != nil {
		panic(err)
	}

	// Starting echo server
	e.Logger.Fatal(e.Start(":1323"))
}

func registerDatastoreEndpoints(e *echo.Echo) {
	dsRepo := infraDs.NewDummyDatastoreRepository()
	e.GET("/datastore", func(context echo.Context) error {
		c := infraDs.NewListDatastoresController(dsRepo)
		dsList, error := c.Invoke()
		if error != nil {
			return context.String(http.StatusBadRequest, error.Error())
		}

		return context.JSON(http.StatusOK, dsList)
	})
}