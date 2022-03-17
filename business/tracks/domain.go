package tracks

import (
	"context"
	"go-drop-logistik/business/agents"
	"go-drop-logistik/business/manifest"
	"time"
)

type Domain struct {
	ID                 int `gorm:"primary_key"`
	StartAgentID       int
	StartAgent         *agents.Domain
	CurrentAgentID     int
	CurrentAgent       *agents.Domain
	DestinationAgentID int
	DestinationAgent   *agents.Domain
	ManifestID         int
	Manifest 		   *manifest.Domain
	Message            string
	CreatedAt          time.Time
}

type Usecase interface {
	StoreTrack(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	StoreTrack(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
}
