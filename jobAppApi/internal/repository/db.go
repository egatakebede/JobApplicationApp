package repository

import (
	"database/sql"

	_ "modernc.org/sqlite" // Must be imported for side effects
)

// initDB initializes the SQLite database and ensures required tables exist.
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./jobportal.db")
	if err != nil {
		return nil, err
	}

	// Create tables
	if err := createTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

// createTable creates the users table if it doesn't already exist.
func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			is_admin BOOLEAN DEFAULT 0,
			profile_picture TEXT
		)
	`)
	return err
}
