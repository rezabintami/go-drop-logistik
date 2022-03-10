package phoneagent

import (
	"context"
)

type Domain struct {
	PhoneID int
	AgentID int
}

type Repository interface {
	Store(ctx context.Context, phoneId, agentId int) error
	GetAllByAgentID(ctx context.Context, id int) ([]Domain, error)
	GetByAgentID(ctx context.Context, id int) (Domain, error)
}
