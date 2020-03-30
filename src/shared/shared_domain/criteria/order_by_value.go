package criteria

type OrderByValue struct {
	value string
}

func NewOrderByValue(value string) *OrderByValue {
	return &OrderByValue{
		value: value,
	}
}

func (o *OrderByValue) GetValue() string {
	return o.value
}

func (o *OrderByValue) String() string {
	return o.value
}
