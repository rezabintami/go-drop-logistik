package tracks

import (
	"context"
	"go-drop-logistik/modules/agents"
	"time"
)

type Domain struct {
	ID                 int
	StartAgentID       int
	StartAgent         *agents.Domain
	CurrentAgentID     int
	CurrentAgent       *agents.Domain
	DestinationAgentID int
	DestinationAgent   *agents.Domain
	Message            string
	CreatedAt          time.Time
}

type Usecase interface {
	StoreTrack(ctx context.Context, data *Domain, agentName string) (int, error)
	Delete(ctx context.Context, trackId, agentId int) error
	Update(ctx context.Context, trackId, agentId int, data *Domain) error
}

type Repository interface {
	StoreTrack(ctx context.Context, data *Domain) (int, error)
	Delete(ctx context.Context, trackId, agentId int) error
	Update(ctx context.Context, trackId, agentId int, data *Domain) error
}
