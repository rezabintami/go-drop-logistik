package phones

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/phoneagent"
	"time"
)

type PhoneUsecase struct {
	phoneRepository      Repository
	phoneAgentRepository phoneagent.Repository
	contextTimeout       time.Duration
	jwtAuth              *middleware.ConfigJWT
}

func NewPhoneUsecase(ur Repository, pa phoneagent.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &PhoneUsecase{
		phoneRepository:      ur,
		phoneAgentRepository: pa,
		jwtAuth:              jwtauth,
		contextTimeout:       timeout,
	}
}

func (usecase *PhoneUsecase) StorePhone(ctx context.Context, phoneDomain *Domain, id int) error {
	phoneId, err := usecase.phoneRepository.StorePhone(ctx, phoneDomain)
	if err != nil {
		return err
	}

	usecase.phoneAgentRepository.Store(ctx, phoneId, id)

	return nil
}

func (usecase *PhoneUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	phone, err := usecase.phoneRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return phone, nil
}

func (usecase *PhoneUsecase) GetAll(ctx context.Context, agentId int) ([]Domain, error) {
	agentPhone, err := usecase.phoneAgentRepository.GetAllByAgentID(ctx, agentId)

	if err != nil {
		return []Domain{}, err
	}

	var phoneDomain []Domain
	for _, phoneId := range agentPhone {
		phone, _ := usecase.phoneRepository.GetByID(ctx, phoneId.PhoneID)
		phoneDomain = append(phoneDomain, phone)
	}

	return phoneDomain, nil
}

func (usecase *PhoneUsecase) Delete(ctx context.Context, agentId, phoneId int) error {
	err := usecase.phoneAgentRepository.Delete(ctx, agentId, phoneId)
	if err != nil {
		return err
	}

	err = usecase.phoneRepository.Delete(ctx, phoneId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *PhoneUsecase) Update(ctx context.Context, phoneDomain *Domain, id int) error {
	err := usecase.phoneRepository.Update(ctx, phoneDomain, id)
	if err != nil {
		return err
	}

	return nil
}
