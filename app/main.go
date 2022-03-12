package main

import (
	"fmt"
	"os"

	_config "go-drop-logistik/app/config"
	_dbMysqlDriver "go-drop-logistik/drivers/mysql"
	"go-drop-logistik/helper/logging"

	_middleware "go-drop-logistik/app/middleware"
	_routes "go-drop-logistik/app/routes"

	_userUsecase "go-drop-logistik/business/users"
	_userController "go-drop-logistik/controllers/users"
	_userRepo "go-drop-logistik/drivers/databases/users"

	_agentUsecase "go-drop-logistik/business/agents"
	_agentController "go-drop-logistik/controllers/agents"
	_agentRepo "go-drop-logistik/drivers/databases/agents"

	_phoneAgentRepo "go-drop-logistik/drivers/databases/phoneagent"

	_manifestreceiptUsecase "go-drop-logistik/business/manifestreceipt"
	_manifestreceiptRepo "go-drop-logistik/drivers/databases/manifestreceipt"

	_phoneUsecase "go-drop-logistik/business/phones"
	_phoneController "go-drop-logistik/controllers/phones"
	_phoneRepo "go-drop-logistik/drivers/databases/phones"

	_adminUsecase "go-drop-logistik/business/admins"
	_adminController "go-drop-logistik/controllers/admins"
	_adminRepo "go-drop-logistik/drivers/databases/admins"

	_receiptUsecase "go-drop-logistik/business/receipts"
	_receiptController "go-drop-logistik/controllers/receipts"
	_receiptRepo "go-drop-logistik/drivers/databases/receipts"

	_manifestUsecase "go-drop-logistik/business/manifest"
	_manifestController "go-drop-logistik/controllers/manifest"
	_manifestRepo "go-drop-logistik/drivers/databases/manifest"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configApp := _config.GetConfig()
	mysqlConfigDB := _dbMysqlDriver.ConfigDB{
		DB_Username: configApp.Mysql.User,
		DB_Password: configApp.Mysql.Pass,
		DB_Host:     configApp.Mysql.Host,
		DB_Port:     configApp.Mysql.Port,
		DB_Database: configApp.Mysql.Name,
	}
	fmt.Println("User :", configApp.Mysql.User)
	fmt.Println("Pass :", configApp.Mysql.Pass)
	fmt.Println("Host :", configApp.Mysql.Host)
	fmt.Println("Port :", configApp.Mysql.Port)
	fmt.Println("Name :", configApp.Mysql.Name)
	// mongoConfigDB := _dbMongoDriver.ConfigDB{
	// 	DB_Username: configApp.MONGO_DB_USER,
	// 	DB_Password: configApp.MONGO_DB_PASS,
	// 	DB_Host:     configApp.MONGO_DB_HOST,
	// 	DB_Port:     configApp.MONGO_DB_PORT,
	// 	DB_Database: configApp.MONGO_DB_NAME,
	// }

	mysql_db := mysqlConfigDB.InitialMysqlDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       configApp.JWT.Secret,
		ExpiresDuration: configApp.JWT.Expired,
	}

	timeoutContext := time.Duration(configApp.JWT.Expired) * time.Second

	e := echo.New()

	logger := logging.NewLogger()

	middlewareLog := _middleware.NewMiddleware(logger)

	//! REPO
	userRepo := _userRepo.NewMySQLUserRepository(mysql_db)
	phoneAgentRepo := _phoneAgentRepo.NewMySQLPhoneAgentRepository(mysql_db)
	phoneRepo := _phoneRepo.NewMySQLPhoneRepository(mysql_db)
	agentRepo := _agentRepo.NewMySQLAgentRepository(mysql_db)
	adminRepo := _adminRepo.NewMySQLAdminRepository(mysql_db)
	receiptRepo := _receiptRepo.NewMySQLReceiptRepository(mysql_db)
	manifestReceiptRepo := _manifestreceiptRepo.NewMySQLManifestReceiptRepository(mysql_db)
	manifestRepo := _manifestRepo.NewMySQLManifestRepository(mysql_db)

	//! USECASE
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext, logger)
	phoneUsecase := _phoneUsecase.NewPhoneUsecase(phoneRepo, phoneAgentRepo, &configJWT, timeoutContext)
	agentUsecase := _agentUsecase.NewAgentUsecase(agentRepo, phoneAgentRepo, phoneUsecase, &configJWT, timeoutContext, logger)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo, &configJWT, timeoutContext, logger)
	receiptUsecase := _receiptUsecase.NewReceiptUsecase(receiptRepo, &configJWT, timeoutContext, logger)
	manifestReceiptUsecase := _manifestreceiptUsecase.NewManifestReceiptUsecase(manifestReceiptRepo, receiptRepo, &configJWT, timeoutContext)
	manifestUsecase := _manifestUsecase.NewManifestUsecase(manifestRepo, &configJWT, timeoutContext)

	//! CONTROLLER
	userCtrl := _userController.NewUserController(userUsecase)
	phoneCtrl := _phoneController.NewPhonesController(phoneUsecase)
	agentCtrl := _agentController.NewAgentController(agentUsecase)
	adminCtrl := _adminController.NewAdminController(adminUsecase, agentUsecase)
	receiptCtrl := _receiptController.NewReceiptController(receiptUsecase, manifestReceiptUsecase)
	manifestCtrl := _manifestController.NewManifestController(manifestUsecase, manifestReceiptUsecase)

	routesInit := _routes.ControllerList{
		MiddlewareLog:      middlewareLog,
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		AgentController:    *agentCtrl,
		AdminController:    *adminCtrl,
		ReceiptController:  *receiptCtrl,
		PhoneController:    *phoneCtrl,
		ManifestController: *manifestCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
