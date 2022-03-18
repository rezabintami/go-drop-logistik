package response

import (
	"go-drop-logistik/business/trucks"
)

type Trucks struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	LicensePlate string `json:"license_plate"`
}

type TrucksPageResponse struct {
	Trucks *[]Trucks `json:"trucks"`
	Total  int       `json:"total"`
}

func FromDomain(truckDomain *trucks.Domain) (res *Trucks) {
	if truckDomain != nil {
		res = &Trucks{
			ID:           truckDomain.ID,
			Name:         truckDomain.Name,
			Type:         truckDomain.Type,
			LicensePlate: truckDomain.LicensePlate,
		}
	}
	return res
}

func FromListDomain(truckDomain []trucks.Domain, Count int) *TrucksPageResponse {
	allTrucks := []Trucks{}
	for _, value := range truckDomain {
		truck := Trucks{
			ID:           value.ID,
			Name:         value.Name,
			Type:         value.Type,
			LicensePlate: value.LicensePlate,
		}
		allTrucks = append(allTrucks, truck)
	}
	result := TrucksPageResponse{}
	result.Trucks = &allTrucks
	result.Total = Count
	return &result
}
