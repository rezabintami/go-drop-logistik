package users

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/users/request"
	"go-drop-logistik/controllers/users/response"
	base_response "go-drop-logistik/helpers"
	"go-drop-logistik/modules/users"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.userUsecase.Register(ctx, req.ToDomain(), false)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Users
	if err := c.Bind(&userLogin); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.userUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}

	return base_response.NewSuccessResponse(c, result)
}

func (controller *UserController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID
	
	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}
