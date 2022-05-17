package phoneagent

import (
	"go-drop-logistik/drivers/databases/agents"
	"go-drop-logistik/drivers/databases/phones"
	"go-drop-logistik/modules/phoneagent"
)

type PhoneAgent struct {
	ID      int `gorm:"primary_key"`
	PhoneID int
	Phone   *phones.Phones `gorm:"foreignkey:PhoneID;references:ID"`
	AgentID int
	Agent   *agents.Agents `gorm:"foreignkey:AgentID;references:ID"`
}

func fromDomain(phoneAgentDomain phoneagent.Domain) *PhoneAgent {
	return &PhoneAgent{
		ID:      phoneAgentDomain.ID,
		PhoneID: phoneAgentDomain.PhoneID,
		AgentID: phoneAgentDomain.AgentID,
	}
}

func (rec *PhoneAgent) ToDomain() *phoneagent.Domain {
	return &phoneagent.Domain{
		ID:      rec.ID,
		PhoneID: rec.PhoneID,
		Phone:   rec.Phone.ToDomain(),
		AgentID: rec.AgentID,
		Agent:   rec.Agent.ToDomain(),
	}
}
