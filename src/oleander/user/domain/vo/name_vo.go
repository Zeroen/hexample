package vo

import "errors"

type NameVO struct {
	value string
}

func NewNameVO(name string) (*NameVO, error) {

	n := NameVO{
		value:name,
	}

	err := n.enforceNonEmptyName(name)
	if err != nil {
		return nil, err
	}

	return &n, nil
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
