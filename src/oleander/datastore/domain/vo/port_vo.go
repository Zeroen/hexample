package vo

type PortVO struct {
	value int
}

func NewPortVO(port int) *PortVO {
	return &PortVO{
		value: port,
	}
}

func (p *PortVO) GetValue() int {
	return p.value
}