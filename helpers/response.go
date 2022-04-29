package helpers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Message string   `json:"message"`
		Errors  []string `json:"error,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func SuccessInsertResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Message = "Success Insert"
	response.Data = param

	return c.JSON(http.StatusCreated, response)
}

func ErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Message = "Something wrong"
	response.Meta.Errors = []string{err.Error()}

	return c.JSON(status, response)
}

func ErrorValidateResponse(c echo.Context, status int, err error, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Message = "Something wrong"
	response.Data = param

	return c.JSON(status, response)
}
