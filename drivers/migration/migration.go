package migration

import (
	"go-drop-logistik/drivers/databases/admins"
	"go-drop-logistik/helper/encrypt"
)

func AdminMigration() interface{} {
	pass, _ := encrypt.Hash("123123")

	return &admins.Admins{
		Name: "Admin",
		Email: "admin@gmail.com",
		Password: pass,
		Roles: "ADMIN",		
	}
}