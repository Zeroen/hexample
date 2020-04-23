package infrastructure

import (
	"fmt"
	domain2 "hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/shared/shared_domain"
)

type DummyDatastoreRepository struct {
}

func NewDummyDatastoreRepository() *DummyDatastoreRepository {
	return &DummyDatastoreRepository{}
}

func (r *DummyDatastoreRepository) Begin() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - Begin \n")
	return nil
}

func (r *DummyDatastoreRepository) Commit() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - Commit \n")

	return nil
}

func (r *DummyDatastoreRepository) Rollback() error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - RollBack \n")
	return nil
}

func (r *DummyDatastoreRepository) SearchAll() ([]*domain2.DatastoreAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - SearchAll \n")
	return nil, nil
}

func (r *DummyDatastoreRepository) SearchByID(id shared_domain.UserIDValueVO) (*domain2.DatastoreAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - SearchByID \n")
	return nil, nil
}

func (r *DummyDatastoreRepository) Save(u *domain2.DatastoreAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - Save \n")
	return nil
}

