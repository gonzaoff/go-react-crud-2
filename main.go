package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Welcome Go!")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	
	app.Use(cors.New())

	app.Static("/","./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error{
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el backend",
		})
	})


	fmt.Println("Servidor iniciado en puerto :3000")
	app.Listen(":3000")
}