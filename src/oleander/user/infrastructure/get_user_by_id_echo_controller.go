package infrastructure

import (
	"github.com/labstack/echo"
	"hexample.com/src/oleander/shared/shared_domain"
	application "hexample.com/src/oleander/user/application/get_id"
	"hexample.com/src/oleander/user/domain"
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
