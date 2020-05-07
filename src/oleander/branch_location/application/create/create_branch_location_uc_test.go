package create

import (
	"github.com/stretchr/testify/require"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/oleander/branch_location/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"testing"
)

func TestNewCreateBranchLocationUC(t *testing.T) {

}

func TestNewCreateBranchLocationUC_NoCreateDuplicatesID(t *testing.T) {
	// TC - Set up
	blName, err := vo.NewNameVO("Name1")
	require.NoError(t,err)
	blCity := vo.NewCityVO("City1")
	blCountry := vo.NewCountryVO("Country1")
	blID, err := shared_domain.NewBranchLocationIDValueVO("3b1e531e-5167-4c81-be07-f67df34f5fad")

	r := NewMockedBranchLocationRepository()
	r.On("Begin").Return(nil)
	r.On("SearchByID", blID).Return(domain.NewBranchLocationAG(blID, blName, blCity, blCountry), nil)
	r.On("Rollback").Return(nil)
	uc := NewCreateBranchLocationUC(r)

	// TC - Action
	err = uc.Invoke(blID, blName, blCity, blCountry)

	// TC - Assertion
	require.Error(t,err)
}


