package vo

type PathVO struct {
	value string
}

func NewPathVO(path string) *PathVO {
	return &PathVO{
		value: path,
	}
}

func (p *PathVO) GetValue() string {
	return p.value
}