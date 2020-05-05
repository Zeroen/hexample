package vo

type NameVO struct {
	value string
}

func NewNameVO(name string) *NameVO {
	return &NameVO{
		value: name,
	}
}

func (n *NameVO) GetValue() string {
	return n.value
}