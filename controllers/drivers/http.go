package drivers

import (
	"net/http"
	"strconv"

	"go-drop-logistik/controllers/drivers/request"
	"go-drop-logistik/controllers/drivers/response"
	helpers "go-drop-logistik/helpers"
	"go-drop-logistik/modules/drivers"

	echo "github.com/labstack/echo/v4"
)

type DriversController struct {
	driversUsecase drivers.Usecase
}

func NewDriversController(uc drivers.Usecase) *DriversController {
	return &DriversController{
		driversUsecase: uc,
	}
}

func (controller *DriversController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Drivers{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.driversUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return helpers.SuccessInsertResponse(c, "Successfully inserted")
}

func (controller *DriversController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	driver, err := controller.driversUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, response.FromDomain(&driver))
}

func (controller *DriversController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.driversUsecase.Delete(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, "Delete Successfully")
}

func (controller *DriversController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Drivers{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// validateMessage, validate, err := helpers.Validate(&req)

	// if validate {
	// 	return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	// }

	err := controller.driversUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, "Update Successfully")
}
