package domain

import (
	"hexample.com/src/shared/shared_domain"
)

type FindByIDBranchLocationDS struct {
	r BranchLocationRepository
}

func NewFindByIDBranchLocationDS(r BranchLocationRepository) *FindByIDBranchLocationDS {
	return &FindByIDBranchLocationDS{
		r: r,
	}
}

func (uc *FindByIDBranchLocationDS) Invoke(id *shared_domain.BranchLocationIDValueVO) (*BranchLocationAG, error) {
	ag, err := uc.r.SearchByID(id)
	if err != nil {
		return nil, err
	}

	return ag, nil
}
