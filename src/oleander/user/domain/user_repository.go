package domain

import (
	"hexample.com/src/oleander/shared/shared_domain"
	"hexample.com/src/oleander/shared/shared_domain/criteria"
)

type UserRepository interface {
	shared_domain.Transactional
	Count() (int, error)
	SearchAll() ([]*UserAG, error)
	SearchBy(criteria *criteria.Criteria) ([]*UserAG, error)
	SearchByID(id shared_domain.UserIDValueVO) (*UserAG, error)
	Save(u *UserAG) error
	Update(u *UserAG) error
	Delete(u *UserAG) error
}
