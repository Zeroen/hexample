package dto

import "hexample.com/src/oleander/branch_location/domain"

type BranchLocationDTO struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

func NewBranchLocationDTO (ag *domain.BranchLocationAG) *BranchLocationDTO{
	return &BranchLocationDTO{
		ID:      ag.GetID(),
		Name:    ag.GetName(),
		City:    ag.GetCity(),
		Country: ag.GetCountry(),
	}
}

