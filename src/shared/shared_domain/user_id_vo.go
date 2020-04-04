package shared_domain

type UserIDValueVO struct {
	value UUIDValueVO
}

func NewUserIDValueVO(value string) (*UserIDValueVO, error) {
	u, err := NewUUIDValueVO(value)
	if err != nil {
		return nil, err
	}

	return &UserIDValueVO{
		value: *u,
	}, nil
}

func (e *UserIDValueVO) GetValue() string {
	return e.value.GetValue()
}

func (e *UserIDValueVO) String() string {
	return e.value.GetValue()
}
