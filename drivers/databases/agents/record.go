package agents

import (
	"go-drop-logistik/modules/agents"
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
	DeletedAt *time.Time
}

func (rec *Agents) ToExistingDomain() *agents.ExistingDomain {
	return &agents.ExistingDomain{
		ID:        rec.ID,
		Password:  rec.Password,
		Name:      rec.Name,
		Email:     rec.Email,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func (rec *Agents) ToDomain() (res *agents.Domain) {
	if rec != nil {
		res = &agents.Domain{
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
	return res
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
