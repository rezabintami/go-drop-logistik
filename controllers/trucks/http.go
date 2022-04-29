package trucks

import (
	"net/http"
	"strconv"

	"go-drop-logistik/controllers/trucks/request"
	"go-drop-logistik/controllers/trucks/response"
	base_response "go-drop-logistik/helpers"
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
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.trucksUsecase.StoreTruck(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *TrucksController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	phone, err := controller.trucksUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, phone)
}

func (controller *TrucksController) DeleteTruck(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.trucksUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *TrucksController) UpdateTruck(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Trucks{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.trucksUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}

func (controller *TrucksController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	agents, count, err := controller.trucksUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(agents, count))
}
