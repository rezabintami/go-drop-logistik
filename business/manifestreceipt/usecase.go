package manifestreceipt

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/receipts"
	"go-drop-logistik/helper/enum"
	"time"
)

type ManifestReceiptUsecase struct {
	manifestReceiptRepository Repository
	receiptRepository         receipts.Repository
	contextTimeout            time.Duration
	jwtAuth                   *middleware.ConfigJWT
}

func NewManifestReceiptUsecase(ur Repository, rr receipts.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ManifestReceiptUsecase{
		manifestReceiptRepository: ur,
		receiptRepository:         rr,
		jwtAuth:                   jwtauth,
		contextTimeout:            timeout,
	}
}

func (usecase *ManifestReceiptUsecase) Store(ctx context.Context, ManifestId, ReceiptId int) error {
	err := usecase.manifestReceiptRepository.Store(ctx, ManifestId, ReceiptId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestReceiptUsecase) GetAllByManifestID(ctx context.Context, id int) ([]Domain, error) {
	res, err := usecase.manifestReceiptRepository.GetAllByManifestID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (usecase *ManifestReceiptUsecase) GetByManifestID(ctx context.Context, id int) (Domain, error) {
	res, err := usecase.manifestReceiptRepository.GetByManifestID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (usecase *ManifestReceiptUsecase) DeleteByReceipt(ctx context.Context, ReceiptId int) error {
	err := usecase.manifestReceiptRepository.DeleteByReceipt(ctx, ReceiptId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestReceiptUsecase) DeleteByManifest(ctx context.Context, manifestId int) error {
	err := usecase.manifestReceiptRepository.DeleteByManifest(ctx, manifestId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestReceiptUsecase) UpdateStatusByManifest(ctx context.Context, manifestId int) error {
	res, err := usecase.manifestReceiptRepository.GetAllByManifestID(ctx, manifestId)
	if err != nil {
		return err
	}

	for _, value := range res {

		err := usecase.receiptRepository.Update(ctx, &receipts.Domain{Status: enum.SUCCESS}, value.ReceiptID)
		if err != nil {
			return err
		}
	}
	return nil
}
