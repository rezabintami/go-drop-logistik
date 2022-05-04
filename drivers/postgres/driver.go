package mysql_driver

import (
	"fmt"

	"log"

	_config "go-drop-logistik/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func GetConnection() string {
	return fmt.Sprintf("user=%s host=%s dbname=%s sslmode=%s password=%s port=%s",
		_config.GetConfiguration("postgres.user"),
		_config.GetConfiguration("postgres.host"),
		_config.GetConfiguration("postgres.name"),
		_config.GetConfiguration("postgres.ssl"),
		_config.GetConfiguration("postgres.pass"),
		_config.GetConfiguration("postgres.port"))
}

func InitialPostgresDB() *gorm.DB {
	db, err = gorm.Open("postgres", GetConnection())
	if err != nil {
		log.Fatal(err)
	}
	if _config.GetConfiguration("app") == "DEV" {
		db.LogMode(true)
	}

	return db
}
