package domain

import (
	"fmt"
	"hexample.com/src/shared/shared_domain"
)

type GetUserByIDDS struct {
	r UserRepository
}

func NewGetUserByIDDS(r UserRepository) *GetUserByIDDS {
	return &GetUserByIDDS{
		r: r,
	}
}

func (ds *GetUserByIDDS) Invoke(id shared_domain.UserIDValueVO) (user *UserAG, rErr error) {
	fmt.Printf("[DOMAIN] | [SERVICE] | GetUserByIDDS \n")

	user, err := ds.r.SearchByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
