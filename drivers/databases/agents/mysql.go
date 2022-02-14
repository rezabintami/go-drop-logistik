package agents

import (
	"context"
	"go-drop-logistik/business/agents"

	"gorm.io/gorm"
)

type mysqlAgentRepository struct {
	Conn *gorm.DB
}

func NewMySQLAgentRepository(conn *gorm.DB) agents.Repository {
	return &mysqlAgentRepository{
		Conn: conn,
	}
}

func (repository *mysqlAgentRepository) GetByID(ctx context.Context, id int) (agents.Domain, error) {
	agentById := Agents{}
	result := repository.Conn.Where("agent.id = ?", id).First(&agentById)
	if result.Error != nil {
		return agents.Domain{}, result.Error
	}

	return *agentById.ToDomain(), nil
}

func (repository *mysqlAgentRepository) GetByEmail(ctx context.Context, email string) (agents.Domain, error) {
	rec := Agents{}

	err := repository.Conn.Where("agent.email = ?", email).First(&rec).Error
	if err != nil {
		return agents.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *mysqlAgentRepository) Register(ctx context.Context, agentDomain *agents.Domain) error {
	rec := fromDomain(*agentDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}