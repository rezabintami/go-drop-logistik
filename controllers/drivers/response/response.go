package response

import (
	truckResp "go-drop-logistik/controllers/trucks/response"
	"go-drop-logistik/modules/drivers"
)

type Drivers struct {
	Name    string            `json:"name"`
	Phone   string            `json:"phone"`
	Address string            `json:"address"`
	Truck   *truckResp.Trucks `json:"truck"`
}

func FromDomain(driverDomain *drivers.Domain) (res *Drivers) {
	if driverDomain != nil {
		res = &Drivers{
			Name:    driverDomain.Name,
			Phone:   driverDomain.Phone,
			Address: driverDomain.Address,
			Truck:   truckResp.FromDomain(driverDomain.Truck),
		}
	}
	return res
}
