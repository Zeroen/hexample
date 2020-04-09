package infrastructure

import (
	"fmt"
	"github.com/labstack/echo"
	application "hexample.com/src/oleander/user/application/get_id"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"net/http"
)

type GetUserByIDEchoController struct {
	uc *application.GetUserByIDUC
}

func NewGetUserByIDEchoController() *GetUserByIDEchoController {
	var r domain.UserRepository
	return &GetUserByIDEchoController{
		uc: application.NewGetUserByIDUC(r),
	}
}

func (c *GetUserByIDEchoController) Invoke(ctx echo.Context) error {

	fmt.Printf("[CONTROLLER] | [INFRASTRUCTURE] - GetUserByIDEchoController \n")

	idParam := ctx.Param("id")

	id, err := shared_domain.NewUserIDValueVO(idParam)
	if err != nil {
		return err
	}

	user, err := c.uc.Invoke(*id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}
