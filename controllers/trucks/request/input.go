package request

import "go-drop-logistik/modules/trucks"

type Trucks struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	LicensePlate string `json:"license_plate"`
}

func (req *Trucks) ToDomain() *trucks.Domain {
	return &trucks.Domain{
		Name:         req.Name,
		Type:         req.Type,
		LicensePlate: req.LicensePlate,
	}
}
