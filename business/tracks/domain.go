package tracks

import (
	"context"
	"go-drop-logistik/business/agents"
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
}

type Repository interface {
	StoreTrack(ctx context.Context, data *Domain) (int, error)
}
