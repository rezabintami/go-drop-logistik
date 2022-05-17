package response

import (
	"go-drop-logistik/modules/phones"
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

func FromListDomain(agentDomain []phones.Domain) *[]Phones {
	allPhone := []Phones{}
	for _, value := range agentDomain {
		phone := Phones{
			ID:        value.ID,
			Phone:     value.Phone,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		allPhone = append(allPhone, phone)
	}

	return &allPhone
}