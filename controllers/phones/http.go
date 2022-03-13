package phones

import (
	"net/http"
	"strconv"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/phones"
	"go-drop-logistik/controllers/phones/request"
	"go-drop-logistik/controllers/phones/response"
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

func (controller *PhonesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	phone, err := controller.phonesUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(phone))
}

func (controller *PhonesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	phone, err := controller.phonesUsecase.GetAll(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(phone))
}

func (controller *PhonesController) DeletePhone(c echo.Context) error {
	phoneId, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	agentId := middleware.GetUser(c).ID

	err := controller.phonesUsecase.Delete(ctx, agentId, phoneId)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *PhonesController) UpdatePhone(c echo.Context) error {
	ctx := c.Request().Context()
	phoneId, _ := strconv.Atoi(c.Param("id"))

	req := request.Phone{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.phonesUsecase.Update(ctx, req.ToDomain(), phoneId)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}
