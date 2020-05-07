package vo

type CountryVO struct {
	value string
}

func NewCountryVO (value string) *CountryVO{
	return &CountryVO{
		value: value,
	}
}

func (n *CountryVO) GetValue() string {
	return n.value
}