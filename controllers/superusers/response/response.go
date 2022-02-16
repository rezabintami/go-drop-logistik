package response

import (
	"go-drop-logistik/business/agents"
	"go-drop-logistik/business/superusers"
)

type Superusers struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Agents struct {
	ID        int     `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	Balance   float64 `json:"balance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func FromDomain(superuserDomain superusers.Domain) Superusers {
	return Superusers{
		ID:    superuserDomain.ID,
		Name:  superuserDomain.Name,
		Email: superuserDomain.Email,
	}
}

func AgentFromDomain(agentDomain agents.Domain) Agents {
	return Agents{
		ID:    agentDomain.ID,
		Name:  agentDomain.Name,
		Email: agentDomain.Email,
		Address:   agentDomain.Address,
		Balance:   agentDomain.Balance,
		Latitude:  agentDomain.Latitude,
		Longitude: agentDomain.Longitude,
	}
}
