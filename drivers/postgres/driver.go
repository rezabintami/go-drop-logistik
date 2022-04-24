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
		_config.GetConfig().Postgres.User,
		_config.GetConfig().Postgres.Host,
		_config.GetConfig().Postgres.Name,
		_config.GetConfig().Postgres.SSL,
		_config.GetConfig().Postgres.Pass,
		_config.GetConfig().Postgres.Port)
}

func InitialPostgresDB() *gorm.DB {
	db, err = gorm.Open("postgres", GetConnection())

	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate(
	// 	&users.Users{},
	// 	&agents.Agents{},
	// 	&admins.Admins{},
	// 	&receipts.Receipts{},
	// 	&phones.Phones{},
	// 	&phoneagent.PhoneAgent{},
	// 	&manifest.Manifest{},
	// 	&manifestreceipt.ManifestReceipt{},
	// 	&tracks.Tracks{},
	// 	&trackmanifest.TrackManifest{},
	// 	&trucks.Trucks{},
	// 	&drivers.Drivers{},
	// )

	return db
}
