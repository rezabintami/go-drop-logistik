package request

import (
	"go-drop-logistik/modules/agents"
)

type Agents struct {
	Name     string `json:"name"  validate:"required" validName:"name"`
	Password string `json:"password,omitempty"  validate:"required" validName:"password"`
	Email    string `json:"email"  validate:"required,email,max=100" validName:"email"`
}

func (req *Agents) ToDomain() *agents.Domain {
	return &agents.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}
