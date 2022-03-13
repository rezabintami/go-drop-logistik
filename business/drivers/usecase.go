package drivers

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type DriversUsecase struct {
	driversRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewDriverUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &DriversUsecase{
		driversRepository: ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
	}
}

func (usecase *DriversUsecase) Store(ctx context.Context, phoneDomain *Domain) error {
	err := usecase.driversRepository.Store(ctx, phoneDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *DriversUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	phone, err := usecase.driversRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return phone, nil
}

func (usecase *DriversUsecase) Delete(ctx context.Context, id int) error {
	err := usecase.driversRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *DriversUsecase) Update(ctx context.Context, phoneDomain *Domain, id int) error {
	err := usecase.driversRepository.Update(ctx, phoneDomain, id)
	if err != nil {
		return err
	}

	return nil
}
