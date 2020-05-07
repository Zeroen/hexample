package infrastucture

import (
	"fmt"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/shared/shared_domain"
)

type DummyBranchLocationRepository struct {
}

func NewDummyBranchLocationRepository() *DummyBranchLocationRepository {
	return &DummyBranchLocationRepository{}
}

func (r *DummyBranchLocationRepository) Begin() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyBranchLocationRepository] - Begin \n")
	return nil
}

func (r *DummyBranchLocationRepository) Commit() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyBranchLocationRepository] - Commit \n")

	return nil
}

func (r *DummyBranchLocationRepository) Rollback() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyBranchLocationRepository] - RollBack \n")
	return nil
}


func (r *DummyBranchLocationRepository) SearchByID(id *shared_domain.BranchLocationIDValueVO) (*domain.BranchLocationAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyBranchLocationRepository] - SearchByID \n")
	return nil, nil
}

func (r *DummyBranchLocationRepository) Save(u *domain.BranchLocationAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyBranchLocationRepository] - Save \n")
	return nil
}

