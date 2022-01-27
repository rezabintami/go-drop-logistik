package routes

import (
	_middleware "go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	MiddlewareLog  _middleware.ConfigMiddleware
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.MiddlewareLog.MiddlewareLogging)

	apiV1 := e.Group("/api/v1")

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	//! USERS
	user := apiV1.Group("/user")
	user.GET("/", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
}
