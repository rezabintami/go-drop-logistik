package seeder

import (
	_config "go-drop-logistik/app/config"
	_dbMysqlDriver "go-drop-logistik/drivers/mysql"
	"os"
	"sort"

	"fmt"
	"go-drop-logistik/helper/encrypt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
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

var (
	Conn *gorm.DB
)

func Seeder() {
	app := &cli.App{
		Name:  "seeder",
		Usage: "seeder",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "seeder",
				Aliases: []string{"s"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("seeder") {
				log.Println("Init seeder")

				configApp := _config.GetConfig()
				mysqlDb := _dbMysqlDriver.ConfigDB{
					DB_Username: configApp.Mysql.User,
					DB_Password: configApp.Mysql.Pass,
					DB_Host:     configApp.Mysql.Host,
					DB_Port:     configApp.Mysql.Port,
					DB_Database: configApp.Mysql.Name,
					Env:         configApp.App.Env,
				}
				fmt.Println("User :", configApp.Mysql.User)
				fmt.Println("Host :", configApp.Mysql.Host)
				fmt.Println("Port :", configApp.Mysql.Port)
				fmt.Println("Name :", configApp.Mysql.Name)

				Conn = mysqlDb.InitialMysqlDB()

				// Admins
				admin := Admins{}
				result := Conn.Where("name = ? ", "Admin").First(&admin)
				if result.Error == nil {
					fmt.Errorf("[error] failed to execute admin seeder query because data already exist")
				}

				admin.Name = "Admin"
				admin.Password, _ = encrypt.Hash("admin")
				admin.Email = "admin@gmail.com"
				admin.Roles = "ADMIN"

				result = Conn.Create(&admin)

				if result.Error != nil {
					fmt.Errorf("[error] failed to execute admin seeder query %s", result.Error)
				}

				fmt.Errorf("[success] success to execute seeder query")
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
