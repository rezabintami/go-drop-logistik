package users

import (
	"go-drop-logistik/business/users"
	"time"
)

type Users struct {
	ID        int `gorm:"primary_key" json:"id"`
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) ToDomain() *users.Domain {
	return &users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Email:     rec.Email,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Email:     userDomain.Email,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
