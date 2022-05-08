package routes

import (
	_middleware "go-drop-logistik/app/middleware"
	"go-drop-logistik/controllers/admins"
	"go-drop-logistik/controllers/agents"
	"go-drop-logistik/controllers/drivers"
	"go-drop-logistik/controllers/manifest"
	"go-drop-logistik/controllers/phones"
	"go-drop-logistik/controllers/receipts"
	"go-drop-logistik/controllers/tracks"
	"go-drop-logistik/controllers/trucks"
	"go-drop-logistik/controllers/users"

	_config "go-drop-logistik/app/config"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	TrackController    tracks.TracksController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.MiddlewareLog.MiddlewareLogging)
	e.HTTPErrorHandler = _middleware.CustomHTTPErrorHandler

	// showing swagger files
	if _config.GetConfiguration("app.env") != "PROD" {
		e.Static("/files", "files")
		url := echoSwagger.URL("/files/swagger.yaml")
		e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
	}
	apiV1 := e.Group("/api/v1")

	//! RESI
	apiV1.POST("/tracking", cl.ReceiptController.GetByCode)

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)
	// auth.POST("/logout", cl.UserController.Logout)
	// auth.POST("/refresh", cl.UserController.Refresh)

	//! USERS
	user := apiV1.Group("/user")
	user.GET("/", cl.UserController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	
	//! AGENTS
	agent := apiV1.Group("/agent")
	agent.POST("/login", cl.AgentController.Login)
	// agent.POST("/refresh", cl.AgentController.Refresh)
	// agent.POST("/logout", cl.AgentController.Logout)
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
	manifest.PUT("/:id/finished", cl.ManifestController.FinishManifest)
	manifest.DELETE("/:id/decline", cl.ManifestController.Delete)

	manifestTrack := manifest.Group("/:manifestId/track")
	manifestTrack.POST("/add", cl.TrackController.CreateTrack)
	manifestTrack.PUT("/:trackId", cl.TrackController.UpdateTrack)
	manifestTrack.DELETE("/:trackId", cl.TrackController.DeleteTrack)

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
	// admin.POST("/refresh", cl.AdminController.Refresh)
	// admin.POST("/logout", cl.AdminController.Logout)
	admin.GET("/profile", cl.AdminController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))

	adminAgent := admin.Group("/agent") //, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("ADMIN"))
	adminAgent.GET("", cl.AdminController.AgentFetch)
	adminAgent.GET("/:id", cl.AdminController.AgentGetByID)
	adminAgent.POST("/add", cl.AdminController.AgentRegister)
	adminAgent.PUT("/:id", cl.AdminController.AgentUpdateByID)
	adminAgent.DELETE("/:id", cl.AdminController.AgentDeleteByID)
}
