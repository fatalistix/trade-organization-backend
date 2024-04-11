package main

import (
	pgConnect "github.com/fatalistix/trade-organization-backend/internal/storage/connection/postgres"
	pgMigrate "github.com/fatalistix/trade-organization-backend/internal/storage/migration/postgres"
	"log"
	"os"
)

func main() {
	// init config

	// setup logger

	// setup database

	// setup router

	// run app

	migrationFiles := os.DirFS("db")

	db, err := pgConnect.NewDB(
		"localhost",
		5434,
		"trade-organization-owner",
		"trade-organization-owner",
		"trade-organization",
		pgConnect.SSLDisable,
	)
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := pgMigrate.NewMigrator(migrationFiles, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.ApplyMigrations(db, "trade-organization")
	if err != nil {
		log.Fatal(err)
	}
}
