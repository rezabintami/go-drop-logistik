package users

import (
	"net/http"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/users/request"
	"go-drop-logistik/controllers/users/response"
	helpers "go-drop-logistik/helpers"
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
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	err = controller.userUsecase.Register(ctx, req.ToDomain(), false)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Users
	if err := c.Bind(&userLogin); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	accessToken, refreshToken, err := controller.userUsecase.Login(ctx, userLogin.Email, userLogin.Password, false)

	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.TokenFromDomain(accessToken, refreshToken))
}

func (controller *UserController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUser(c).ID

	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, response.FromDomain(user))
}
