package request

import "go-drop-logistik/modules/trucks"

type Trucks struct {
	Name         string `json:"name" validate:"required" validName:"name"`
	Type         string `json:"type" validate:"required" validName:"type"`
	LicensePlate string `json:"license_plate" validate:"required" validName:"license_plate"`
}

func (req *Trucks) ToDomain() *trucks.Domain {
	return &trucks.Domain{
		Name:         req.Name,
		Type:         req.Type,
		LicensePlate: req.LicensePlate,
	}
}
