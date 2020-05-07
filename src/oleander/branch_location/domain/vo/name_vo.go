package vo

import "errors"

type NameVO struct {
	value string
}

func NewNameVO (value string) (*NameVO, error){
	n := &NameVO{
		value: value,
	}

	err := n.enforceNumberChars()
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (n *NameVO) GetValue() string {
	return n.value
}

func (n *NameVO) enforceNumberChars() error {
	if len(n.value) <= 3 {
		return errors.New("name should have more than 3 chars")
	}
	return nil
}