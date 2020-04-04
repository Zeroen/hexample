package shared_domain

import (
	"errors"
	"fmt"
	"regexp"
)

type UUIDValueVO struct {
	value string
}

func NewUUIDValueVO(value string) (*UUIDValueVO, error) {

	u := UUIDValueVO{
		value: value,
	}

	err := u.enforceUUIDformat(value)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (vo *UUIDValueVO) GetValue() string {
	return vo.value
}

func (vo *UUIDValueVO) enforceUUIDformat(value string) error {
	r, _ := regexp.Compile("[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}")
	if !r.Match([]byte(value)) {
		return errors.New(fmt.Sprintf("The value it is not a valid UUIDValueVO V4 [%s]", value))
	}
	return nil
}
