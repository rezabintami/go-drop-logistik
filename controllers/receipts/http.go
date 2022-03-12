package receipts

import (
	"net/http"
	"strconv"

	"go-drop-logistik/business/manifestreceipt"
	"go-drop-logistik/business/receipts"
	"go-drop-logistik/controllers/receipts/request"
	"go-drop-logistik/controllers/receipts/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type ReceiptController struct {
	receiptUsecase         receipts.Usecase
	manifestreceiptUsecase manifestreceipt.Usecase
}

func NewReceiptController(uc receipts.Usecase, mr manifestreceipt.Usecase) *ReceiptController {
	return &ReceiptController{
		receiptUsecase:         uc,
		manifestreceiptUsecase: mr,
	}
}

func (controller *ReceiptController) CreateReceipt(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Receipts{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	receiptId, err := controller.receiptUsecase.StoreReceipt(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.manifestreceiptUsecase.Store(ctx, req.ManifestID, receiptId)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *ReceiptController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	receipt, err := controller.receiptUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(receipt))
}

func (controller *ReceiptController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	receipts, count, err := controller.receiptUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(receipts, count))
}

func (controller *ReceiptController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.receiptUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.manifestreceiptUsecase.DeleteByReceipt(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *ReceiptController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Receipts{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.receiptUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}
