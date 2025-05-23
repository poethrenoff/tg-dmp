package app

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDatabase() *sqlx.DB {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current directory: %s", err.Error())
	}

	for {
		migrationsPath := filepath.Join(currentDir, "migrations")

		if _, err := os.Stat(migrationsPath); err == nil {
			migrator, err := migrate.New("file://"+migrationsPath, os.Getenv("POSTGRES_DSN"))
			if err != nil {
				log.Fatalf("failed to apply migrations: %s", err.Error())
			}
			migrator.Up()
			break
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}

	return db
}

func CloseDatabase(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}
}
