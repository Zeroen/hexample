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

	id1, err := shared_domain.NewDatastoreIDValueVO("fbf3287a-1032-4072-939b-e705c08fcd75")
	if err != nil {
		return nil, err
	}
	id2, err := shared_domain.NewDatastoreIDValueVO("fd49cef4-f317-4840-8dc0-164d539efd77")
	if err != nil {
		return nil, err
	}
	id3, err := shared_domain.NewDatastoreIDValueVO("fd49cef4-f317-4840-8dc0-164d539efd86")
	if err != nil {
		return nil, err
	}

	obj1, err := domain2.NewDatastoreAG(id1, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	obj2, err := domain2.NewDatastoreAG(id2, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	obj3, err := domain2.NewDatastoreAG(id3, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	datastore := []*domain2.DatastoreAG{
		obj1, obj2, obj3,
	}

	return datastore, nil
}

func (r *DummyDatastoreRepository) SearchByID(id shared_domain.DatastoreIDValueVO) (*domain2.DatastoreAG, error) {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - SearchByID \n")
	return nil, nil
}

func (r *DummyDatastoreRepository) Save(u *domain2.DatastoreAG) error {
	fmt.Printf("[DOMAIN] | [INVERSION] | [INFRASTRUCTURE] | [DummyDatastoreRepository] - Save \n")
	return nil
}

