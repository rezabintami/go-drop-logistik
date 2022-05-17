package phoneagent

import (
	"context"
	"go-drop-logistik/modules/phoneagent"
	"log"

	"github.com/jinzhu/gorm"
)

type postgrePhoneAgentRepository struct {
	tx *gorm.DB
}

func NewPostgrePhoneAgentRepository(tx *gorm.DB) phoneagent.Repository {
	return &postgrePhoneAgentRepository{
		tx: tx,
	}
}

func (repository *postgrePhoneAgentRepository) Store(ctx context.Context, phoneId, agentId int) error {
	phoneAgent := &phoneagent.Domain{
		PhoneID: phoneId,
		AgentID: agentId,
	}

	rec := fromDomain(*phoneAgent)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] phoneagents.repository.Store : failed to execute store phoneagent query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgrePhoneAgentRepository) GetByAgentID(ctx context.Context, id int) (phoneagent.Domain, error) {
	phoneAgent := PhoneAgent{}
	result := repository.tx.Preload("Phone").Where("agent_id = ?", id).First(&phoneAgent)
	if result.Error != nil {
		log.Println("[error] phoneagents.repository.GetByAgentID : failed to execute get data phoneagent query", result.Error)
		return phoneagent.Domain{}, result.Error
	}

	return *phoneAgent.ToDomain(), nil
}

func (repository *postgrePhoneAgentRepository) GetAllByAgentID(ctx context.Context, id int) ([]phoneagent.Domain, error) {
	allPhoneAgent := []PhoneAgent{}
	result := repository.tx.Preload("Phone").Where("agent_id = ?", id).Find(&allPhoneAgent)
	if result.Error != nil {
		log.Println("[error] phoneagents.repository.GetAllByAgentID : failed to execute get data phoneagents query", result.Error)
		return []phoneagent.Domain{}, result.Error
	}

	allPhoneAgentDomain := []phoneagent.Domain{}
	for _, value := range allPhoneAgent {
		allPhoneAgentDomain = append(allPhoneAgentDomain, *value.ToDomain())
	}

	return allPhoneAgentDomain, nil
}

func (repository *postgrePhoneAgentRepository) Delete(ctx context.Context, agentId, phoneId int) error {
	phoneDelete := PhoneAgent{}
	result := repository.tx.Preload("Phone").Preload("Agent").Where("agent_id = ?", agentId).Where("phone_id = ?", phoneId).Delete(&phoneDelete)
	if result.Error != nil {
		log.Println("[error] phoneagents.repository.Delete : failed to execute delete phoneagent query", result.Error)
		return result.Error
	}

	return nil
}
