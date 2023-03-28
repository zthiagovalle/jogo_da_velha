package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zthiagovalle/jogo_da_velha/cmd/api/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/", handlers.RunHashGame)
	app.Listen(":3000")
}
