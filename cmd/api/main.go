package main

import (
	"log"

	"root_academy_go/internal/platform/db"
)

func main() {
	if err := db.RunMigrations(db.MigrationURLFromEnv(), "migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	conn, err := db.NewPostgresConnection(db.DSNFromEnv())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("failed to get underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()

	log.Println("connected to database, migrations applied")
}
