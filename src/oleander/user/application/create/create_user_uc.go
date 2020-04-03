package create

import (
	"errors"
	"fmt"
	"hexample.com/src/oleander/shared/shared_domain"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/oleander/user/domain/vo"
)

type CreateUserUC struct {
	r domain.UserRepository
	searchUserDS *domain.GetUserByIDDS
}

func NewCreateUserUC(r domain.UserRepository) *CreateUserUC {
	return &CreateUserUC{
		r: r,
		searchUserDS: domain.NewGetUserByIDDS(r),
	}
}

func (uc *CreateUserUC) Invoke(id shared_domain.UserIDValueVO, name *vo.NameVO, age *vo.AgeVO, emailID shared_domain.EmailIDValueVO) (rErr error) {

	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	user, err := uc.searchUserDS.Invoke(id)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New(fmt.Sprintf("User already exists. %s", id))
	}

	user, err = domain.NewUser(id, name, age, emailID)
	if err != nil {
		return err
	}

	err = uc.r.Save(user)
	if err != nil {
		return err
	}

	// EMIT EVENT

	return nil
}
