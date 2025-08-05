package main

import (
	"cars/pranay/github.com/config"
	"cars/pranay/github.com/handlers"
	"cars/pranay/github.com/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.ConnectDB()

	app := fiber.New()
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"pranay":  "password",
			"admin":   "password",
			"manager": "password",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Put("/cars", handlers.UpdateCar)
	app.Use(middleware.SecurityHeaders)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting server", err)
		panic(err)
	}
}
