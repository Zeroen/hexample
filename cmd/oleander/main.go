package main

import (
	"github.com/labstack/echo"
	infraDs "hexample.com/src/oleander/datastore/infrastructure"
	"hexample.com/src/oleander/user/application/datastore_counter"
	"hexample.com/src/oleander/user/application/reset_password"
	"hexample.com/src/oleander/user/infrastructure"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
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
	incrementDsCounterController := infraDs.NewIncreaseDsCounterPutController(ur, eventBus)
	e.PUT("/user/:id/datastore/increment", func(context echo.Context) error {
		userId, err := shared_domain.NewUserIDValueVO(context.Param("id"))
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}
		err = incrementDsCounterController.Invoke(userId)
		if err != nil {
			return context.String(http.StatusInternalServerError, err.Error())
		}
		return nil
	})

	// Register datastore endpoints
	registerDatastoreEndpoints(e, eventBus)

	// Register subscribers
	rpoucd := reset_password.NewResetPasswordUCOnUserCreatedDomainEvent(ur)
	err = eventBus.RegisterOnEventBus(rpoucd)
	if err != nil {
		panic(err)
	}
	ndscsubs := datastore_counter.NewIncreaseDatastoreCounterOnDsCreatedEvent(ur, eventBus)
	err = eventBus.RegisterOnEventBus(ndscsubs)
	if err != nil {
		panic(err)
	}

	// Starting echo server
	e.Logger.Fatal(e.Start(":1323"))
}

func registerDatastoreEndpoints(e *echo.Echo, eventBus shared_domain_event_bus.EventBus) {
	dsRepo := infraDs.NewDummyDatastoreRepository()
	e.GET("/datastore", func(context echo.Context) error {
		c := infraDs.NewListDatastoresController(dsRepo)
		dsList, error := c.Invoke()
		if error != nil {
			return context.String(http.StatusBadRequest, error.Error())
		}

		return context.JSON(http.StatusOK, dsList)
	})
	e.POST("/datastore/:id", func(context echo.Context) error {
		id := context.Param("id")
		var dto infraDs.CreateDatastoreDTO
		err := context.Bind(&dto)
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}
		dto.ID = id
		userId, err := shared_domain.NewUserIDValueVO("056006bf-20e0-46f8-b63c-b2f5f32272b7")
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}
		err = infraDs.NewCreateDatastoreController(dsRepo, eventBus).Invoke(&dto, userId)
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}
		return context.JSON(http.StatusOK, nil)
	})
	e.GET("/datastore/:id", func(context echo.Context) error {
		id := context.Param("id")
		dto, err := infraDs.NewGetDatastoreByIdController(dsRepo).Invoke(id)
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}
		if dto == nil {
			return context.String(http.StatusNotFound, "Datastore not found")
		}

		return context.JSON(http.StatusOK, dto)
	})
}