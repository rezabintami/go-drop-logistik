package mysql_driver

import (
	"fmt"
	"go-drop-logistik/drivers/databases/admins"
	"go-drop-logistik/drivers/databases/agents"
	"go-drop-logistik/drivers/databases/drivers"
	"go-drop-logistik/drivers/databases/manifest"
	"go-drop-logistik/drivers/databases/manifestreceipt"
	"go-drop-logistik/drivers/databases/phoneagent"
	"go-drop-logistik/drivers/databases/phones"
	"go-drop-logistik/drivers/databases/receipts"
	"go-drop-logistik/drivers/databases/trackmanifest"
	"go-drop-logistik/drivers/databases/tracks"
	"go-drop-logistik/drivers/databases/trucks"
	"go-drop-logistik/drivers/databases/users"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
	DB_SSL      string
	Env         string
}

var (
	db  *gorm.DB
	err error
)

func (config *ConfigDB) InitialPostgresDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s host=%s dbname=%s sslmode=%s password=%s port=%s",
		config.DB_Username,
		config.DB_Host,
		config.DB_Database,
		config.DB_SSL,
		config.DB_Password,
		config.DB_Port)

	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&users.Users{},
		&agents.Agents{},
		&admins.Admins{},
		&receipts.Receipts{},
		&phones.Phones{},
		&phoneagent.PhoneAgent{},
		&manifest.Manifest{},
		&manifestreceipt.ManifestReceipt{},
		&tracks.Tracks{},
		&trackmanifest.TrackManifest{},
		&trucks.Trucks{},
		&drivers.Drivers{},
	)

	return db
}
