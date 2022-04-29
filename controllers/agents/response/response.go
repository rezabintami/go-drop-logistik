package response

import (
	"go-drop-logistik/modules/agents"
)

type Agents struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Phone     []string `json:"phone"`
	Address   string   `json:"address"`
	Balance   float64  `json:"balance"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

type TrackAgentResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

func FromDomain(userDomain *agents.Domain) (res *Agents) {
	if userDomain != nil {
		res = &Agents{
			ID:        userDomain.ID,
			Name:      userDomain.Name,
			Email:     userDomain.Email,
			Phone:     userDomain.Phone,
			Address:   userDomain.Address,
			Balance:   userDomain.Balance,
			Latitude:  userDomain.Latitude,
			Longitude: userDomain.Longitude,
		}
	}
	return res
}

func FromTrackDomain(userDomain *agents.Domain) (res *TrackAgentResponse) {
	if userDomain != nil {
		res = &TrackAgentResponse{
			ID:        userDomain.ID,
			Name:      userDomain.Name,
			Address:   userDomain.Address,
			Latitude:  userDomain.Latitude,
			Longitude: userDomain.Longitude,
		}
	}
	return res
}