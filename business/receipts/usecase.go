package receipts

import (
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helper/logging"
	"time"
)

type ReceiptUsecase struct {
	receiptRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
	logger         logging.Logger
}

func NewReceiptUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &ReceiptUsecase{
		receiptRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
		logger:         logger,
	}
}