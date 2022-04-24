package main

import (
	"fmt"
	"os"

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
	

	postgres_db := _dbPostgresDriver.InitialPostgresDB()

	// Init Seeding
	err := seeder.Seeder(postgres_db)

	// Init Validation
	validation.Init()

	log.Println(err)

	// configJWT := _middleware.ConfigJWT{
	// 	SecretJWT:       configApp.JWT.Secret,
	// 	ExpiresDuration: configApp.JWT.Expired,
	// }

	// timeoutContext := time.Duration(configApp.JWT.Expired) * time.Second

	e := echo.New()

	logger := logging.NewLogger()

	middlewareLog := _middleware.NewMiddleware(logger)

	plugins := _plugins.ConfigurationPlugins(postgres_db,logger,middlewareLog)

	route := plugins.RoutePlugins()	
	route.RouteRegister(e)
		
	// //! REPO
	// userRepo := _userRepo.NewMySQLUserRepository(postgres_db)
	// phoneAgentRepo := _phoneAgentRepo.NewMySQLPhoneAgentRepository(postgres_db)
	// phoneRepo := _phoneRepo.NewMySQLPhoneRepository(postgres_db)
	// agentRepo := _agentRepo.NewMySQLAgentRepository(postgres_db)
	// adminRepo := _adminRepo.NewMySQLAdminRepository(postgres_db)
	// receiptRepo := _receiptRepo.NewMySQLReceiptRepository(postgres_db)
	// manifestReceiptRepo := _manifestReceiptRepo.NewMySQLManifestReceiptRepository(postgres_db)
	// manifestRepo := _manifestRepo.NewMySQLManifestRepository(postgres_db)
	// truckRepo := _truckRepo.NewMySQLTruckRepository(postgres_db)
	// driverRepo := _driverRepo.NewMySQLDriverRepository(postgres_db)
	// trackRepo := _trackRepo.NewMySQLTrackRepository(postgres_db)
	// trackManifestRepo := _trackManifestRepo.NewMySQLTrackManifestRepository(postgres_db)

	// //! USECASE
	// userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	// phoneUsecase := _phoneUsecase.NewPhoneUsecase(phoneRepo, &configJWT, timeoutContext)
	// agentUsecase := _agentUsecase.NewAgentUsecase(agentRepo, &configJWT, timeoutContext)
	// adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo, &configJWT, timeoutContext, logger)
	// receiptUsecase := _receiptUsecase.NewReceiptUsecase(receiptRepo, &configJWT, timeoutContext, logger)
	// manifestReceiptUsecase := _manifestReceiptUsecase.NewManifestReceiptUsecase(manifestReceiptRepo, receiptRepo, &configJWT, timeoutContext)
	// manifestUsecase := _manifestUsecase.NewManifestUsecase(manifestRepo, &configJWT, timeoutContext)
	// truckUsecase := _truckUsecase.NewTrucksUsecase(truckRepo, &configJWT, timeoutContext)
	// driverUsecase := _driverUsecase.NewDriverUsecase(driverRepo, &configJWT, timeoutContext)
	// phoneAgentUsecase := _phoneAgentUsecase.NewPhoneAgentUsecase(phoneAgentRepo, &configJWT, timeoutContext)
	// trackUsecase := _trackUsecase.NewTrackUsecase(trackRepo, &configJWT, timeoutContext)
	// trackManifestUsecase := _trackManifestUsecase.NewTrackManifestUsecase(trackManifestRepo, &configJWT, timeoutContext)

	// //! CONTROLLER
	// userCtrl := _userController.NewUserController(userUsecase)
	// phoneCtrl := _phoneController.NewPhonesController(phoneUsecase, phoneAgentUsecase)
	// agentCtrl := _agentController.NewAgentController(agentUsecase, phoneAgentUsecase, phoneUsecase)
	// adminCtrl := _adminController.NewAdminController(adminUsecase, agentUsecase)
	// receiptCtrl := _receiptController.NewReceiptController(receiptUsecase, manifestReceiptUsecase, trackManifestUsecase)
	// manifestCtrl := _manifestController.NewManifestController(manifestUsecase, manifestReceiptUsecase, trackManifestUsecase)
	// truckCtrl := _truckController.NewTrucksController(truckUsecase)
	// driverCtrl := _driverController.NewDriversController(driverUsecase)
	// trackCtrl := _trackController.NewTracksController(trackUsecase, trackManifestUsecase)

	

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	fmt.Println("Location :", loc, " Time :", now.Format(time.RFC3339))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	logger.LogServer("Server is running").Info("Server started at port ", port)
	log.Println("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))

}
