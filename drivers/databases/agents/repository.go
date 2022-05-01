package agents

import (
	"context"
	"go-drop-logistik/modules/agents"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreAgentRepository struct {
	tx *gorm.DB
}

func NewPostgreAgentRepository(tx *gorm.DB) agents.Repository {
	return &postgreAgentRepository{
		tx: tx,
	}
}

func (repository *postgreAgentRepository) GetByID(ctx context.Context, id int) (agents.Domain, error) {
	agentById := Agents{}
	result := repository.tx.Where("agents.id = ?", id).First(&agentById)
	if result.Error != nil {
		log.Println("[error] agents.repository.GetByID : failed to execute get data agent query", result.Error)
		return agents.Domain{}, result.Error
	}

	return *agentById.ToDomain(), nil
}

func (repository *postgreAgentRepository) GetByEmail(ctx context.Context, email string) (agents.ExistingDomain, error) {
	rec := Agents{}

	err := repository.tx.Where("agents.email = ?", email).First(&rec).Error
	if err != nil {
		log.Println("[error] agents.repository.GetByEmail : email ", email, "failed to execute get data agent query", err)
		return agents.ExistingDomain{}, err
	}

	return *rec.ToExistingDomain(), nil
}

func (repository *postgreAgentRepository) Register(ctx context.Context, agentDomain *agents.Domain) error {
	rec := fromDomain(*agentDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] agents.repository.Register : failed to execute register agent query", result.Error)
		return result.Error
	}
	return nil
}

func (repository *postgreAgentRepository) Fetch(ctx context.Context, page, perpage int) ([]agents.Domain, int, error) {
	rec := []Agents{}

	offset := (page - 1) * perpage
	err := repository.tx.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		log.Println("[error] agents.repository.Fetch : failed to execute fetch agents query", err)
		return []agents.Domain{}, 0, err
	}

	var totalData int64
	err = repository.tx.Model(&rec).Count(&totalData).Error
	if err != nil {
		log.Println("[error] agents.repository.Fetch : failed to execute count agents query", err)
		return []agents.Domain{}, 0, err
	}

	var result []agents.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *postgreAgentRepository) Update(ctx context.Context, userDomain *agents.Domain, id int) error {
	agentUpdate := fromDomain(*userDomain)

	result := repository.tx.Where("users.id = ?", id).Updates(&agentUpdate)
	if result.Error != nil {
		log.Println("[error] agents.repository.Update : failed to execute update agent query", result.Error)
		return result.Error
	}

	return nil
}
