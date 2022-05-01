package phones

import (
	"net/http"
	"strconv"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/phones/request"
	"go-drop-logistik/controllers/phones/response"
	helpers "go-drop-logistik/helpers"
	"go-drop-logistik/modules/phoneagent"
	"go-drop-logistik/modules/phones"

	echo "github.com/labstack/echo/v4"
)

type PhonesController struct {
	phonesUsecase     phones.Usecase
	phoneAgentUsecase phoneagent.Usecase
}

func NewPhonesController(ph phones.Usecase, pa phoneagent.Usecase) *PhonesController {
	return &PhonesController{
		phonesUsecase:     ph,
		phoneAgentUsecase: pa,
	}
}

func (controller *PhonesController) StorePhone(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	req := request.Phone{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}


	phoneId, err := controller.phonesUsecase.StorePhone(ctx, req.ToDomain(), id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.phoneAgentUsecase.Store(ctx, phoneId, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessInsertResponse(c, "Successfully inserted")
}

func (controller *PhonesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	phone, err := controller.phonesUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, response.FromDomain(phone))
}

func (controller *PhonesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	allPhone, err := controller.phoneAgentUsecase.GetAllByAgentID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	var phoneDomain []phones.Domain
	for _, value := range allPhone {
		phone, _ := controller.phonesUsecase.GetByID(ctx, value.PhoneID)
		phoneDomain = append(phoneDomain, phone)
	}

	return helpers.SuccessResponse(c, response.FromListDomain(phoneDomain))
}

func (controller *PhonesController) DeletePhone(c echo.Context) error {
	phoneId, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	agentId := middleware.GetUser(c).ID

	err := controller.phoneAgentUsecase.Delete(ctx, agentId, phoneId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.phonesUsecase.Delete(ctx, phoneId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, "Delete Successfully")
}

func (controller *PhonesController) UpdatePhone(c echo.Context) error {
	ctx := c.Request().Context()
	phoneId, _ := strconv.Atoi(c.Param("id"))

	req := request.Phone{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.phonesUsecase.Update(ctx, req.ToDomain(), phoneId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, "Update Successfully")
}
