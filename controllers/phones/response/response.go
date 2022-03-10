package response

import (
	"go-drop-logistik/business/phones"
	"time"
)

type Phones struct {
	ID        int       `json:"id"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(phoneDomain phones.Domain) Phones {
	return Phones{
		ID:        phoneDomain.ID,
		Phone:     phoneDomain.Phone,
		CreatedAt: phoneDomain.CreatedAt,
		UpdatedAt: phoneDomain.UpdatedAt,
	}
}
