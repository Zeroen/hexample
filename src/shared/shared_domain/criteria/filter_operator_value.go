package criteria

type FilterOperatorValue struct {
	value string
}

func NewFilterOperatorEqual() *FilterOperatorValue {
	return &FilterOperatorValue{value: "="}
}

func NewFilterOperatorNotEqual() *FilterOperatorValue {
	return &FilterOperatorValue{value: "!="}
}

func NewFilterOperatorGT() *FilterOperatorValue {
	return &FilterOperatorValue{value: ">"}
}

func NewFilterOperatorLT() *FilterOperatorValue {
	return &FilterOperatorValue{value: "<"}
}

func NewFilterOperatorContains() *FilterOperatorValue {
	return &FilterOperatorValue{value: "LIKE"}
}

func NewFilterOperatorNotContains() *FilterOperatorValue {
	return &FilterOperatorValue{value: "NOT LIKE"}
}

func (f *FilterOperatorValue) GetValue() string {
	return f.value
}

func (f *FilterOperatorValue) IsContain() bool {
	return f.value == "LIKE" || f.value == "NOT LIKE"
}
