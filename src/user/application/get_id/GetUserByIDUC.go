package get_id

import (
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/user/domain"
)

type GetUserByIDUC struct {
	r domain.UserRepository
}

func NewGetUserByIDUC(r domain.UserRepository) *GetUserByIDUC {
	return &GetUserByIDUC{
		r: r,
	}
}

func (uc *GetUserByIDUC) Invoke(id *shared_domain.UserIDValueVO) (ag *domain.UserAG, rErr error) {

	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return nil, err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	return uc.r.SearchByID(id)
}
