package get_id

import (
	"hexample.com/src/oleander/shared/shared_domain"
	"hexample.com/src/oleander/user/domain"
)

type GetUserByIDUC struct {
	searchUserDS *domain.GetUserByIDDS
}

func NewGetUserByIDUC(r domain.UserRepository) *GetUserByIDUC {
	return &GetUserByIDUC{
		searchUserDS: domain.NewGetUserByIDDS(r),
	}
}

func (uc *GetUserByIDUC) Invoke(id shared_domain.UserIDValueVO) (dto *UserDTO, rErr error) {

	user, err := uc.searchUserDS.Invoke(id)
	if err != nil {
		return nil, err
	}

	dto = NewUserDTO(user)

	return dto, nil
}
