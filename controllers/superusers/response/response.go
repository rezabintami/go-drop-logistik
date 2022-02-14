package response

import (
	"go-drop-logistik/business/superusers"
)

type Superusers struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(userDomain superusers.Domain) Superusers {
	return Superusers{
		ID:    userDomain.ID,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}
