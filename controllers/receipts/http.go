package receipts

import (
	"net/http"
	"strconv"

	"go-drop-logistik/business/receipts"
	"go-drop-logistik/controllers/receipts/request"
	"go-drop-logistik/controllers/receipts/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type ReceiptController struct {
	receiptUsecase receipts.Usecase
}

func NewReceiptController(uc receipts.Usecase) *ReceiptController {
	return &ReceiptController{
		receiptUsecase: uc,
	}
}

func (controller *ReceiptController) CreateReceipt(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Receipts{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.receiptUsecase.StoreReceipt(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *ReceiptController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.receiptUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
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
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.receiptUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}
