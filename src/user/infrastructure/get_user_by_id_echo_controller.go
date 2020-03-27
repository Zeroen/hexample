package infrastructure

import (
	"github.com/labstack/echo"
	"hexample.com/src/shared/shared_domain"
)

type GetUserByIDEchoController struct {
}

func NewGetUserByIDEchoController() *GetUserByIDEchoController {
	return &GetUserByIDEchoController{}
}

func (c *GetUserByIDEchoController) Invoke(ctx echo.Context) error {

	idParam := ctx.Param("id")

	id, err := shared_domain.NewUserIDValueVO(idParam)
	if err != nil {
		return err
	}

	return nil
}
