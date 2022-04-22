package manifest

import (
	"go-drop-logistik/business/manifest"
	"go-drop-logistik/drivers/databases/drivers"
	"time"
)

type Manifest struct {
	ID        int `gorm:"primary_key"`
	Code      string
	Status    string
	DriverID  int
	Driver    *drivers.Drivers `gorm:"foreignkey:DriverID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"column:deletedAt"`
}

func (rec *Manifest) ToDomain() (res *manifest.Domain) {
	if rec != nil {
		res = &manifest.Domain{
			ID:        rec.ID,
			Code:      rec.Code,
			Status:    rec.Status,
			DriverID:  rec.DriverID,
			Driver:    rec.Driver.ToDomain(),
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return res
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
