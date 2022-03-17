package routes

import (
	_middleware "go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/admins"
	"go-drop-logistik/controllers/agents"
	"go-drop-logistik/controllers/drivers"
	"go-drop-logistik/controllers/manifest"
	"go-drop-logistik/controllers/phones"
	"go-drop-logistik/controllers/receipts"
	"go-drop-logistik/controllers/trucks"
	"go-drop-logistik/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	MiddlewareLog      _middleware.ConfigMiddleware
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	AgentController    agents.AgentController
	AdminController    admins.AdminController
	ReceiptController  receipts.ReceiptController
	PhoneController    phones.PhonesController
	ManifestController manifest.ManifestController
	TruckController    trucks.TrucksController
	DriverController   drivers.DriversController
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
	resi.PUT("/:id", cl.ReceiptController.Update)
	resi.DELETE("/:id/decline", cl.ReceiptController.Delete)

	manifest := agent.Group("/manifest", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	manifest.POST("/add", cl.ManifestController.CreateManifest)
	manifest.GET("", cl.ManifestController.Fetch)
	manifest.GET("/:id", cl.ManifestController.GetByID)
	manifest.PUT("/:id", cl.ManifestController.Update)
	manifest.PUT("/:id/finish", cl.ManifestController.UpdateStatus)
	manifest.DELETE("/:id/decline", cl.ManifestController.Delete)

	agentPhone := agent.Group("/phone", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	agentPhone.POST("/add", cl.PhoneController.StorePhone)
	agentPhone.GET("", cl.PhoneController.GetAll)
	agentPhone.DELETE("/:id", cl.PhoneController.DeletePhone)
	agentPhone.PUT("/:id", cl.PhoneController.UpdatePhone)

	drivers := agent.Group("/driver", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	drivers.POST("/add", cl.DriverController.Store)
	drivers.GET("/:id", cl.DriverController.GetByID)
	drivers.DELETE("/:id", cl.DriverController.Delete)
	drivers.PUT("/:id", cl.DriverController.Update)

	trucks := agent.Group("/truck", middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("AGENT"))
	trucks.POST("/add", cl.TruckController.StoreTruck)
	trucks.GET("", cl.TruckController.Fetch)
	trucks.DELETE("/:id", cl.TruckController.DeleteTruck)
	trucks.PUT("/:id", cl.TruckController.UpdateTruck)

	//! ADMINS
	admin := apiV1.Group("/admin")
	admin.POST("/register", cl.AdminController.Register)
	admin.POST("/login", cl.AdminController.Login)
	admin.GET("/profile", cl.AdminController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))

	adminAgent := admin.Group("/agent") //, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))
	adminAgent.GET("", cl.AdminController.AgentFetch)
	adminAgent.GET("/:id", cl.AdminController.AgentGetByID)
	adminAgent.POST("/add", cl.AdminController.AgentRegister)
	adminAgent.PUT("/:id", cl.AdminController.AgentUpdateByID)

	//! USERS
	user := apiV1.Group("/user")
	user.GET("/", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
}
