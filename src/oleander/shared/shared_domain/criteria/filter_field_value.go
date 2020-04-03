package criteria

type FilterFieldValue struct {
	value string
}

func NewFilterFieldValue(value string) *FilterFieldValue {
	return &FilterFieldValue{
		value: value,
	}
}

func (o *FilterFieldValue) GetValue() string {
	return o.value
}

func (o *FilterFieldValue) String() string {
	return o.value
}
