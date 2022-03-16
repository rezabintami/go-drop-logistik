package agents

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/agents"
	"go-drop-logistik/business/phoneagent"
	"go-drop-logistik/business/phones"
	"go-drop-logistik/controllers/agents/request"
	"go-drop-logistik/controllers/agents/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type AgentController struct {
	agentUsecase      agents.Usecase
	phoneAgentUsecase phoneagent.Usecase
	phoneUsecase      phones.Usecase
}

func NewAgentController(ag agents.Usecase, pa phoneagent.Usecase, ph phones.Usecase) *AgentController {
	return &AgentController{
		agentUsecase:      ag,
		phoneAgentUsecase: pa,
		phoneUsecase:      ph,
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

	phone, _ := controller.phoneAgentUsecase.GetAllByAgentID(ctx, id)

	for _, phones := range phone {
		number, _ := controller.phoneUsecase.GetByID(ctx, phones.PhoneID)
		user.Phone = append(user.Phone, number.Phone)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}
