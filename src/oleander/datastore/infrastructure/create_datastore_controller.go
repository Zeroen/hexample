package infrastructure

import (
	"fmt"
	"hexample.com/src/oleander/datastore/application/create"
	"hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/oleander/datastore/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type CreateDatastoreController struct {
	uc *create.CreateDatastoreUc
}

type CreateDatastoreDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

func NewCreateDatastoreController(r domain.DatastoreRepository, bus shared_domain_event_bus.EventBus) *CreateDatastoreController {
	return &CreateDatastoreController{
		uc: create.NewCreateDatastoreUc(r, bus),
	}
}

func (c *CreateDatastoreController) Invoke(dto *CreateDatastoreDTO, userId *shared_domain.UserIDValueVO) error {

	fmt.Printf("[CONTROLLER] | [INFRASTRUCTURE] - CreateDatastoreController \n")

	id, err := shared_domain.NewDatastoreIDValueVO(dto.ID)
	if err != nil {
		return err
	}

	name := vo.NewNameVO(dto.Name)
	path:= vo.NewPathVO(dto.Path)
	hostname := vo.NewHostnameVO(dto.Hostname)
	port := vo.NewPortVO(dto.Port)

	err = c.uc.Invoke(id, name, path, hostname, port, userId)
	if err != nil {
		return err
	}
	return nil
}
