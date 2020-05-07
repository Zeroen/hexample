package infrastucture

import (
	"hexample.com/src/oleander/branch_location/application/dto"
	"hexample.com/src/oleander/branch_location/application/find_by_id"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/shared/shared_domain"
)

type FindByIDBranchLocationController struct {
	uc *find_by_id.FindBYIDBranchLocationUC
}

type FindByIDBranchLocationDTO struct {
	ID      string `json:"id,omitempty"`
}

func NewFindByIDBranchLocationController(r domain.BranchLocationRepository) *FindByIDBranchLocationController {
	return &FindByIDBranchLocationController{
		uc: find_by_id.NewFindBYIDBranchLocationUC(r),
	}
}

func (c *FindByIDBranchLocationController) Invoke(dto *FindByIDBranchLocationDTO) (*dto.BranchLocationDTO, error) {
	blID, err := shared_domain.NewBranchLocationIDValueVO(dto.ID)
	if err != nil {
		return nil, err
	}
	res, err := c.uc.Invoke(blID)
	if err != nil {
		return nil, err
	}
	return res, nil
}