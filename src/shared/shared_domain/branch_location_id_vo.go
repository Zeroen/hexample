package shared_domain

type BranchLocationIDValueVO struct {
	value UUIDValueVO
}

func NewBranchLocationIDValueVO (value string) (*BranchLocationIDValueVO, error) {
	u, err := NewUUIDValueVO(value)
	if err != nil {
		return nil, err
	}

	return &BranchLocationIDValueVO{
		value: *u,
	}, nil
}

func (e *BranchLocationIDValueVO) GetValue() string {
	return e.value.GetValue()
}

func (e *BranchLocationIDValueVO) String() string {
	return e.value.GetValue()
}