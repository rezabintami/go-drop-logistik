package phones

import (
	"go-drop-logistik/business/phones"
	"time"

	"gorm.io/gorm"
)

type Phones struct {
	ID        int `gorm:"primary_key"`
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func fromDomain(phoneDomain phones.Domain) *Phones {
	return &Phones{
		ID:        phoneDomain.ID,
		Phone:     phoneDomain.Phone,
		CreatedAt: phoneDomain.CreatedAt,
		UpdatedAt: phoneDomain.UpdatedAt,
	}
}

func (rec *Phones) ToDomain() (res *phones.Domain) {
	if rec != nil {
		res = &phones.Domain{
			ID:        rec.ID,
			Phone:     rec.Phone,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
}
