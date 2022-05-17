package phones

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type PhoneUsecase struct {
	phoneRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewPhoneUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &PhoneUsecase{
		phoneRepository: ur,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
	}
}

func (usecase *PhoneUsecase) StorePhone(ctx context.Context, phoneDomain *Domain, id int) (int, error) {
	phoneId, err := usecase.phoneRepository.StorePhone(ctx, phoneDomain)
	if err != nil {
		return 0, err
	}

	return phoneId, nil
}

func (usecase *PhoneUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	phone, err := usecase.phoneRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return phone, nil
}

func (usecase *PhoneUsecase) Delete(ctx context.Context, phoneId int) error {
	err := usecase.phoneRepository.CheckByID(ctx, phoneId)
	if err != nil {
		return err
	}

	err = usecase.phoneRepository.Delete(ctx, phoneId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *PhoneUsecase) Update(ctx context.Context, phoneDomain *Domain, phoneId int) error {
	err := usecase.phoneRepository.CheckByID(ctx, phoneId)
	if err != nil {
		return err
	}

	err = usecase.phoneRepository.Update(ctx, phoneDomain, phoneId)
	if err != nil {
		return err
	}

	return nil
}
