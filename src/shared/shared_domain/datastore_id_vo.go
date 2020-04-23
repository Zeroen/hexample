package shared_domain

type DatastoreIDValueVO struct {
	value UUIDValueVO
}

func NewDatastoreIDValueVO(value string) (*DatastoreIDValueVO, error) {
	d, err := NewUUIDValueVO(value)
	if err != nil {
		return nil, err
	}

	return &DatastoreIDValueVO{
		value: *d,
	}, nil
}

func (e *DatastoreIDValueVO) GetValue() string {
	return e.value.GetValue()
}

func (e *DatastoreIDValueVO) String() string {
	return e.value.GetValue()
}
