package domain

import (
	"fmt"
	"hexample.com/src/oleander/user/domain/events"
	"hexample.com/src/oleander/user/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type UserAG struct {
	shared_domain_event_bus.AggregateRoot

	id shared_domain.UserIDValueVO

	name     *vo.NameVO
	age      *vo.AgeVO
	password vo.PasswordVO

	email shared_domain.EmailIDValueVO
}

func NewUser(
	id shared_domain.UserIDValueVO,
	name *vo.NameVO,
	age *vo.AgeVO,
	email shared_domain.EmailIDValueVO) (*UserAG, error) {

	userAG := &UserAG{
		id:    id,
		name:  name,
		age:   age,
		email: email,
	}

	userAG.Record(events.NewUserCreatedDomainEvent(id.GetValue()))

	return userAG, nil
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

func (a *UserAG) ResetPassword() {
	fmt.Printf("[DOMAIN] | [USERAG] - RESET PASSWORD \n")

	a.password.CalculateNewPassword()
}
