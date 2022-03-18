package drivers

import (
	"go-drop-logistik/business/drivers"
	"go-drop-logistik/drivers/databases/trucks"
)

type Drivers struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Phone   string
	Address string
	TruckID int
	Truck   *trucks.Trucks `gorm:"foreignkey:TruckID;references:ID"`
}

func fromDomain(driverDomain drivers.Domain) *Drivers {
	return &Drivers{
		ID:      driverDomain.ID,
		Name:    driverDomain.Name,
		Phone:   driverDomain.Phone,
		Address: driverDomain.Address,
		TruckID: driverDomain.TruckID,
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
	}
}
