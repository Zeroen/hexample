package domain

import (
	"errors"
	"hexample.com/src/oleander/datastore/domain/vo"
	"hexample.com/src/shared/shared_domain"
)

type DatastoreAG struct {
	id       *shared_domain.DatastoreIDValueVO
	name     *vo.NameVO
	path     *vo.PathVO
	hostname *vo.HostnameVO
	port     *vo.PortVO
}

func NewDatastoreAG(id *shared_domain.DatastoreIDValueVO, h *vo.HostnameVO,
	n *vo.NameVO, path *vo.PathVO, port *vo.PortVO) (*DatastoreAG, error) {
	if id == nil {
		return nil, errors.New("The data store id can't be null")
	}

	return &DatastoreAG{
		id: id,
		name: n,
		path: path,
		hostname: h,
		port: port,
	}, nil
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
