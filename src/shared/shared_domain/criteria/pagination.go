package criteria

import (
	"errors"
	"fmt"
)

type Pagination struct {
	offset int
	limit  int
}

func NewPagination(offset int, limit int) (*Pagination, error) {
	var pag Pagination

	err := pag.enforceNonNegativeOffset(offset)
	if err != nil {
		return nil, err
	}
	pag.offset = offset

	err = pag.enforceNonNegativeLimit(limit)
	if err != nil {
		return nil, err
	}
	pag.limit = limit

	return &pag, nil
}

func (*Pagination) enforceNonNegativeOffset(offset int) error {
	if offset < 0 {
		return errors.New(fmt.Sprintf("Invalid value offset can not be negative: %s", offset))
	}
	return nil
}

func (*Pagination) enforceNonNegativeLimit(limit int) error {
	if limit < 0 {
		return errors.New(fmt.Sprintf("Invalid value limit can not be negative: %s", limit))
	}
	return nil
}

func (p *Pagination) BuildQuery() string {
	return fmt.Sprintf("LIMIT %v,%v", p.offset, p.limit)
}
