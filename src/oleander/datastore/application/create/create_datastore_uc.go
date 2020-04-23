package create

import (
	"errors"
	"hexample.com/src/oleander/datastore/domain"
	"hexample.com/src/oleander/datastore/domain/vo"
	"hexample.com/src/shared/shared_domain"
)

type CreateDatastoreUc struct {
	r domain.DatastoreRepository
}

func NewCreateDatastoreUc(r domain.DatastoreRepository) *CreateDatastoreUc {
	return &CreateDatastoreUc{
		r: r,
	}
}

func (uc *CreateDatastoreUc) Invoke(id *shared_domain.DatastoreIDValueVO,
	name *vo.NameVO, path *vo.PathVO, hostname *vo.HostnameVO, port *vo.PortVO) (rErr error) {
	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	// Check if the element already exists
	dsSearch, err := uc.r.SearchByID(*id)
	if err != nil {
		return err
	}

	if dsSearch != nil {
		return errors.New("The datastore already exists")
	}

	ds, err := domain.NewDatastoreAG(id, hostname, name, path, port)
	if err != nil {
		return err
	}
	err = uc.r.Save(ds)
	if err != nil {
		return err
	}

	return nil
}