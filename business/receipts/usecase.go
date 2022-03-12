package receipts

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helper/code"
	"go-drop-logistik/helper/enum"
	"go-drop-logistik/helper/logging"
	"time"
)

type ReceiptUsecase struct {
	receiptRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
	logger            logging.Logger
}

func NewReceiptUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &ReceiptUsecase{
		receiptRepository: ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
		logger:            logger,
	}
}

func (usecase *ReceiptUsecase) StoreReceipt(ctx context.Context, receiptDomain *Domain) error {
	receiptDomain.Code = code.GenerateReceipt()
	receiptDomain.Status = enum.PROCESS

	err := usecase.receiptRepository.StoreReceipt(ctx, receiptDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ReceiptUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.receiptRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}

func (usecase *ReceiptUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
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

func (usecase *ReceiptUsecase) Delete(ctx context.Context, id int) error {
	err := usecase.receiptRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
