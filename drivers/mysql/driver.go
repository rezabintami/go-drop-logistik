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
	"go-drop-logistik/drivers/databases/trucks"
	"go-drop-logistik/drivers/databases/users"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
		&trucks.Trucks{},
		&drivers.Drivers{},
	)

	return db
}
