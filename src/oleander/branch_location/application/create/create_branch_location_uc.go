package create

import (
	"errors"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/oleander/branch_location/domain/vo"
	"hexample.com/src/shared/shared_domain"
)

type CreateBranchLocationUC struct {
	r domain.BranchLocationRepository
	dms *domain.FindByIDBranchLocationDS
}

func NewCreateBranchLocationUC(r domain.BranchLocationRepository) *CreateBranchLocationUC {
	return &CreateBranchLocationUC{
		r: r,
		dms: domain.NewFindByIDBranchLocationDS(r),
	}
}

func (uc *CreateBranchLocationUC) Invoke(id *shared_domain.BranchLocationIDValueVO, name *vo.NameVO,
	city *vo.CityVO, country *vo.CountryVO) (rErr error) {
	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	ag, err := uc.dms.Invoke(id)
	if err != nil {
		return err
	}
	if ag != nil {
		return errors.New("branch Location ID already exists")
	}

	bl := domain.NewBranchLocationAG(id, name, city, country)
	err = uc.r.Save(bl)
	if err != nil {
		return err
	}

	// TODO: Emit events

	return nil
}
