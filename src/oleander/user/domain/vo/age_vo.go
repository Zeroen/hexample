package vo

import "errors"

type AgeVO struct {
	value int
}

func NewAgeVO(age int) (*AgeVO, error) {

	a := AgeVO{
		value: age,
	}

	err := a.enforcePositiveAgeValue(age)
	if err != nil {
		return nil, err
	}

	err = a.enforceLess130AgeValue(age)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (vo *AgeVO) enforcePositiveAgeValue(value int) error {
	if value <= 0 {
		return errors.New("User age cannot be a negative number.")
	}
	return nil
}

func (vo *AgeVO) enforceLess130AgeValue(value int) error {
	if value <= 0 {
		return errors.New("User age cannot be greater than 130 years old.")
	}
	return nil
}

func (vo *AgeVO) GetValue() int {
	return vo.value
}
