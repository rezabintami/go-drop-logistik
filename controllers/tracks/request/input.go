package request

import "go-drop-logistik/business/tracks"

type Track struct {
	StartAgentID       int    `json:"start_agent_id"`
	CurrentAgentID     int    `json:"current_agent_id"`
	DestinationAgentID int    `json:"destination_agent_id"`
	Message            string `json:"message"`
}

func (req *Track) ToDomain() *tracks.Domain {
	return &tracks.Domain{
		StartAgentID:       req.StartAgentID,
		CurrentAgentID:     req.CurrentAgentID,
		DestinationAgentID: req.DestinationAgentID,
		Message:            req.Message,
	}
}
