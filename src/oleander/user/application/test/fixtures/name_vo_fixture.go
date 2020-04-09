package fixtures

import "hexample.com/src/oleander/user/domain/vo"

func NewNameVOFixture(name string) *vo.NameVO {
	vo, _ := vo.NewNameVO(name)
	return vo
}

func NewNameVOFixture1() *vo.NameVO {
	vo, _ := vo.NewNameVO("Cristian")
	return vo
}
