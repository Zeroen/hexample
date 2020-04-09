package infrastructure

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"hexample.com/src/oleander/user/application/create"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/oleander/user/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
	"net/http"
)

type CreateUserEchoController struct {
	uc *create.CreateUserUC
}

type CreateUserDTO struct {
	ID      string `json:"elementID"`
	Age     int    `json:"age"`
	Name    string `json:"name"`
	EmailID string `json:"emailID"`
}

func NewCreateUserEchoController(eventBus shared_domain_event_bus.EventBus, r domain.UserRepository) *CreateUserEchoController {
	return &CreateUserEchoController{
		uc: create.NewCreateUserUC(eventBus, r),
	}
}

func (c *CreateUserEchoController) Invoke(ctx echo.Context) error {

	fmt.Printf("[CONTROLLER] | [INFRASTRUCTURE] - CreateUserEchoController \n")

	var dto CreateUserDTO
	err := ctx.Bind(&dto)
	if err != nil {
		return errors.New("Error trying to read the CreateUSerDTO" + err.Error())
	}

	id, err := shared_domain.NewUserIDValueVO(dto.ID)
	if err != nil {
		return err
	}

	name, err := vo.NewNameVO(dto.Name)
	if err != nil {
		return err
	}

	age, err := vo.NewAgeVO(dto.Age)
	if err != nil {
		return err
	}

	emailID, err := shared_domain.NewEmailIDValueVO(dto.EmailID)
	if err != nil {
		return err
	}

	err = c.uc.Invoke(*id, name, age, *emailID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}
