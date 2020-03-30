package criteria

import "fmt"

type Filter struct {
	field    *FilterFieldValue
	operator *FilterOperatorValue
	value    *FilterValue
}

func NewFilter(field *FilterFieldValue, operator *FilterOperatorValue, value *FilterValue) (*Filter, error) {
	return &Filter{
		field:    field,
		operator: operator,
		value:    value,
	}, nil
}

func (f *Filter) GetField() *FilterFieldValue {
	return f.field
}

func (f *Filter) GetOperator() *FilterOperatorValue {
	return f.operator
}

func (f *Filter) GetValue() *FilterValue {
	return f.value
}

func (f *Filter) String() string {
	return fmt.Sprintf("%s.%s.%s", f.field.GetValue(), f.operator.GetValue(), f.value.GetValue())
}

func (f *Filter) BuildQuery() string {
	if f.operator.IsContain() {
		return f.field.GetValue() + " " + f.operator.GetValue() + " '%" + f.value.BuildQuery() + "%'"
	} else {
		if f.value.IsNumber() {
			return f.field.GetValue() + " " + f.operator.GetValue() + " " + f.value.BuildQuery()
		} else {
			return f.field.GetValue() + " " + f.operator.GetValue() + " '" + f.value.BuildQuery() + "' "
		}
	}
}
