package phoneagent

import (
	"context"
	"go-drop-logistik/business/agents"
	"go-drop-logistik/business/phones"
)

type Domain struct {
	PhoneID int
	Phone   *phones.Domain
	AgentID int
	Agent	*agents.Domain
}

type Usecase interface {
	Store(ctx context.Context, phoneId, agentId int) error
	GetAllByAgentID(ctx context.Context, id int) ([]Domain, error)
	GetByAgentID(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, agentId, phoneId int) error
}

type Repository interface {
	Store(ctx context.Context, phoneId, agentId int) error
	GetAllByAgentID(ctx context.Context, id int) ([]Domain, error)
	GetByAgentID(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, agentId, phoneId int) error
}
