package fixtures

import (
	"hexample.com/src/shared/shared_domain"
)

func NewUserIDVOFixture(name string) *shared_domain.UserIDValueVO {
	vo, _ := shared_domain.NewUserIDValueVO(name)
	return vo
}

func NewUserIDVOFixture1() *shared_domain.UserIDValueVO {
	vo, _ := shared_domain.NewUserIDValueVO("2d481fcb-c0cf-4ae4-bb43-745fa6f19ea4")
	return vo
}
