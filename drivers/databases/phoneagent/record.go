package phoneagent

import (
	"go-drop-logistik/business/phoneagent"
)

type PhoneAgent struct {
	PhoneID int
	AgentID int
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
		AgentID: rec.AgentID,
	}
}
