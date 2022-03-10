package response

import (
	"go-drop-logistik/business/agents"
)

type Agents struct {
	ID        int      `gorm:"primary_key" json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Phone     []string `json:"phone"`
	Address   string   `json:"address"`
	Balance   float64  `json:"balance"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

func FromDomain(userDomain agents.Domain) Agents {
	return Agents{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Phone:     userDomain.Phone,
		Address:   userDomain.Address,
		Balance:   userDomain.Balance,
		Latitude:  userDomain.Latitude,
		Longitude: userDomain.Longitude,
	}
}
