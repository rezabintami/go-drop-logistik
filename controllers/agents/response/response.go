package response

import (
	"go-drop-logistik/business/agents"
)

type Agents struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(userDomain agents.Domain) Agents {
	return Agents{
		ID:    userDomain.ID,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}
