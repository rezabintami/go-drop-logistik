package manifest

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helper/receipt"
	"time"
)

type ManifestUsecase struct {
	receiptRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewManifestUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ManifestUsecase{
		receiptRepository: ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
	}
}

func (usecase *ManifestUsecase) StoreManifest(ctx context.Context, manifestDomain *Domain) error {
	manifestDomain.Code = receipt.GenerateReceipt()

	err := usecase.receiptRepository.StoreManifest(ctx, manifestDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.receiptRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}

func (usecase *ManifestUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := usecase.receiptRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (usecase *ManifestUsecase) Delete(ctx context.Context, id int) error {
	err := usecase.receiptRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
