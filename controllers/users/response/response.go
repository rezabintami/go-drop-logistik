package response

import (
	"go-drop-logistik/modules/users"
)

type Users struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(userDomain users.Domain) Users {
	return Users{
		ID:    userDomain.ID,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}
