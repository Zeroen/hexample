package create

import (
	"errors"
	"hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/oleander/datastore/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type CreateDatastoreUc struct {
	r                      domain.DatastoreRepository
	GetDatastoreSearchById *domain.GetDatastoreByIdDs
	EventBus               shared_domain_event_bus.EventBus
}

func NewCreateDatastoreUc(r domain.DatastoreRepository, eb shared_domain_event_bus.EventBus) *CreateDatastoreUc {
	return &CreateDatastoreUc{
		r:                      r,
		GetDatastoreSearchById: domain.NewGetDatastoreByIdDs(r),
		EventBus:               eb,
	}
}

func (uc *CreateDatastoreUc) Invoke(id *shared_domain.DatastoreIDValueVO,
	name *vo.NameVO, path *vo.PathVO, hostname *vo.HostnameVO, port *vo.PortVO,
	user *shared_domain.UserIDValueVO) (rErr error) {
	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	// Check if the element already exists
	dsSearch, err := uc.GetDatastoreSearchById.Invoke(*id)
	if err != nil {
		return err
	}

	if dsSearch != nil {
		return errors.New("The datastore already exists")
	}

	ds, err := domain.NewDatastoreAG(id, hostname, name, path, port, user)
	if err != nil {
		return err
	}
	err = uc.r.Save(ds)
	if err != nil {
		return err
	}

	// Emit ds created event
	err = uc.EventBus.Publish(ds.PullDomainEvents())
	if err != nil {
		return err
	}

	return nil
}
