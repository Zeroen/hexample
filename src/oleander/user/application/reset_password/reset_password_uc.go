package reset_password

import (
	"errors"
	"fmt"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
)

type ResetPasswordUC struct {
	r            domain.UserRepository
	searchUserDS *domain.GetUserByIDDS
}

func NewResetPasswordUC(r domain.UserRepository) *ResetPasswordUC {
	return &ResetPasswordUC{
		r:            r,
		searchUserDS: domain.NewGetUserByIDDS(r),
	}
}

func (uc *ResetPasswordUC) Invoke(id shared_domain.UserIDValueVO) (rErr error) {
	fmt.Printf("[APPLICATION] | [SERVICE] | ResetPasswordUC \n")

	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	user, err := uc.searchUserDS.Invoke(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(fmt.Sprintf("User does not exists. %s", id))
	}

	user.ResetPassword()

	err = uc.r.Update(user)
	if err != nil {
		return err
	}

	// TODO Publish events

	return nil
}
