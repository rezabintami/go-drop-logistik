package phoneagent

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type PhoneAgentUsecase struct {
	phoneAgentRepository Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewPhoneAgentUsecase(pa Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &PhoneAgentUsecase{
		phoneAgentRepository: pa,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (usecase *PhoneAgentUsecase) Store(ctx context.Context, phoneId, agentId int) error {
	err := usecase.phoneAgentRepository.Store(ctx, phoneId, agentId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *PhoneAgentUsecase) GetAllByAgentID(ctx context.Context, Id int) ([]Domain, error) {
	result, err := usecase.phoneAgentRepository.GetAllByAgentID(ctx, Id)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (usecase *PhoneAgentUsecase) GetByAgentID(ctx context.Context, Id int) (Domain, error) {
	result, err := usecase.phoneAgentRepository.GetByAgentID(ctx, Id)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (usecase *PhoneAgentUsecase) Delete(ctx context.Context, agentId, phoneId int) error {
	err := usecase.phoneAgentRepository.Delete(ctx, agentId, phoneId)
	if err != nil {
		return err
	}

	return nil
}
