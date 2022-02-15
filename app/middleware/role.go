package middleware

import (
	"errors"
	base_response "go-drop-logistik/helper/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RoleValidation(roles ...string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)

			for _, role := range roles {
				if claims.Role == role {
					return hf(c)
				}
			}
			return base_response.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
		}
	}
}
