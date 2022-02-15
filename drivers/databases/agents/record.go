package agents

import (
	"go-drop-logistik/business/agents"
	"time"
)

type Agents struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Password  string
	Email     string
	Roles     string
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
		CreatedAt: agentDomain.CreatedAt,
		UpdatedAt: agentDomain.UpdatedAt,
	}
}
