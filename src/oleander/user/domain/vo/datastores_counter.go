package vo

type DatastoresCounterVo struct {
	value int
}

func NewDatastoresCounterVo() *DatastoresCounterVo {

	a := DatastoresCounterVo{
		value: 0,
	}

	return &a
}


func (vo *DatastoresCounterVo) Increase() {
	vo.value += 1
}

func (vo *DatastoresCounterVo) GetValue() int {
	return vo.value
}
