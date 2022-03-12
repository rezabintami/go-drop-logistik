package manifest

import (
	"go-drop-logistik/business/manifest"
	"time"

	"gorm.io/gorm"
)

type Manifest struct {
	ID        int `gorm:"primary_key"`
	Code      string
	Status    string
	DriverID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Manifest) ToDomain() *manifest.Domain {
	return &manifest.Domain{
		ID:        rec.ID,
		Code:      rec.Code,
		Status:    rec.Status,
		DriverID:  rec.DriverID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(manifestDomain manifest.Domain) *Manifest {
	return &Manifest{
		ID:        manifestDomain.ID,
		Code:      manifestDomain.Code,
		Status:    manifestDomain.Status,
		DriverID:  manifestDomain.DriverID,
		CreatedAt: manifestDomain.CreatedAt,
		UpdatedAt: manifestDomain.UpdatedAt,
	}
}
