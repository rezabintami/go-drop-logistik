package request

import "go-drop-logistik/business/manifest"

type Manifest struct {
	DriverID int `json:"driver_id"`
}

func (req *Manifest) ToDomain() *manifest.Domain {
	return &manifest.Domain{
		DriverID: req.DriverID,
	}
}

type ManifestUpdate struct {
	Status   string `json:"status"`
	DriverID int    `json:"driver_id"`
}

func (req *ManifestUpdate) ToDomain() *manifest.Domain {
	return &manifest.Domain{
		Status:   req.Status,
		DriverID: req.DriverID,
	}
}
