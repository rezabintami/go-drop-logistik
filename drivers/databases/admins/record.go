package admins

import (
	"go-drop-logistik/business/admins"
	"time"
)

type Admins struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Password  string
	Email     string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Admins) ToDomain() *admins.Domain {
	return &admins.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(adminDomain admins.Domain) *Admins {
	return &Admins{
		ID:        adminDomain.ID,
		Name:      adminDomain.Name,
		Password:  adminDomain.Password,
		Email:     adminDomain.Email,
		Roles:     adminDomain.Roles,
		CreatedAt: adminDomain.CreatedAt,
		UpdatedAt: adminDomain.UpdatedAt,
	}
}
