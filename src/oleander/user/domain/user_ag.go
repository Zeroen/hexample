package domain

import (
	"hexample.com/src/oleander/shared/shared_domain"
	"hexample.com/src/oleander/user/domain/vo"
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

func (a *UserAG) GetID() string {
	return a.id.GetValue()
}

func (a *UserAG) GetAge() int {
	return a.age.GetValue()
}

func (a *UserAG) GetName() string {
	return a.name.GetValue()
}

func (a *UserAG) GetEmail() string {
	return a.GetEmail()
}
