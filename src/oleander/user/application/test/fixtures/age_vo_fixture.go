package fixtures

import "hexample.com/src/oleander/user/domain/vo"

func NewAgeVOFixture(age int) *vo.AgeVO {
	vo, _ := vo.NewAgeVO(age)
	return vo
}

func NewAgeVOFixture1() *vo.AgeVO {
	vo, _ := vo.NewAgeVO(32)
	return vo
}
