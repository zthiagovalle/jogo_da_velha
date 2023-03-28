package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zthiagovalle/jogo_da_velha/cmd/api/handlers"
)

type HashGame struct {
	Matriz [3][3]string `json:"matriz"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/", handlers.RunHashGame)
	app.Listen(":3000")
}
