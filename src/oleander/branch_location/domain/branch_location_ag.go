package domain

import (
	"hexample.com/src/oleander/branch_location/domain/vo"
	"hexample.com/src/shared/shared_domain"
)

type BranchLocationAG struct {
	id      *shared_domain.BranchLocationIDValueVO
	name    *vo.NameVO
	city    *vo.CityVO
	country *vo.CountryVO
}

func NewBranchLocationAG(id *shared_domain.BranchLocationIDValueVO, name *vo.NameVO,
	city *vo.CityVO, country *vo.CountryVO) *BranchLocationAG {
	return &BranchLocationAG{
		id:      id,
		name:    name,
		city:    city,
		country: country,
	}
}

func (b *BranchLocationAG)GetID() string {
	return b.id.GetValue()
}
func (b *BranchLocationAG)GetName () string {
	return b.name.GetValue()
}
func (b *BranchLocationAG)GetCity () string {
	return b.city.GetValue()
}
func (b *BranchLocationAG)GetCountry () string {
	return b.country.GetValue()
}
