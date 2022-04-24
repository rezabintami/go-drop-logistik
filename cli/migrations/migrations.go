package main

import (
	"log"
	"os"
	"sort"

	_dbPostgresDriver "go-drop-logistik/drivers/postgres"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "migrations",
		Usage:                "migrations",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "go-drop-logistik:migrate",
				Aliases: []string{"m"},
				Usage:   "migrate",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "up",
						Aliases: []string{"u"},
					},
					&cli.BoolFlag{
						Name:    "down",
						Aliases: []string{"d"},
					},
				},
				Action: func(ctx *cli.Context) error {
					db, err := gorm.Open("postgres", _dbPostgresDriver.GetConnection())
					if err != nil {
						log.Fatalf("Error connection to main db %v \n", err)
					}

					defer db.Close()
					db.LogMode(true)

					driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
					if err != nil {
						log.Fatalf("could not start sql migration... %v", err)
					}
					m, err := migrate.NewWithDatabaseInstance(
						"file://drivers/postgres/files/migrations",
						"postgres", driver)
					if err != nil {
						log.Fatalf("migration failed... %v", err)
					}
					if ctx.Bool("up") {
						if err := m.Up(); err != nil {
							log.Fatalf("An error occurred while syncing the database.. %v", err)
						}
						log.Println("Database go-drop-logistik migrated")
					}
					if ctx.Bool("down") {
						if err := m.Down(); err != nil {
							log.Fatalf("An error occurred while syncing the database.. %v", err)
						}
						log.Println("Database go-drop-logistik down")
					}
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
