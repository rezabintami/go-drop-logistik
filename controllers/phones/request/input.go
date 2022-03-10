package request

import "go-drop-logistik/business/phones"

type Phone struct {
	Phone string `json:"phone"`
}

func (req *Phone) ToDomain() *phones.Domain {
	return &phones.Domain{
		Phone: req.Phone,
	}
}
