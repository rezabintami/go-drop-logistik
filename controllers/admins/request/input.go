package request

import (
	"go-drop-logistik/modules/admins"
	"go-drop-logistik/modules/agents"
)

type (
	Admins struct {
		Name     string `json:"name" validate:"required" validName:"name"`
		Password string `json:"password,omitempty"  validate:"required" validName:"password"`
		Email    string `json:"email" validate:"required,email,max=100" validName:"email"`
	}
	Agents struct {
		Name      string  `json:"name" validate:"required,max=100" validName:"name"`
		Password  string  `json:"password,omitempty" validate:"required" validName:"password"`
		Email     string  `json:"email" validate:"omitempty,email,max=100" validName:"email"`
		Address   string  `json:"address" validate:"required" validName:"address"`
		Balance   float64 `json:"balance"`
		Latitude  float64 `json:"latitude" validate:"required" validName:"latitude"`
		Longitude float64 `json:"longitude" validate:"required" validName:"longitude"`
	}
)

func (req *Admins) ToDomain() *admins.Domain {
	return &admins.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}

func (req *Agents) AgentToDomain() *agents.Domain {
	return &agents.Domain{
		Name:      req.Name,
		Password:  req.Password,
		Email:     req.Email,
		Address:   req.Address,
		Balance:   req.Balance,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
