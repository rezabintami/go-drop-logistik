package app_plugins

import (
	_config "go-drop-logistik/app/config"
	_middleware "go-drop-logistik/app/middleware"
	_routes "go-drop-logistik/app/routes"

	"time"

	"go-drop-logistik/helper/logging"

	_userUsecase "go-drop-logistik/business/users"
	_userController "go-drop-logistik/controllers/users"
	_userRepo "go-drop-logistik/drivers/databases/users"

	_agentUsecase "go-drop-logistik/business/agents"
	_agentController "go-drop-logistik/controllers/agents"
	_agentRepo "go-drop-logistik/drivers/databases/agents"

	_phoneAgentUsecase "go-drop-logistik/business/phoneagent"
	_phoneAgentRepo "go-drop-logistik/drivers/databases/phoneagent"

	_manifestReceiptUsecase "go-drop-logistik/business/manifestreceipt"
	_manifestReceiptRepo "go-drop-logistik/drivers/databases/manifestreceipt"

	_trackManifestUsecase "go-drop-logistik/business/trackmanifest"
	_trackManifestRepo "go-drop-logistik/drivers/databases/trackmanifest"

	_phoneUsecase "go-drop-logistik/business/phones"
	_phoneController "go-drop-logistik/controllers/phones"
	_phoneRepo "go-drop-logistik/drivers/databases/phones"

	_adminUsecase "go-drop-logistik/business/admins"
	_adminController "go-drop-logistik/controllers/admins"
	_adminRepo "go-drop-logistik/drivers/databases/admins"

	_receiptUsecase "go-drop-logistik/business/receipts"
	_receiptController "go-drop-logistik/controllers/receipts"
	_receiptRepo "go-drop-logistik/drivers/databases/receipts"

	_trackUsecase "go-drop-logistik/business/tracks"
	_trackController "go-drop-logistik/controllers/tracks"
	_trackRepo "go-drop-logistik/drivers/databases/tracks"

	_manifestUsecase "go-drop-logistik/business/manifest"
	_manifestController "go-drop-logistik/controllers/manifest"
	_manifestRepo "go-drop-logistik/drivers/databases/manifest"

	_truckUsecase "go-drop-logistik/business/trucks"
	_truckController "go-drop-logistik/controllers/trucks"
	_truckRepo "go-drop-logistik/drivers/databases/trucks"

	_driverUsecase "go-drop-logistik/business/drivers"
	_driverController "go-drop-logistik/controllers/drivers"
	_driverRepo "go-drop-logistik/drivers/databases/drivers"

	"github.com/jinzhu/gorm"
)

type ConfigurationPlugins struct {
	ConfigApp     _config.Config
	Postgres_DB   *gorm.DB
	Logger        logging.Logger
	MiddlewareLog _middleware.ConfigMiddleware
}

func (route *ConfigurationPlugins) RoutePlugins() _routes.ControllerList {


	

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       route.ConfigApp.JWT.Secret,
		ExpiresDuration: route.ConfigApp.JWT.Expired,
	}

	timeoutContext := time.Duration(route.ConfigApp.JWT.Expired) * time.Second

	//! REPO
	userRepo := _userRepo.NewMySQLUserRepository(route.Postgres_DB)
	phoneAgentRepo := _phoneAgentRepo.NewMySQLPhoneAgentRepository(route.Postgres_DB)
	phoneRepo := _phoneRepo.NewMySQLPhoneRepository(route.Postgres_DB)
	agentRepo := _agentRepo.NewMySQLAgentRepository(route.Postgres_DB)
	adminRepo := _adminRepo.NewMySQLAdminRepository(route.Postgres_DB)
	receiptRepo := _receiptRepo.NewMySQLReceiptRepository(route.Postgres_DB)
	manifestReceiptRepo := _manifestReceiptRepo.NewMySQLManifestReceiptRepository(route.Postgres_DB)
	manifestRepo := _manifestRepo.NewMySQLManifestRepository(route.Postgres_DB)
	truckRepo := _truckRepo.NewMySQLTruckRepository(route.Postgres_DB)
	driverRepo := _driverRepo.NewMySQLDriverRepository(route.Postgres_DB)
	trackRepo := _trackRepo.NewMySQLTrackRepository(route.Postgres_DB)
	trackManifestRepo := _trackManifestRepo.NewMySQLTrackManifestRepository(route.Postgres_DB)

	//! USECASE
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	phoneUsecase := _phoneUsecase.NewPhoneUsecase(phoneRepo, &configJWT, timeoutContext)
	agentUsecase := _agentUsecase.NewAgentUsecase(agentRepo, &configJWT, timeoutContext)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo, &configJWT, timeoutContext, route.Logger)
	receiptUsecase := _receiptUsecase.NewReceiptUsecase(receiptRepo, &configJWT, timeoutContext, route.Logger)
	manifestReceiptUsecase := _manifestReceiptUsecase.NewManifestReceiptUsecase(manifestReceiptRepo, receiptRepo, &configJWT, timeoutContext)
	manifestUsecase := _manifestUsecase.NewManifestUsecase(manifestRepo, &configJWT, timeoutContext)
	truckUsecase := _truckUsecase.NewTrucksUsecase(truckRepo, &configJWT, timeoutContext)
	driverUsecase := _driverUsecase.NewDriverUsecase(driverRepo, &configJWT, timeoutContext)
	phoneAgentUsecase := _phoneAgentUsecase.NewPhoneAgentUsecase(phoneAgentRepo, &configJWT, timeoutContext)
	trackUsecase := _trackUsecase.NewTrackUsecase(trackRepo, &configJWT, timeoutContext)
	trackManifestUsecase := _trackManifestUsecase.NewTrackManifestUsecase(trackManifestRepo, &configJWT, timeoutContext)

	//! CONTROLLER
	userCtrl := _userController.NewUserController(userUsecase)
	phoneCtrl := _phoneController.NewPhonesController(phoneUsecase, phoneAgentUsecase)
	agentCtrl := _agentController.NewAgentController(agentUsecase, phoneAgentUsecase, phoneUsecase)
	adminCtrl := _adminController.NewAdminController(adminUsecase, agentUsecase)
	receiptCtrl := _receiptController.NewReceiptController(receiptUsecase, manifestReceiptUsecase, trackManifestUsecase)
	manifestCtrl := _manifestController.NewManifestController(manifestUsecase, manifestReceiptUsecase, trackManifestUsecase)
	truckCtrl := _truckController.NewTrucksController(truckUsecase)
	driverCtrl := _driverController.NewDriversController(driverUsecase)
	trackCtrl := _trackController.NewTracksController(trackUsecase, trackManifestUsecase)

	return _routes.ControllerList{
		MiddlewareLog:      route.MiddlewareLog,
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		AgentController:    *agentCtrl,
		AdminController:    *adminCtrl,
		ReceiptController:  *receiptCtrl,
		PhoneController:    *phoneCtrl,
		ManifestController: *manifestCtrl,
		TruckController:    *truckCtrl,
		DriverController:   *driverCtrl,
		TrackController:    *trackCtrl,
		ConfigApp:          route.ConfigApp,
	}
}
