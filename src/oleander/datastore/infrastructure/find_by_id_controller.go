package infrastructure

import (
	"hexample.com/src/oleander/datastore/application/dto"
	"hexample.com/src/oleander/datastore/application/get_id"
	"hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/shared/shared_domain"
)

type GetDatastoreByIdController struct {
	uc *get_id.GetDatastoreByIdUc
}

func NewGetDatastoreByIdController(r domain.DatastoreRepository) *GetDatastoreByIdController {
	return &GetDatastoreByIdController{
		uc: get_id.NewGetDatastoreByIdUc(r),
	}
}

func (c *GetDatastoreByIdController) Invoke(id string) (*dto.DatastoreDTO, error) {
	idVo, err := shared_domain.NewDatastoreIDValueVO(id)
	if err != nil {
		return nil, err
	}
	dto, err := c.uc.Invoke(*idVo)
	if err != nil {
		return nil, err
	}
	return dto, nil
}