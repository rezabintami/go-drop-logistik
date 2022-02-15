package superusers

import (
	"go-drop-logistik/business/superusers"
	"time"
)

type Superusers struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Password  string
	Email     string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Superusers) ToDomain() *superusers.Domain {
	return &superusers.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(superuserDomain superusers.Domain) *Superusers {
	return &Superusers{
		ID:        superuserDomain.ID,
		Name:      superuserDomain.Name,
		Password:  superuserDomain.Password,
		Email:     superuserDomain.Email,
		Roles:     superuserDomain.Roles,
		CreatedAt: superuserDomain.CreatedAt,
		UpdatedAt: superuserDomain.UpdatedAt,
	}
}
