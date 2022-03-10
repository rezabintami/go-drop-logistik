package phones

import (
	"go-drop-logistik/business/phones"
	"time"
)

type Phones struct {
	ID        int `gorm:"primary_key"`
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(phoneDomain phones.Domain) *Phones {
	return &Phones{
		Phone:     phoneDomain.Phone,
		CreatedAt: phoneDomain.CreatedAt,
	}
}

func (rec *Phones) ToDomain() *phones.Domain {
	return &phones.Domain{
		ID:        rec.ID,
		Phone:     rec.Phone,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
