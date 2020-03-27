package shared_domain

type EmailIDValueVO struct {
	value UUIDValueVO
}

func NewEmailIDValueVO(value string) (*EmailIDValueVO, error) {
	u, err := NewUUIDValueVO(value)
	if err != nil {
		return nil, err
	}

	return &EmailIDValueVO{
		value: *u,
	}, nil
}

func (e *EmailIDValueVO) GetValue() string {
	return e.value.GetValue()
}
