package vo

import "errors"

type NameVO struct {
	value string
}

func NewNameVO(name string) (*NameVO, error) {

	err := NameVO{}.enforceNonEmptyName(name)
	if err != nil {
		return nil, err
	}

	return &NameVO{
		value: name,
	}, nil
}

func (vo *NameVO) enforceNonEmptyName(value string) error {
	if value == "" {
		return errors.New("User name cannot be empty.")
	}
	return nil
}

func (vo *NameVO) GetValue() string {
	return vo.value
}
