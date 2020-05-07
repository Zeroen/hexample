package vo

type CityVO struct {
	value string
}

func NewCityVO (value string) *CityVO{
	return &CityVO{
		value: value,
	}
}

func (n *CityVO) GetValue() string {
	return n.value
}