package main

import (
	healthCheck "server/modules/HealthCheck"
	"server/modules/dbManager"

	"server/modules/migrations"
	"server/modules/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	dbConn := dbManager.GetDb()

	migrations.Migrate(dbConn)

	sqlDB, err := dbConn.DB()

	if err != nil {
		println("DB connection error: ", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${latency} ${method} ${path}\n",
	}))
	api := app.Group("/api")

	healthCheck.Router(api)
	user.Router(api)

	app.Listen(":3000")
}
