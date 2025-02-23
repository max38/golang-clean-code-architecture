package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger"
)

// @title Golang Clean Code Architecture API
// @version 1.0
// @description This is a sample server for a Golang Clean Code Architecture project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// SetupSwagger sets up the Swagger handler
func SetupSwagger(app *fiber.App) {
    app.Get("/swagger/*", swagger.HandlerDefault)
}