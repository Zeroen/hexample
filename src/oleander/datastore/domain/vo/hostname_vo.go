package vo

type HostnameVO struct {
	value string
}

func NewHostnameVO(hostname string) *HostnameVO {
	return &HostnameVO{
		value: hostname,
	}
}

func (h *HostnameVO) GetValue() string {
	return h.value
}
