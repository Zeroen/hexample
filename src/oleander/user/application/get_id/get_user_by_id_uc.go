package get_id

import (
	"fmt"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
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
	fmt.Printf("[APPLICATION] | [SERVICE] | GetUserByIDUC \n")

	user, err := uc.searchUserDS.Invoke(id)
	if err != nil {
		return nil, err
	}

	dto = NewUserDTO(user)

	return dto, nil
}
