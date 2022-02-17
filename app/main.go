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

	_superuserUsecase "go-drop-logistik/business/superusers"
	_superuserController "go-drop-logistik/controllers/superusers"
	_superuserRepo "go-drop-logistik/drivers/databases/superusers"

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

	userRepo := _userRepo.NewMySQLUserRepository(mysql_db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext, logger)
	userCtrl := _userController.NewUserController(userUsecase)

	agentRepo := _agentRepo.NewMySQLAgentRepository(mysql_db)
	agentUsecase := _agentUsecase.NewAgentUsecase(agentRepo, &configJWT, timeoutContext, logger)
	agentCtrl := _agentController.NewAgentController(agentUsecase)

	superuserRepo := _superuserRepo.NewMySQLSuperusersRepository(mysql_db)
	superuserUsecase := _superuserUsecase.NewSuperuserUsecase(superuserRepo, &configJWT, timeoutContext, logger)
	superuserCtrl := _superuserController.NewSuperuserController(superuserUsecase, agentUsecase)

	routesInit := _routes.ControllerList{
		MiddlewareLog:       middlewareLog,
		JWTMiddleware:       configJWT.Init(),
		UserController:      *userCtrl,
		AgentController:     *agentCtrl,
		SuperuserController: *superuserCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
