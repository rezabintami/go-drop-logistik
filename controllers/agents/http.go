package agents

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/agents"
	"go-drop-logistik/controllers/agents/request"
	"go-drop-logistik/controllers/agents/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type AgentController struct {
	agentUsecase agents.Usecase
}

func NewAgentController(uc agents.Usecase) *AgentController {
	return &AgentController{
		agentUsecase: uc,
	}
}

func (controller *AgentController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Agents
	if err := c.Bind(&userLogin); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.agentUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}

	return base_response.NewSuccessResponse(c, result)
}

func (controller *AgentController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.agentUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(&user))
}
