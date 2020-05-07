package domain

import (
	"hexample.com/src/shared/shared_domain"
)

type BranchLocationRepository interface {
	shared_domain.Transactional
	SearchByID(id *shared_domain.BranchLocationIDValueVO) (*BranchLocationAG, error)
	Save(u *BranchLocationAG) error
}
