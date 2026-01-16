package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	m, err := migrate.New(
		"file://migrations",
		"sqlite://mydb.sqlite",
	)
	if err != nil {
		log.Fatal("Migration init error:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration error:", err)
	}

	log.Println("Migrations applied")
}
