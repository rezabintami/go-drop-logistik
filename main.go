package main

import (
	"log"
	"os"

	_config "go-drop-logistik/app/config"
	_middleware "go-drop-logistik/app/middleware"
	_plugins "go-drop-logistik/app/plugins"
	_dbPostgresDriver "go-drop-logistik/drivers/postgres"
	"go-drop-logistik/helpers"

	echo "github.com/labstack/echo/v4"

	"go-drop-logistik/cmd/seeder"
)

func main() {
	log.Println("Starting application version :", _config.GetConfiguration("app.version"))
	log.Println("Environment :", _config.GetConfiguration("app.env"))
	log.Println("Server is running on port : " + _config.GetConfiguration("server.port"))

	log.Println("User :", _config.GetConfiguration("postgres.user"))
	log.Println("Host :", _config.GetConfiguration("postgres.host"))
	log.Println("Port :", _config.GetConfiguration("postgres.port"))
	log.Println("Name :", _config.GetConfiguration("postgres.name"))

	postgres_db := _dbPostgresDriver.InitialPostgresDB()

	// Init Seeding
	err := seeder.Seeder(postgres_db)
	log.Println(err)

	// Init Validation
	helpers.InitValidation()

	e := echo.New()
	logger := helpers.NewLogger()
	middlewareLog := _middleware.NewMiddleware(logger)

	plugins := _plugins.ConfigurationPlugins{
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
