package infrastucture

import (
	"hexample.com/src/oleander/branch_location/application/create"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/oleander/branch_location/domain/vo"
	"hexample.com/src/shared/shared_domain"
)

type CreateBranchLocationController struct {
	uc *create.CreateBranchLocationUC
}

type CreateBranchLocationDTO struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

func NewCreateBranchLocationController(r domain.BranchLocationRepository) *CreateBranchLocationController {
	return &CreateBranchLocationController{
		uc: create.NewCreateBranchLocationUC(r),
	}
}

func (c *CreateBranchLocationController) Invoke(dto *CreateBranchLocationDTO) error {
	blName, err := vo.NewNameVO(dto.Name)
	if err != nil {
		return err
	}
	blCity := vo.NewCityVO(dto.City)
	blCountry := vo.NewCountryVO(dto.Country)
	blID, err := shared_domain.NewBranchLocationIDValueVO(dto.ID)
	if err != nil {
		return err
	}
	err = c.uc.Invoke(blID, blName, blCity, blCountry)
	if err != nil {
		return err
	}
	return nil
}
