package request

import (
	"go-drop-logistik/business/agents"
)

type Agents struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

func (req *Agents) ToDomain() *agents.Domain {
	return &agents.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}
