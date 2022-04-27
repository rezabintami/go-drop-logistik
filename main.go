package main

import (
	"log"
	"os"

	_config "go-drop-logistik/app/config"
	_middleware "go-drop-logistik/app/middleware"
	_plugins "go-drop-logistik/app/plugins"
	_dbPostgresDriver "go-drop-logistik/drivers/postgres"

	echo "github.com/labstack/echo/v4"

	"go-drop-logistik/cmd/seeder"
	"go-drop-logistik/helper/logging"
	"go-drop-logistik/helper/validation"
)

func main() {
	configApp := _config.GetConfig()

	log.Println("Starting application version :", configApp.App.Version)
	log.Println("Environment :", configApp.App.Env)
	log.Println("Server is running on port : " + configApp.Server.Address)

	log.Println("User :", configApp.Postgres.User)
	log.Println("Host :", configApp.Postgres.Host)
	log.Println("Port :", configApp.Postgres.Port)
	log.Println("Name :", configApp.Postgres.Name)

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	logger.LogServer("Server is running").Info("Server started at port ", port)
	log.Fatal(e.Start(":" + port))

}
