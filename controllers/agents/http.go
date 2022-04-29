package agents

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/agents/request"
	"go-drop-logistik/controllers/agents/response"
	"go-drop-logistik/helpers"
	"go-drop-logistik/modules/agents"
	"go-drop-logistik/modules/phoneagent"
	"go-drop-logistik/modules/phones"

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
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.agentUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}

	return helpers.SuccessResponse(c, result)
}

func (controller *AgentController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.agentUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	phone, _ := controller.phoneAgentUsecase.GetAllByAgentID(ctx, id)

	for _, phones := range phone {
		number, _ := controller.phoneUsecase.GetByID(ctx, phones.PhoneID)
		user.Phone = append(user.Phone, number.Phone)
	}

	return helpers.SuccessResponse(c, response.FromDomain(&user))
}
