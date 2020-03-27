package vo

import "errors"

type AgeVO struct {
	value int
}

func NewAgeVO(age int) (*AgeVO, error) {

	err := AgeVO{}.enforcePositiveAgeValue(age)
	if err != nil {
		return nil, err
	}

	err = AgeVO{}.enforceLess130AgeValue(age)
	if err != nil {
		return nil, err
	}

	return &AgeVO{
		value: age,
	}, nil
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
