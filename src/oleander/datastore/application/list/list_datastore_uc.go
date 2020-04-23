package list

import "hexample.com/src/oleander/datastore/domain"

type ListDatastoreUc struct {
	r domain.DatastoreRepository
}

func NewListDatastoresUc(r domain.DatastoreRepository) *ListDatastoreUc {
	return &ListDatastoreUc{
		r: r,
	}
}

func (uc *ListDatastoreUc) Invoke() ([]*DatastoreDTO, error){
	list, err := uc.r.SearchAll()
	if err != nil {
		return nil, err
	}

	return NewDatastoreDTOList(list), nil
}