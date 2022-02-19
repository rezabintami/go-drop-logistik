package routes

import (
	_middleware "go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/admins"
	"go-drop-logistik/controllers/agents"
	"go-drop-logistik/controllers/receipts"
	"go-drop-logistik/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	MiddlewareLog   _middleware.ConfigMiddleware
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	AgentController agents.AgentController
	AdminController admins.AdminController
	ReceiptController receipts.ReceiptController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.MiddlewareLog.MiddlewareLogging)

	apiV1 := e.Group("/api/v1")

	//! RESI
	// apiV1.POST("/tracking", cl.UserController.Login)

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	//! AGENTS
	agent := apiV1.Group("/agent")
	agent.POST("/login", cl.AgentController.Login)
	agent.GET("/profile", cl.AgentController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT", "ADMIN"))

	resi := agent.Group("/resi", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	resi.POST("/add", cl.ReceiptController.CreateReceipt)
	resi.GET("", cl.ReceiptController.Fetch)
	resi.GET("/:id", cl.ReceiptController.GetByID)
	// resi.PUT("/:id", cl.AgentController.Login)
	// resi.GET("/:id/finish", cl.AgentController.Login)
	resi.DELETE("/:id/decline", cl.ReceiptController.Delete)

	// manifest := agent.Group("/manifest", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	// manifest.POST("/add", cl.AgentController.Login)
	// manifest.GET("", cl.AgentController.Login)
	// manifest.GET("/:id", cl.AgentController.Login)
	// manifest.PUT("/:id", cl.AgentController.Login)
	// manifest.GET("/:id/finish", cl.AgentController.Login)
	// manifest.DELETE("/:id/decline", cl.AgentController.Login)

	//! ADMINS
	admin := apiV1.Group("/admin")
	admin.POST("/register", cl.AdminController.Register)
	admin.POST("/login", cl.AdminController.Login)
	admin.GET("/profile", cl.AdminController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))

	adminAgent := admin.Group("/agent", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))
	adminAgent.GET("", cl.AdminController.AgentFetch)
	adminAgent.GET("/:id", cl.AdminController.AgentGetByID)
	adminAgent.POST("/add", cl.AdminController.AgentRegister)
	adminAgent.PUT("/:id", cl.AdminController.AgentUpdateByID)

	//! USERS
	user := apiV1.Group("/user")
	user.GET("/", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
}
