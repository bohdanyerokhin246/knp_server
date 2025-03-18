package app

import (
	"knp_server/internal/database/postgresql"
	"log"
)

func dbConnectAndMigrate() {
	// Connecting to Postgresql
	err := postgresql.Connect()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	// Postgresql migration
	err = postgresql.Migrate()
	if err != nil {
		log.Fatalf("Error during DB migration: %v", err)
	}
}
