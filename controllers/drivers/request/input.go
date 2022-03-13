package request

import "go-drop-logistik/business/drivers"

type Drivers struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	TruckID int    `json:"truck_id"`
}

func (req *Drivers) ToDomain() *drivers.Domain {
	return &drivers.Domain{
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
		TruckID: req.TruckID,
	}
}
