package criteria

import "fmt"

type Order struct {
	orderBy   *OrderByValue
	orderType *OrderByType
}

func NewOrderDesc(orderBy *OrderByValue) *Order {
	return NewOrder(orderBy, NewOrderByTypeDesc())
}

func NewOrderAsc(orderBy *OrderByValue) *Order {
	return NewOrder(orderBy, NewOrderByTypeAsc())
}

func NewOrderNone() *Order {
	return NewOrder(NewOrderByValue(""), NewOrderByTypeNone())
}

func NewOrder(orderBy *OrderByValue, orderType *OrderByType) *Order {
	o := Order{}
	o.orderType = orderType
	o.orderBy = orderBy
	return &o
}

func (o *Order) GetOrderBy() *OrderByValue {
	return o.GetOrderBy()
}

func (o *Order) GetOrderTypeBy() *OrderByType {
	return o.GetOrderTypeBy()
}

func (o *Order) IsNone() bool {
	return o.orderType.isNone()
}

func (o *Order) String() string {
	return fmt.Sprintf("%s.%s", o.orderBy.GetValue(), o.orderType.GetValue())
}

func (o *Order) BuildQuery() string {
	query := "ORDER BY " + o.orderBy.GetValue()
	if o.IsNone() {
		return ""
	}

	return query + " " + o.orderType.GetValue()
}
