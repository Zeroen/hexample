package domain

import (
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/user/domain/vo"
)

type UserAG struct {
	id shared_domain.UserIDValueVO

	name *vo.NameVO
	age  *vo.AgeVO

	email shared_domain.EmailIDValueVO
}

func NewUser(
	id shared_domain.UserIDValueVO,
	name *vo.NameVO,
	age *vo.AgeVO,
	email shared_domain.EmailIDValueVO) (*UserAG, error) {

	return &UserAG{
		id:    id,
		name:  name,
		age:   age,
		email: email,
	}, nil
}
