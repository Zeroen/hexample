package get_id

import "hexample.com/src/oleander/user/domain"

type UserDTO struct {
	ID    string `json:"id,omitempty"`
	Age   int    `json:"age,omitempty"`
	Name  string `json:"name,omitempty"`
	EmailID string `json:"emailID,omitempty"`
}

func NewUserDTO(ag *domain.UserAG) *UserDTO {
	return &UserDTO{
		ID:    ag.GetID(),
		Age:   ag.GetAge(),
		Name:  ag.GetName(),
		EmailID: ag.GetEmail(),
	}
}
