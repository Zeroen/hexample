package test

import (
	"github.com/gemalto/requester"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"hexample.com/src/oleander/branch_location/infrastucture"
	"net/http"
	"testing"
	"time"
)

func TestNewCreateBranchLocationE2E(t *testing.T) {
	// TC - Arrange
	e := echo.New()

	blr := infrastucture.NewDummyBranchLocationRepository()
	blCreate := infrastucture.NewCreateBranchLocationController(blr)

	blDTO :=  &infrastucture.CreateBranchLocationDTO{
		ID:      "76ff36b5-11cf-49da-9fb9-1e28b942c328",
		Name:    "DTOName",
		City:    "DTOCity",
		Country: "DTOCountry",
	}

	e.POST("/branchlocation/:id", func(c echo.Context) error {
		err := blCreate.Invoke(blDTO)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Starting echo server
	go func () {
		err := e.Start(":1323")
		require.NoError(t, err)
	}()
	defer e.Server.Close()

	// TC - Action
	time.Sleep( 5 * time.Second)
	h, err := requester.Send(
		requester.Post("http://localhost:1323/branchlocation"),
		requester.AppendPath(blDTO.ID),
		requester.Body(blDTO))
	require.NoError(t,err)

	// TC - Assertion
	require.Equal(t, h.StatusCode, http.StatusOK)
}