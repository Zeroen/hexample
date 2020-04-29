package domain

import (
	"errors"
	"hexample.com/src/oleander/datastore/domain/events"
	"hexample.com/src/oleander/datastore/domain/vo"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type DatastoreAG struct {
	shared_domain_event_bus.AggregateRoot
	id       *shared_domain.DatastoreIDValueVO
	name     *vo.NameVO
	path     *vo.PathVO
	hostname *vo.HostnameVO
	port     *vo.PortVO
}

func NewDatastoreAG(id *shared_domain.DatastoreIDValueVO, h *vo.HostnameVO,
	n *vo.NameVO, path *vo.PathVO, port *vo.PortVO, userId *shared_domain.UserIDValueVO) (*DatastoreAG, error) {
	if id == nil {
		return nil, errors.New("The data store id can't be null")
	}

	// TODO Here can go all required guard clause

	ar := &DatastoreAG{
		id: id,
		name: n,
		path: path,
		hostname: h,
		port: port,
	}

	ar.Record(events.NewDatastoreCreatedEvent(ar.GetID(), userId.GetValue()))

	return ar, nil
}

func (d *DatastoreAG) GetID() string {
	return d.id.GetValue()
}

func (d *DatastoreAG) GetName() string {
	if d.name == nil {
		return ""
	}
	return d.name.GetValue()
}

func (d *DatastoreAG) GetPath() string {
	if d.path == nil {
		return ""
	}
	return d.path.GetValue()
}

func (d *DatastoreAG) GetHostname() string {
	if d.hostname == nil {
		return ""
	}
	return d.hostname.GetValue()
}

func (d *DatastoreAG) GetPort() int {
	if d.port == nil {
		return 0
	}
	return d.port.GetValue()
}
