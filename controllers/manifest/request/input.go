package request

import "go-drop-logistik/business/manifest"

type Manifest struct {
	DriverID int    `json:"driver_id"`
}

func (req *Manifest) ToDomain() *manifest.Domain {
	return &manifest.Domain{
		DriverID: req.DriverID,
	}
}
