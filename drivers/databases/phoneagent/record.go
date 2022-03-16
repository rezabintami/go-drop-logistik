package phoneagent

import (
	"go-drop-logistik/business/phoneagent"
	"go-drop-logistik/drivers/databases/agents"
	"go-drop-logistik/drivers/databases/phones"
)

type PhoneAgent struct {
	PhoneID int
	Phone   *phones.Phones `gorm:"foreignkey:PhoneID;references:ID"`
	AgentID int
	Agent   *agents.Agents `gorm:"foreignkey:AgentID;references:ID"`
}

func fromDomain(phoneAgentDomain phoneagent.Domain) *PhoneAgent {
	return &PhoneAgent{
		PhoneID: phoneAgentDomain.PhoneID,
		AgentID: phoneAgentDomain.AgentID,
	}
}

func (rec *PhoneAgent) ToDomain() *phoneagent.Domain {
	return &phoneagent.Domain{
		PhoneID: rec.PhoneID,
		Phone:   rec.Phone.ToDomain(),
		AgentID: rec.AgentID,
		Agent:   rec.Agent.ToDomain(),
	}
}
