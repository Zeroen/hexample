package criteria

import "fmt"

type Criteria struct {
	filters    []*Filter
	order      *Order
	pagination *Pagination
}

func NewCriteria(filters []*Filter, order *Order, pagination *Pagination) *Criteria {
	return &Criteria{
		filters:    filters,
		order:      order,
		pagination: pagination,
	}
}

func (c *Criteria) AddFilter(filter *Filter) *Criteria {
	if c.filters == nil {
		c.filters = make([]*Filter, 0)
	}
	c.filters = append(c.filters, filter)
	return c
}

func (c *Criteria) hasFilters() bool {
	return len(c.filters) > 0
}

func (c *Criteria) hasOrder() bool {
	return !c.order.IsNone()
}

func (c *Criteria) GetFilters() []*Filter {
	return c.filters
}

func (c *Criteria) GetOrder() *Order {
	return c.order
}

func (c *Criteria) GetOffset() int {
	return c.pagination.offset
}

func (c *Criteria) GetLimit() int {
	return c.pagination.limit
}

func (c *Criteria) String() string {
	return fmt.Sprintf("%s~~%s~~%s~~%s", c.filters, c.order, c.pagination)
}

func (c *Criteria) BuildQuery() string {
	query := ""

	sep := " "
	if c.filters != nil {
		for _, filters := range c.filters {
			query = query + sep + filters.BuildQuery()
			sep = " AND "
		}
	}

	if c.order != nil {
		query = query + " " + c.order.BuildQuery()
	}

	if c.pagination != nil {
		query = query + " " + c.pagination.BuildQuery()
	}

	return query
}
