package find_by_id

import (
	"hexample.com/src/oleander/branch_location/application/dto"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/shared/shared_domain"
)

type FindBYIDBranchLocationUC struct {
	r domain.BranchLocationRepository
	dms *domain.FindByIDBranchLocationDS
}

func NewFindBYIDBranchLocationUC(r domain.BranchLocationRepository) *FindBYIDBranchLocationUC {
	return &FindBYIDBranchLocationUC{
		r: r,
		dms: domain.NewFindByIDBranchLocationDS(r),
	}
}

func (uc *FindBYIDBranchLocationUC) Invoke(id *shared_domain.BranchLocationIDValueVO) (bl *dto.BranchLocationDTO, rErr error) {
	ag, err := uc.dms.Invoke(id)
	if err != nil {
		return nil, err
	}

	bl = dto.NewBranchLocationDTO(ag)

	// TODO: Emit events

	return bl, nil
}

