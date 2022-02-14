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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Superusers) ToDomain() *superusers.Domain {
	return &superusers.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain superusers.Domain) *Superusers {
	return &Superusers{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Email:     userDomain.Email,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
