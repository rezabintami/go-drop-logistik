package response

import (
	"go-drop-logistik/business/tracks"
	agentResp "go-drop-logistik/controllers/agents/response"
	manifestResp "go-drop-logistik/controllers/manifest/response"
)

type Track struct {
	StartAgent       *agentResp.Agents      `json:"start_agent"`
	CurrentAgent     *agentResp.Agents      `json:"current_agent"`
	DestinationAgent *agentResp.Agents      `json:"destination_agent"`
	Manifest         *manifestResp.Manifest `json:"manifest"`
	Message          string                 `json:"message"`
}

func FromDomain(trackDomain tracks.Domain) Track {
	return Track{
		StartAgent:       agentResp.FromDomain(trackDomain.StartAgent),
		CurrentAgent:     agentResp.FromDomain(trackDomain.CurrentAgent),
		DestinationAgent: agentResp.FromDomain(trackDomain.DestinationAgent),
		Manifest:         manifestResp.FromDomain(trackDomain.Manifest),
		Message:          trackDomain.Message,
	}
}
