package shared_domain

type DatastoreIDValueVO struct {
	value UUIDValueVO
}

func NewDatastoreIDValueVO(value string) (*DatastoreIDValueVO, error) {
	// TODO Can this method be reused?
	u, err := NewUUIDValueVO(value)
	if err != nil {
		return nil, err
	}

	return &DatastoreIDValueVO{
		value: *u,
	}, nil
}

func (e *DatastoreIDValueVO) GetValue() string {
	return e.value.GetValue()
}

func (e *DatastoreIDValueVO) String() string {
	return e.value.GetValue()
}
