package criteria

const ValueTypeString = "string"
const ValueTypeNumber = "number"

type FilterValue struct {
	valueType string
	value     string
}

func NewFilterStringValue(value string) *FilterValue {
	return &FilterValue{
		value:     value,
		valueType: ValueTypeString,
	}
}

func NewFilterIntValue(value string) *FilterValue {
	return &FilterValue{
		value:     value,
		valueType: ValueTypeNumber,
	}
}

func (o *FilterValue) GetValue() string {
	return o.value
}

func (o *FilterValue) String() string {
	return o.value
}

func (o *FilterValue) IsNumber() bool {
	return o.valueType == ValueTypeNumber
}

func (o *FilterValue) BuildQuery() string {
	return o.value
}
