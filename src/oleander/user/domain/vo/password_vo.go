package vo

type PasswordVO struct {
	value     string
	temporary bool
}

func NewPasswordVO() (*PasswordVO, error) {

	a := PasswordVO{}
	a.CalculateNewPassword()

	return &a, nil
}

func (vo *PasswordVO) GetValue() string {
	return vo.value
}

func (vo *PasswordVO) CalculateNewPassword() {
	vo.value = "temporary"
	vo.temporary = true
}
