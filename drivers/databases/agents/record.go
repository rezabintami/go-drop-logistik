package agents

import (
	"go-drop-logistik/business/agents"
	"time"
)

type Agents struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Password  string
	Email     string
	Roles     string
	Address   string
	Balance   float64
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Agents) ToDomain() *agents.Domain {
	return &agents.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		Roles:     rec.Roles,
		Address:   rec.Address,
		Balance:   rec.Balance,
		Latitude:  rec.Latitude,
		Longitude: rec.Longitude,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(agentDomain agents.Domain) *Agents {
	return &Agents{
		ID:        agentDomain.ID,
		Name:      agentDomain.Name,
		Password:  agentDomain.Password,
		Email:     agentDomain.Email,
		Roles:     agentDomain.Roles,
		Address:   agentDomain.Address,
		Balance:   agentDomain.Balance,
		Latitude:  agentDomain.Latitude,
		Longitude: agentDomain.Longitude,
		CreatedAt: agentDomain.CreatedAt,
		UpdatedAt: agentDomain.UpdatedAt,
	}
}
