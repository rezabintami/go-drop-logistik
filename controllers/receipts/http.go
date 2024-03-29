package receipts

import (
	"net/http"
	"strconv"

	"go-drop-logistik/controllers/receipts/request"
	"go-drop-logistik/controllers/receipts/response"
	helpers "go-drop-logistik/helpers"
	"go-drop-logistik/modules/manifestreceipt"
	"go-drop-logistik/modules/receipts"
	"go-drop-logistik/modules/trackmanifest"

	echo "github.com/labstack/echo/v4"
)

type ReceiptController struct {
	receiptUsecase         receipts.Usecase
	manifestreceiptUsecase manifestreceipt.Usecase
	trackManifestUsecase   trackmanifest.Usecase
}

func NewReceiptController(uc receipts.Usecase, mr manifestreceipt.Usecase, tr trackmanifest.Usecase) *ReceiptController {
	return &ReceiptController{
		receiptUsecase:         uc,
		manifestreceiptUsecase: mr,
		trackManifestUsecase:   tr,
	}
}

func (controller *ReceiptController) CreateReceipt(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Receipts{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	receiptId, err := controller.receiptUsecase.StoreReceipt(ctx, req.ToDomain())
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if req.ManifestID != 0 {
		err = controller.manifestreceiptUsecase.Store(ctx, req.ManifestID, receiptId)
		if err != nil {
			return helpers.ErrorResponse(c, http.StatusBadRequest, err)
		}
	}

	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *ReceiptController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	receipt, err := controller.receiptUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	manifestId, _ := controller.manifestreceiptUsecase.GetByReceiptID(ctx, receipt.ID)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.FromDomainManifest(receipt, manifestId))
}

func (controller *ReceiptController) GetByCode(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.TrackingReceipts{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	receipt, err := controller.receiptUsecase.GetByCode(ctx, req.Code)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	manifestId, err := controller.manifestreceiptUsecase.GetByReceiptID(ctx, receipt.ID)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	tracks, err := controller.trackManifestUsecase.GetAllByManifestID(ctx, manifestId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	for _, value := range tracks {
		receipt.Tracks = append(receipt.Tracks, *value.Track)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.TrackFromDomain(receipt))
}

func (controller *ReceiptController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	receipts, count, err := controller.receiptUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.FromListDomain(receipts, count))
}

func (controller *ReceiptController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.receiptUsecase.Delete(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.manifestreceiptUsecase.DeleteByReceipt(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}

func (controller *ReceiptController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Receipts{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// validateMessage, validate, err := helpers.Validate(&req)

	// if validate {
	// 	return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	// }

	err := controller.receiptUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if req.ManifestID != 0 {
		err = controller.manifestreceiptUsecase.Store(ctx, req.ManifestID, id)
		if err != nil {
			return helpers.ErrorResponse(c, http.StatusBadRequest, err)
		}
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}
