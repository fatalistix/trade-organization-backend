package main

import (
	pgconnect "github.com/fatalistix/trade-organization-backend/internal/database/connection/postgres"
	pgmigrate "github.com/fatalistix/trade-organization-backend/internal/database/migration/postgres"
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

	db, err := pgconnect.NewDB(
		"localhost",
		5434,
		"trade-organization-owner",
		"trade-organization-owner",
		"trade-organization",
		pgconnect.SSLDisable,
	)
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := pgmigrate.NewMigrator(migrationFiles, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.ApplyMigrations(db, "trade-organization")
	if err != nil {
		log.Fatal(err)
	}
}
