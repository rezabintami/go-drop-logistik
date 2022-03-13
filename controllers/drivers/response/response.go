package response

import (
	"go-drop-logistik/business/drivers"
	truckResp "go-drop-logistik/controllers/trucks/response"
)

type Drivers struct {
	Name    string            `json:"name"`
	Phone   string            `json:"phone"`
	Address string            `json:"address"`
	Truck   *truckResp.Trucks `json:"truck"`
}

func FromDomain(driverDomain drivers.Domain) Drivers {
	return Drivers{
		Name:    driverDomain.Name,
		Phone:   driverDomain.Phone,
		Address: driverDomain.Address,
		Truck:   truckResp.FromDomain(driverDomain.Truck),
	}
}
