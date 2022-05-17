package manifestreceipt

import (
	"context"
	"errors"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/modules/receipts"
	"log"
	"sync"
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

func (usecase *ManifestReceiptUsecase) GetByReceiptID(ctx context.Context, id int) (int, error) {
	res, err := usecase.manifestReceiptRepository.GetByReceiptID(ctx, id)
	if err != nil {
		return 0, err
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

func (usecase *ManifestReceiptUsecase) UpdateStatusByManifest(ctx context.Context, status string, manifestId int) error {
	res, err := usecase.manifestReceiptRepository.GetAllByManifestID(ctx, manifestId)
	if err != nil {
		return err
	}

	message := make(chan error, len(res))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range res {
			err := usecase.receiptRepository.Update(ctx, &receipts.Domain{Status: status}, value.ReceiptID)
			if err != nil {
				log.Println("[error] manifestreceipts.usecase.UpdateStatusByManifest : failed to execute update all status manifest", err)
				message <- err
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()

	select {
	case check := <-message:
		return errors.New("update all status manifest failed : " + check.Error())
	default:
		return nil
	}
}
