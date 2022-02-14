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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Agents) ToDomain() *agents.Domain {
	return &agents.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain agents.Domain) *Agents {
	return &Agents{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Email:     userDomain.Email,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
