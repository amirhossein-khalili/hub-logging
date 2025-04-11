package db

import (
	"database/sql"
	"fmt"
	"log"

	"hub_logging/config"

	_ "github.com/lib/pq"
)

// EnsureDatabase connects to the default Postgres database and creates the target database if it does not exist.
func EnsureDatabase(cfg config.AppConfig) error {
	// Connect to the default "postgres" database.
	defaultDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort)
	sqlDB, err := sql.Open("postgres", defaultDSN)
	if err != nil {
		return fmt.Errorf("failed to connect to default database: %w", err)
	}
	defer sqlDB.Close()

	// Check if the target database exists.
	var exists int
	checkQuery := "SELECT 1 FROM pg_database WHERE datname = $1"
	err = sqlDB.QueryRow(checkQuery, cfg.DBName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	// Create the database if it does not exist.
	if exists != 1 {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", cfg.DBName)
		if _, err = sqlDB.Exec(createQuery); err != nil {
			return fmt.Errorf("failed to create database %s: %w", cfg.DBName, err)
		}
		log.Printf("Database %s created.\n", cfg.DBName)
	} else {
		log.Printf("Database %s already exists.\n", cfg.DBName)
	}
	return nil
}

// CreateExtension connects to the given database (via the provided *sql.DB) and creates the extension.
func CreateExtension(sqlDB *sql.DB) error {
	query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	if _, err := sqlDB.Exec(query); err != nil {
		return fmt.Errorf("failed to create extension: %w", err)
	}
	return nil
}
