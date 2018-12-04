package utils

import (
	"errors"
	"log"
	"os"

	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

var (
	ErrDBConnection = errors.New("failed to connect database")
	errDBMigration  = errors.New("failed to migrate database")
	errEnvDBUrl     = errors.New("Env variable found DATABASE_URL not found")
	errEnvDBType    = errors.New("Env variable found DATABASE_TYPE not found")
	errEnvApi       = errors.New("Env variable found API_PORT not found")
)

//Validates Environment Variables
func ValidateEnvVars() error {
	if len(os.Getenv("DATABASE_URL")) == 0 {
		return errEnvDBUrl
	}
	if len(os.Getenv("DATABASE_TYPE")) == 0 {
		return errEnvDBType
	}
	if len(os.Getenv("API_PORT")) == 0 {
		return errEnvApi
	}
	return nil
}

//ConnectToDB func that creates a DB connection
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_URL"))
	return db, err
}

//Execute DB Migrations
func MigrateDB() {
	db, err := ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrator, _ := gomigrate.NewMigratorWithLogger(db, gomigrate.Postgres{}, "./migrations", logrus.New())
	err = migrator.Migrate()

	if err != nil {
		migrator.Rollback()
		log.Fatal(err)
	}
}
