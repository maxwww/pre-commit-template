package main

import (
	"log"
	_ "precommit/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus" //nolint:depguard
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	logrus.Println("hello")
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/ping", pingHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Println(app.Listen(":8080"))
}

// pingHandler godoc
//
//	@Summary	ping handler
//	@Tags		Ping
//	@Success	200
//	@Router		/ping [get]
func pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
