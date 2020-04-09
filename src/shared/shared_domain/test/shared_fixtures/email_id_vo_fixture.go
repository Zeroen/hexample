package fixtures

import (
	"hexample.com/src/shared/shared_domain"
)

func NewEmailIDVOFixture(name string) *shared_domain.EmailIDValueVO {
	vo, _ := shared_domain.NewEmailIDValueVO(name)
	return vo
}

func NewEmailIDVOFixture1() *shared_domain.EmailIDValueVO {
	vo, _ := shared_domain.NewEmailIDValueVO("b189a87d-b60a-4c02-bd22-f234193026a2")
	return vo
}
