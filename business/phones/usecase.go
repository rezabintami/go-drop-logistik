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
