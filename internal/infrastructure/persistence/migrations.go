package persistence

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) error {
	log.Println("Running database migrations...")

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}