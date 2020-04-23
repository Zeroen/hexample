package get_id

import (
	"hexample.com/src/oleander/datastore/application/dto"
	"hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/shared/shared_domain"
)

type GetDatastoreByIdUc struct {
	ds *domain.GetDatastoreByIdDs
}

func NewGetDatastoreByIdUc(r domain.DatastoreRepository) *GetDatastoreByIdUc {
	return &GetDatastoreByIdUc{
		ds: domain.NewGetDatastoreByIdDs(r),
	}
}

func (uc *GetDatastoreByIdUc) Invoke(vo shared_domain.DatastoreIDValueVO) (*dto.DatastoreDTO, error) {
	datastore, err := uc.ds.Invoke(vo)
	if err != nil {
		return nil, err
	}
	if datastore == nil {
		return nil, nil
	}
	dtoResult := dto.NewDatastoreDTO(datastore)
	return dtoResult, nil
}
