package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial database
	database.DatabaseInit()
	// RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// initial route
	route.RouteInit(app)

	app.Listen(":8080")
}
