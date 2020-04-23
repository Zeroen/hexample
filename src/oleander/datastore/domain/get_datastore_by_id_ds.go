package domain

import (
	"hexample.com/src/shared/shared_domain"
)

type GetDatastoreByIdDs struct {
	r DatastoreRepository
}

func NewGetDatastoreByIdDs(r DatastoreRepository) *GetDatastoreByIdDs {
	return &GetDatastoreByIdDs{
		r: r,
	}
}

func (uc *GetDatastoreByIdDs) Invoke(vo shared_domain.DatastoreIDValueVO) (*DatastoreAG, error){
	datastore, err := uc.r.SearchByID(vo)
	if err != nil {
		return nil, err
	}
	return datastore, nil
}