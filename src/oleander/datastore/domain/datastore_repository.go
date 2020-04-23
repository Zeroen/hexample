package domain

import (
	"hexample.com/src/shared/shared_domain"
)

type DatastoreRepository interface {
	shared_domain.Transactional
	SearchAll() ([]*DatastoreAG, error)
	SearchByID(id shared_domain.UserIDValueVO) (*DatastoreAG, error)
	Save(u *DatastoreAG) error
}
