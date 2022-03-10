package phoneagent

import (
	"context"
	"go-drop-logistik/business/phoneagent"

	"gorm.io/gorm"
)

type mysqlPhoneAgentRepository struct {
	Conn *gorm.DB
}

func NewMySQLPhoneAgentRepository(conn *gorm.DB) phoneagent.Repository {
	return &mysqlPhoneAgentRepository{
		Conn: conn,
	}
}

func (repository *mysqlPhoneAgentRepository) Store(ctx context.Context, phoneId, agentId int) error {
	phoneAgent := &phoneagent.Domain{
		PhoneID: phoneId,
		AgentID: agentId,
	}
	
	rec := fromDomain(*phoneAgent)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlPhoneAgentRepository) GetByAgentID(ctx context.Context, id int) (phoneagent.Domain, error) {
	phoneAgent := PhoneAgent{}
	result := repository.Conn.Where("agent_id = ?", id).First(&phoneAgent)
	if result.Error != nil {
		return phoneagent.Domain{}, result.Error
	}

	return *phoneAgent.ToDomain(), nil
}

func (repository *mysqlPhoneAgentRepository) GetAllByAgentID(ctx context.Context, id int) ([]phoneagent.Domain, error) {
	allPhoneAgent := []PhoneAgent{}
	result := repository.Conn.Where("agent_id = ?", id).Find(&allPhoneAgent)
	if result.Error != nil {
		return []phoneagent.Domain{}, result.Error
	}

	allPhoneAgentDomain := []phoneagent.Domain{}
	for _, value := range allPhoneAgent {
		allPhoneAgentDomain = append(allPhoneAgentDomain, *value.ToDomain())
	}

	return allPhoneAgentDomain, nil
}
