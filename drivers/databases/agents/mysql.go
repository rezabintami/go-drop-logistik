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
	result := repository.Conn.Where("agents.id = ?", id).First(&agentById)
	if result.Error != nil {
		return agents.Domain{}, result.Error
	}

	return *agentById.ToDomain(), nil
}

func (repository *mysqlAgentRepository) GetByEmail(ctx context.Context, email string) (agents.ExistingDomain, error) {
	rec := Agents{}

	err := repository.Conn.Where("agents.email = ?", email).First(&rec).Error
	if err != nil {
		return agents.ExistingDomain{}, err
	}

	return *rec.ToExistingDomain(), nil
}

func (repository *mysqlAgentRepository) Register(ctx context.Context, agentDomain *agents.Domain) error {
	rec := fromDomain(*agentDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlAgentRepository) Fetch(ctx context.Context, page, perpage int) ([]agents.Domain, int, error) {
	rec := []Agents{}

	offset := (page - 1) * perpage
	err := repository.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []agents.Domain{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []agents.Domain{}, 0, err
	}

	var result []agents.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *mysqlAgentRepository) Update(ctx context.Context, userDomain *agents.Domain, id int) error {
	agentUpdate := fromDomain(*userDomain)

	result := repository.Conn.Where("users.id = ?", id).Updates(&agentUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}