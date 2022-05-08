package app_plugins

import (
	_config "go-drop-logistik/app/config"
	_middleware "go-drop-logistik/app/middleware"
	_routes "go-drop-logistik/app/routes"
	_helpers "go-drop-logistik/helpers"
	"time"

	_userController "go-drop-logistik/controllers/users"
	_userRepo "go-drop-logistik/drivers/databases/users"
	_userUsecase "go-drop-logistik/modules/users"

	_agentController "go-drop-logistik/controllers/agents"
	_agentRepo "go-drop-logistik/drivers/databases/agents"
	_agentUsecase "go-drop-logistik/modules/agents"

	_phoneAgentRepo "go-drop-logistik/drivers/databases/phoneagent"
	_phoneAgentUsecase "go-drop-logistik/modules/phoneagent"

	_manifestReceiptRepo "go-drop-logistik/drivers/databases/manifestreceipt"
	_manifestReceiptUsecase "go-drop-logistik/modules/manifestreceipt"

	_trackManifestRepo "go-drop-logistik/drivers/databases/trackmanifest"
	_trackManifestUsecase "go-drop-logistik/modules/trackmanifest"

	_phoneController "go-drop-logistik/controllers/phones"
	_phoneRepo "go-drop-logistik/drivers/databases/phones"
	_phoneUsecase "go-drop-logistik/modules/phones"

	_adminController "go-drop-logistik/controllers/admins"
	_adminRepo "go-drop-logistik/drivers/databases/admins"
	_adminUsecase "go-drop-logistik/modules/admins"

	_receiptController "go-drop-logistik/controllers/receipts"
	_receiptRepo "go-drop-logistik/drivers/databases/receipts"
	_receiptUsecase "go-drop-logistik/modules/receipts"

	_trackController "go-drop-logistik/controllers/tracks"
	_trackRepo "go-drop-logistik/drivers/databases/tracks"
	_trackUsecase "go-drop-logistik/modules/tracks"

	_manifestController "go-drop-logistik/controllers/manifest"
	_manifestRepo "go-drop-logistik/drivers/databases/manifest"
	_manifestUsecase "go-drop-logistik/modules/manifest"

	_truckController "go-drop-logistik/controllers/trucks"
	_truckRepo "go-drop-logistik/drivers/databases/trucks"
	_truckUsecase "go-drop-logistik/modules/trucks"

	_driverController "go-drop-logistik/controllers/drivers"
	_driverRepo "go-drop-logistik/drivers/databases/drivers"
	_driverUsecase "go-drop-logistik/modules/drivers"

	"github.com/jinzhu/gorm"
)

type ConfigurationPlugins struct {
	Postgres_DB   *gorm.DB
	Logger        _helpers.Logger
	MiddlewareLog _middleware.ConfigMiddleware
}

func (route *ConfigurationPlugins) RoutePlugins() _routes.ControllerList {

	configJWT := _middleware.ConfigJWT{
		SecretJWT:        _config.GetConfiguration("jwt.access_token"),
		RefreshSecretJWT: _config.GetConfiguration("jwt.refresh_token"),
		ExpiresDuration:  _helpers.ConvertStringtoInt(_config.GetConfiguration("jwt.expired")),
	}

	timeoutContext := time.Duration(_helpers.ConvertStringtoInt(_config.GetConfiguration("server.timeout"))) * time.Second

	//! REPOSITORY
	userRepo := _userRepo.NewPostgreUsersRepository(route.Postgres_DB)
	phoneAgentRepo := _phoneAgentRepo.NewPostgrePhoneAgentRepository(route.Postgres_DB)
	phoneRepo := _phoneRepo.NewpostgrePhoneRepository(route.Postgres_DB)
	agentRepo := _agentRepo.NewPostgreAgentRepository(route.Postgres_DB)
	adminRepo := _adminRepo.NewPostgreAdminRepository(route.Postgres_DB)
	receiptRepo := _receiptRepo.NewPostgreReceiptRepository(route.Postgres_DB)
	manifestReceiptRepo := _manifestReceiptRepo.NewPostgreManifestReceiptRepository(route.Postgres_DB)
	manifestRepo := _manifestRepo.NewPostgreManifestRepository(route.Postgres_DB)
	truckRepo := _truckRepo.NewPostgreTruckRepository(route.Postgres_DB)
	driverRepo := _driverRepo.NewPostgreDriverRepository(route.Postgres_DB)
	trackRepo := _trackRepo.NewPostgreTrackRepository(route.Postgres_DB)
	trackManifestRepo := _trackManifestRepo.NewPostgreTrackManifestRepository(route.Postgres_DB)

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
	trackCtrl := _trackController.NewTracksController(trackUsecase, trackManifestUsecase, manifestReceiptUsecase, manifestUsecase)

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
	}
}
