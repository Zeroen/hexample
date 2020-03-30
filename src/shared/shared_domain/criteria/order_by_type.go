package criteria

type OrderByType struct {
	value string
}

func NewOrderByTypeDesc() *OrderByType {
	return &OrderByType{value: "desc"}
}

func NewOrderByTypeAsc() *OrderByType {
	return &OrderByType{value: "asc"}
}

func NewOrderByTypeNone() *OrderByType {
	return &OrderByType{value: "none"}
}

func (o *OrderByType) GetValue() string {
	return o.value
}

func (o OrderByType) isNone() bool {
	return o.value == "none"
}
