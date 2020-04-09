package main

import (
	"github.com/labstack/echo"
	"hexample.com/src/oleander/user/application/reset_password"
	"hexample.com/src/oleander/user/infrastructure"
	"hexample.com/src/shared/shared_infrastructure/shared_infra_event_bus"
	"net/http"
)

func main() {
	e := echo.New()

	ur := infrastructure.NewDummyUserRepository()

	/*eventBus, err := shared_infra_event_bus.NewRabbitEventBus(
		"rabbitmq",
		"rabbitmq",
		"localhost",
		5672,
		nil)
	if err != nil {
		panic(err)
	}*/

	eventBus, err := shared_infra_event_bus.NewMemoryEventBus(nil)
	if err != nil {
		panic(err)
	}

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

	// Register event
	rpoucd := reset_password.NewResetPasswordUCOnUserCreatedDomainEvent(ur)
	err = eventBus.RegisterOnEventBus(rpoucd)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
