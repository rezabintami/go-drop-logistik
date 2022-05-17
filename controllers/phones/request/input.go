package request

import "go-drop-logistik/modules/phones"

type Phone struct {
	Phone string `json:"phone" validate:"required,phone,min=10,max=16" validName:"phoneNumber"`
}

func (req *Phone) ToDomain() *phones.Domain {
	return &phones.Domain{
		Phone: req.Phone,
	}
}
