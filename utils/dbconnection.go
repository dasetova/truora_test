package utils

import (
	"errors"
	"log"

	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

var (
	ErrDBConnection = errors.New("failed to connect database")
	errDBMigration  = errors.New("failed to migrate database")
)

//ConnectToDB func that creates a DB connection
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=26257 user=truora_test dbname=recipes_db sslmode=disable")
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
