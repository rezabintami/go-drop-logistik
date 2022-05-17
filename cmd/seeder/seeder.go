package seeder

import (
	"errors"
	"fmt"
	"go-drop-logistik/helpers"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Admins struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Password  string
	Email     string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Seeder(Conn *gorm.DB) error {
	log.Println("Initializing Seeder...")

	// Admins
	admin := Admins{}
	result := Conn.Where("name = ? ", "Admin").First(&admin)
	if result.Error == nil {
		return errors.New("[error] failed to execute admin seeder query because data already exist")
	}

	admin.Name = "Admin"
	admin.Password, _ = helpers.Hash("admin")
	admin.Email = "admin@gmail.com"
	admin.Roles = "ADMIN"

	result = Conn.Create(&admin)

	if result.Error != nil {
		return fmt.Errorf("[error] failed to execute admin seeder query %s", result.Error)
	}

	return fmt.Errorf("[success] success to execute seeder query")
}
