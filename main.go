package main

import (
	"fmt"
	"os"

	_config "go-drop-logistik/app/config"
	_dbPostgresDriver "go-drop-logistik/drivers/postgres"
	"go-drop-logistik/helper/logging"
	"go-drop-logistik/helper/validation"

	_middleware "go-drop-logistik/app/middleware"
	// _routes "go-drop-logistik/app/routes"

	_plugins "go-drop-logistik/app/plugins"

	// _userUsecase "go-drop-logistik/business/users"
	// _userController "go-drop-logistik/controllers/users"
	// _userRepo "go-drop-logistik/drivers/databases/users"

	// _agentUsecase "go-drop-logistik/business/agents"
	// _agentController "go-drop-logistik/controllers/agents"
	// _agentRepo "go-drop-logistik/drivers/databases/agents"

	// _phoneAgentUsecase "go-drop-logistik/business/phoneagent"
	// _phoneAgentRepo "go-drop-logistik/drivers/databases/phoneagent"

	// _manifestReceiptUsecase "go-drop-logistik/business/manifestreceipt"
	// _manifestReceiptRepo "go-drop-logistik/drivers/databases/manifestreceipt"

	// _trackManifestUsecase "go-drop-logistik/business/trackmanifest"
	// _trackManifestRepo "go-drop-logistik/drivers/databases/trackmanifest"

	// _phoneUsecase "go-drop-logistik/business/phones"
	// _phoneController "go-drop-logistik/controllers/phones"
	// _phoneRepo "go-drop-logistik/drivers/databases/phones"

	// _adminUsecase "go-drop-logistik/business/admins"
	// _adminController "go-drop-logistik/controllers/admins"
	// _adminRepo "go-drop-logistik/drivers/databases/admins"

	// _receiptUsecase "go-drop-logistik/business/receipts"
	// _receiptController "go-drop-logistik/controllers/receipts"
	// _receiptRepo "go-drop-logistik/drivers/databases/receipts"

	// _trackUsecase "go-drop-logistik/business/tracks"
	// _trackController "go-drop-logistik/controllers/tracks"
	// _trackRepo "go-drop-logistik/drivers/databases/tracks"

	// _manifestUsecase "go-drop-logistik/business/manifest"
	// _manifestController "go-drop-logistik/controllers/manifest"
	// _manifestRepo "go-drop-logistik/drivers/databases/manifest"

	// _truckUsecase "go-drop-logistik/business/trucks"
	// _truckController "go-drop-logistik/controllers/trucks"
	// _truckRepo "go-drop-logistik/drivers/databases/trucks"

	// _driverUsecase "go-drop-logistik/business/drivers"
	// _driverController "go-drop-logistik/controllers/drivers"
	// _driverRepo "go-drop-logistik/drivers/databases/drivers"

	"log"
	"time"

	"go-drop-logistik/cli/seeder"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configApp := _config.GetConfig()

	fmt.Println("Server is running on port :" + configApp.Server.Address)

	fmt.Println("User :", configApp.Postgres.User)
	fmt.Println("Host :", configApp.Postgres.Host)
	fmt.Println("Port :", configApp.Postgres.Port)
	fmt.Println("Name :", configApp.Postgres.Name)

	postgres_db := _dbPostgresDriver.InitialPostgresDB()

	// Init Seeding
	err := seeder.Seeder(postgres_db)
	log.Println(err)

	// Init Validation
	validation.Init()

	e := echo.New()
	logger := logging.NewLogger()
	middlewareLog := _middleware.NewMiddleware(logger)

	plugins := _plugins.ConfigurationPlugins{
		ConfigApp:     configApp,
		Postgres_DB:   postgres_db,
		Logger:        logger,
		MiddlewareLog: middlewareLog,
	}

	route := plugins.RoutePlugins()
	route.RouteRegister(e)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	fmt.Println("Location :", loc, " Time :", now.Format(time.RFC3339))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("App :", configApp.App.Env)
	log.Println("Debug :", configApp.App.Debug)
	log.Println("App Version :", configApp.App.Version)

	logger.LogServer("Server is running").Info("Server started at port ", port)
	log.Println("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))

}
