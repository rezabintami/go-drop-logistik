package superusers

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/agents"
	"go-drop-logistik/business/superusers"
	"go-drop-logistik/controllers/superusers/request"
	"go-drop-logistik/controllers/superusers/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type SuperuserController struct {
	superuserUsecase superusers.Usecase
	agentUsecase     agents.Usecase
}

func NewSuperuserController(su superusers.Usecase, au agents.Usecase) *SuperuserController {
	return &SuperuserController{
		superuserUsecase: su,
	}
}

func (controller *SuperuserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Superusers{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.superuserUsecase.Register(ctx, req.ToDomain(), false)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *SuperuserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Superusers
	if err := c.Bind(&userLogin); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.superuserUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}

	return base_response.NewSuccessResponse(c, result)
}

func (controller *SuperuserController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.superuserUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}

func (controller *SuperuserController) AgentGetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.agentUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.AgentFromDomain(user))
}

func (controller *SuperuserController) AgentRegister(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Agents{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	controller.agentUsecase.Register(ctx, req.AgentToDomain(), true)

	err := controller.agentUsecase.Register(ctx, req.AgentToDomain(), false)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}
