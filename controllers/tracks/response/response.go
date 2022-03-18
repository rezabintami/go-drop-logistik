package response

import (
	"go-drop-logistik/business/tracks"
	agentResp "go-drop-logistik/controllers/agents/response"
)

type Track struct {
	StartAgent       *agentResp.TrackAgentResponse `json:"start_agent"`
	CurrentAgent     *agentResp.TrackAgentResponse `json:"current_agent"`
	DestinationAgent *agentResp.TrackAgentResponse `json:"destination_agent"`
	Message          string            `json:"message"`
}

func FromDomain(trackDomain tracks.Domain) Track {
	return Track{
		StartAgent:       agentResp.FromTrackDomain(trackDomain.StartAgent),
		CurrentAgent:     agentResp.FromTrackDomain(trackDomain.CurrentAgent),
		DestinationAgent: agentResp.FromTrackDomain(trackDomain.DestinationAgent),
		Message:          trackDomain.Message,
	}
}

func FromListDomain(trackDomain *[]tracks.Domain) *[]Track {
	allTrack := []Track{}
	for _, value := range *trackDomain {
		manifest := Track{
			StartAgent:       agentResp.FromTrackDomain(value.StartAgent),
			CurrentAgent:     agentResp.FromTrackDomain(value.CurrentAgent),
			DestinationAgent: agentResp.FromTrackDomain(value.DestinationAgent),
			Message:          value.Message,
		}
		allTrack = append(allTrack, manifest)
	}

	return &allTrack
}
