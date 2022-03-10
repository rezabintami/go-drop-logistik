package phones

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/phones"
	"go-drop-logistik/controllers/phones/request"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type PhonesController struct {
	phonesUsecase phones.Usecase
}

func NewPhonesController(uc phones.Usecase) *PhonesController {
	return &PhonesController{
		phonesUsecase: uc,
	}
}

func (controller *PhonesController) StorePhone(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	req := request.Phone{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.phonesUsecase.StorePhone(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}
