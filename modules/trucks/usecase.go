package trucks

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type TrucksUsecase struct {
	truckRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewTrucksUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &TrucksUsecase{
		truckRepository: ur,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
	}
}

func (usecase *TrucksUsecase) StoreTruck(ctx context.Context, truckDomain *Domain) error {
	err := usecase.truckRepository.StoreTruck(ctx, truckDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrucksUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	phone, err := usecase.truckRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return phone, nil
}

func (usecase *TrucksUsecase) Delete(ctx context.Context, id int) error {
	err := usecase.truckRepository.CheckByID(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.truckRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrucksUsecase) Update(ctx context.Context, truckDomain *Domain, id int) error {
	err := usecase.truckRepository.CheckByID(ctx, id)
	if err != nil {
		return err
	}
	
	err = usecase.truckRepository.Update(ctx, truckDomain, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrucksUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := usecase.truckRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

