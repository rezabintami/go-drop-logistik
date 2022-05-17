package response

import (
	"go-drop-logistik/modules/admins"
	"go-drop-logistik/modules/agents"
)

type Admins struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Agents struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	Balance   float64 `json:"balance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AgentsPageResponse struct {
	Users *[]Agents `json:"agents"`
	Total int       `json:"total"`
}

func TokenFromDomain(accessToken, refreshToken string) Token {
	return Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

func FromDomain(adminDomain admins.Domain) Admins {
	return Admins{
		ID:    adminDomain.ID,
		Name:  adminDomain.Name,
		Email: adminDomain.Email,
	}
}

func AgentFromDomain(agentDomain agents.Domain) Agents {
	return Agents{
		ID:        agentDomain.ID,
		Name:      agentDomain.Name,
		Email:     agentDomain.Email,
		Address:   agentDomain.Address,
		Balance:   agentDomain.Balance,
		Latitude:  agentDomain.Latitude,
		Longitude: agentDomain.Longitude,
	}
}

func AgentFromListDomain(agentDomain []agents.Domain, Count int) *AgentsPageResponse {
	allAgent := []Agents{}
	for _, value := range agentDomain {
		agent := Agents{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Address:   value.Address,
			Balance:   value.Balance,
			Latitude:  value.Latitude,
			Longitude: value.Longitude,
		}
		allAgent = append(allAgent, agent)
	}
	result := AgentsPageResponse{}
	result.Users = &allAgent
	result.Total = Count
	return &result
}
