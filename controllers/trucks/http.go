package trucks

import (
	"net/http"
	"strconv"

	"go-drop-logistik/controllers/trucks/request"
	"go-drop-logistik/controllers/trucks/response"
	helpers "go-drop-logistik/helpers"
	"go-drop-logistik/modules/trucks"

	echo "github.com/labstack/echo/v4"
)

type TrucksController struct {
	trucksUsecase trucks.Usecase
}

func NewTrucksController(uc trucks.Usecase) *TrucksController {
	return &TrucksController{
		trucksUsecase: uc,
	}
}

func (controller *TrucksController) StoreTruck(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Trucks{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.trucksUsecase.StoreTruck(ctx, req.ToDomain())
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *TrucksController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	phone, err := controller.trucksUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, phone)
}

func (controller *TrucksController) DeleteTruck(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.trucksUsecase.Delete(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}

func (controller *TrucksController) UpdateTruck(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Trucks{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// validateMessage, validate, err := helpers.Validate(&req)

	// if validate {
	// 	return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	// }

	err := controller.trucksUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}

func (controller *TrucksController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	agents, count, err := controller.trucksUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.FromListDomain(agents, count))
}
