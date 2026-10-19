package main

import (
	"log"

	"root_academy_go/internal/platform/db"
	"root_academy_go/internal/user/adapter/in/web"
	"root_academy_go/internal/user/adapter/out/persistance"
	"root_academy_go/internal/user/application/service"

	"github.com/gin-gonic/gin"
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

	userRepo := persistance.NewUserRepositoryAdapter(conn)
	userService := service.NewUserService(userRepo)
	registerController := web.NewRegisterController(userService)

	router := gin.Default()
	router.POST("/register", registerController.Handle)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
