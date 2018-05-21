package cmd

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/database/mysql"
	"fmt"
)

const (
	//when we add a new migration this constant need to be updated
	step = 1
	//if we change the folder in which we keep our migration
	//files we need to update this
	sourceUrl = "file://migrations/"
)

func MigrateUp(db *sql.DB, databaseName string) {
	fmt.Print("Running migrations")
	migrate := setUpForMigration(db, databaseName)
	migrate.Steps(step)
}
func setUpForMigration(db *sql.DB, databaseName string)(*migrate.Migrate){
	migrationConfig := &mysql.Config{}
	driver, _ := mysql.WithInstance(db, migrationConfig)
	migrate, err := migrate.NewWithDatabaseInstance(
		sourceUrl,
		databaseName,
		driver,
	)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	return migrate
}
func MigrateDown(db *sql.DB, databaseName string){
	fmt.Print("Undoing  migrations")
	migrate := setUpForMigration(db, databaseName)
	migrate.Down()
}
