package get_id

import (
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/user/domain"
)

type GetUserByIDUC struct {

}

func NewGetUserByIDUC() *GetUserByIDUC {

	return &GetUserByIDUC{

	}
}

func (uc *GetUserByIDUC) Invoke(id *shared_domain.UserIDValueVO) *domain.UserAG {
	return nil
}