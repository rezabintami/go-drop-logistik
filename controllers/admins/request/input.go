package request

import (
	"go-drop-logistik/business/admins"
	"go-drop-logistik/business/agents"
)

type Admins struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

type Agents struct {
	Name      string  `json:"name"`
	Password  string  `json:"password,omitempty"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	Balance   float64 `json:"balance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

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
