package infrastructure

import (
	"fmt"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/criteria"
)

type DummyUserRepository struct {
}

func NewDummyUserRepository() *DummyUserRepository {
	return &DummyUserRepository{}
}

func (r *DummyUserRepository) Begin() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Begin \n")
	return nil
}

func (r *DummyUserRepository) Commit() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Commit \n")

	return nil
}

func (r *DummyUserRepository) Rollback() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - RollBack \n")
	return nil
}

func (r *DummyUserRepository) Count() (int, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Count \n")
	return 0, nil
}

func (r *DummyUserRepository) SearchAll() ([]*domain.UserAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - SearchAll \n")
	return nil, nil
}

func (r *DummyUserRepository) SearchBy(criteria *criteria.Criteria) ([]*domain.UserAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - SearchBy \n")
	return nil, nil
}

func (r *DummyUserRepository) SearchByID(id shared_domain.UserIDValueVO) (*domain.UserAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - SearchByID \n")
	return nil, nil
}

func (r *DummyUserRepository) Save(u *domain.UserAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Save \n")
	return nil
}

func (r *DummyUserRepository) Update(u *domain.UserAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Update \n")
	return nil
}

func (r *DummyUserRepository) Delete(u *domain.UserAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyUserRepository] - Delete \n")
	return nil
}
