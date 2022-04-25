package trucks

import (
	"go-drop-logistik/business/trucks"
	"time"
)

type Trucks struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Type         string
	LicensePlate string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func fromDomain(truckDomain trucks.Domain) *Trucks {
	return &Trucks{
		ID:           truckDomain.ID,
		Name:         truckDomain.Name,
		Type:         truckDomain.Type,
		LicensePlate: truckDomain.LicensePlate,
		CreatedAt:    truckDomain.CreatedAt,
		UpdatedAt:    truckDomain.UpdatedAt,
	}
}

func (rec *Trucks) ToDomain() (res *trucks.Domain) {
	if rec != nil {
		res = &trucks.Domain{
			ID:           rec.ID,
			Name:         rec.Name,
			Type:         rec.Type,
			LicensePlate: rec.LicensePlate,
			CreatedAt:    rec.CreatedAt,
			UpdatedAt:    rec.UpdatedAt,
		}
	}
	return res
}
