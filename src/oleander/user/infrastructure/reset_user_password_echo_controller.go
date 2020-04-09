package infrastructure

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"hexample.com/src/oleander/user/application/reset_password"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"net/http"
)

type ResetUserPasswordEchoController struct {
	uc *reset_password.ResetPasswordUC
}

type ResetUserPasswordDTO struct {
	ID string `json:"elementID"`
}

func NewResetUserPasswordEchoController(r domain.UserRepository) *ResetUserPasswordEchoController {
	return &ResetUserPasswordEchoController{
		uc: reset_password.NewResetPasswordUC(r),
	}
}

func (c *ResetUserPasswordEchoController) Invoke(ctx echo.Context) error {

	fmt.Printf("[CONTROLLER] | [INFRASTRUCTURE] - ResetUserPasswordEchoController \n")

	var dto ResetUserPasswordDTO
	err := ctx.Bind(&dto)
	if err != nil {
		return errors.New("Error trying to read the ResetUserPasswordDTO" + err.Error())
	}

	id, err := shared_domain.NewUserIDValueVO(dto.ID)
	if err != nil {
		return err
	}

	err = c.uc.Invoke(*id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}
