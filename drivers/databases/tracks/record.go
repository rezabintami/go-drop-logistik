package tracks

import (
	"go-drop-logistik/business/tracks"
	"go-drop-logistik/drivers/databases/agents"
	"time"

	"gorm.io/gorm"
)

type Tracks struct {
	ID                 int `gorm:"primary_key"`
	StartAgentID       int
	StartAgent         *agents.Agents `gorm:"foreignkey:StartAgentID;references:ID"`
	CurrentAgentID     int
	CurrentAgent       *agents.Agents `gorm:"foreignkey:CurrentAgentID;references:ID"`
	DestinationAgentID int
	DestinationAgent   *agents.Agents `gorm:"foreignkey:DestinationAgentID;references:ID"`
	ManifestID         int
	Message            string
	CreatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

func (rec *Tracks) ToDomain() (res *tracks.Domain) {
	if rec != nil {
		res = &tracks.Domain{
			ID:                 rec.ID,
			StartAgentID:       rec.StartAgentID,
			StartAgent:         rec.StartAgent.ToDomain(),
			CurrentAgentID:     rec.CurrentAgentID,
			CurrentAgent:       rec.CurrentAgent.ToDomain(),
			DestinationAgentID: rec.DestinationAgentID,
			DestinationAgent:   rec.DestinationAgent.ToDomain(),
			ManifestID:         rec.ManifestID,
			Message:            rec.Message,
			CreatedAt:          rec.CreatedAt,
		}
	}
	return res
}

func fromDomain(trackDomain tracks.Domain) *Tracks {
	return &Tracks{
		ID:                 trackDomain.ID,
		StartAgentID:       trackDomain.StartAgentID,
		CurrentAgentID:     trackDomain.CurrentAgentID,
		DestinationAgentID: trackDomain.DestinationAgentID,
		ManifestID:         trackDomain.ManifestID,
		Message:            trackDomain.Message,
		CreatedAt:          trackDomain.CreatedAt,
	}
}
