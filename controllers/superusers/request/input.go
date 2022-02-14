package request

import (
	"go-drop-logistik/business/superusers"
)

type Superusers struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

func (req *Superusers) ToDomain() *superusers.Domain {
	return &superusers.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
}
