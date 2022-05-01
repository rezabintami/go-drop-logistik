package request

import "go-drop-logistik/modules/drivers"

type Drivers struct {
	Name    string `json:"name"  validate:"required" validName:"name"`
	Phone   string `json:"phone" validate:"required,phone,min=10,max=16" validName:"phoneNumber"`
	Address string `json:"address" validate:"required" validName:"address"`
	TruckID int    `json:"truck_id" validate:"required" validName:"truck_id"`
}

func (req *Drivers) ToDomain() *drivers.Domain {
	return &drivers.Domain{
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
		TruckID: req.TruckID,
	}
}
