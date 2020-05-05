package infrastructure

import (
	"hexample.com/src/oleander/datastore/application/dto"
	"hexample.com/src/oleander/datastore/application/list"
	"hexample.com/src/oleander/datastore/domain"
)

type ListDatastoresController struct {
	uc *list.ListDatastoreUc
}

func NewListDatastoresController(repository domain.DatastoreRepository) *ListDatastoresController {
	return &ListDatastoresController{
		uc: list.NewListDatastoresUc(repository),
	}
}

func (c *ListDatastoresController) Invoke() ([]*dto.DatastoreDTO, error) {
	return c.uc.Invoke()
}