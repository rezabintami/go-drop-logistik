package admins

import (
	"net/http"
	"strconv"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/admins/request"
	"go-drop-logistik/controllers/admins/response"
	"go-drop-logistik/helpers"
	"go-drop-logistik/modules/admins"
	"go-drop-logistik/modules/agents"

	echo "github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUsecase admins.Usecase
	agentUsecase agents.Usecase
}

func NewAdminController(su admins.Usecase, au agents.Usecase) *AdminController {
	return &AdminController{
		adminUsecase: su,
		agentUsecase: au,
	}
}

func (controller *AdminController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.adminUsecase.Register(ctx, req.ToDomain(), false)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *AdminController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Admins
	if err := c.Bind(&userLogin); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.adminUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}

	return helpers.SuccessResponse(c, http.StatusOK, result)
}

func (controller *AdminController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.adminUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.FromDomain(user))
}

func (controller *AdminController) AgentGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.agentUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.AgentFromDomain(user))
}

func (controller *AdminController) AgentRegister(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Agents{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.agentUsecase.Register(ctx, req.AgentToDomain(), false)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *AdminController) AgentFetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	agents, count, err := controller.agentUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.AgentFromListDomain(agents, count))
}

func (controller *AdminController) AgentUpdateByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	req := request.Agents{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// validateMessage, validate, err := helpers.Validate(&req)

	// if validate {
	// 	return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	// }

	err := controller.agentUsecase.Update(ctx, req.AgentToDomain(), idInt)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	user, err := controller.agentUsecase.GetByID(ctx, idInt)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.AgentFromDomain(user))
}
