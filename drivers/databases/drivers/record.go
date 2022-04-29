package drivers

import (
	"go-drop-logistik/drivers/databases/trucks"
	"go-drop-logistik/modules/drivers"
	"time"
)

type Drivers struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Phone   string
	Address string
	TruckID int
	Truck   *trucks.Trucks `gorm:"foreignkey:TruckID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(driverDomain drivers.Domain) *Drivers {
	return &Drivers{
		ID:      driverDomain.ID,
		Name:    driverDomain.Name,
		Phone:   driverDomain.Phone,
		Address: driverDomain.Address,
		TruckID: driverDomain.TruckID,
		CreatedAt: driverDomain.CreatedAt,
		UpdatedAt: driverDomain.UpdatedAt,
	}
}

func (rec *Drivers) ToDomain() *drivers.Domain {
	return &drivers.Domain{
		ID:      rec.ID,
		Name:    rec.Name,
		Phone:   rec.Phone,
		Address: rec.Address,
		TruckID: rec.TruckID,
		Truck:   rec.Truck.ToDomain(),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
