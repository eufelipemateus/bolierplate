package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	config "github.com/eufelipemateus/bolierplate-go/configs"
	"github.com/eufelipemateus/bolierplate-go/database"
	"github.com/eufelipemateus/bolierplate-go/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		panic("Erro on load config.toml")
	}
	database.OpenConnection()

	if config.IsGenerateDB() {
		database.GenerateDB()
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(config.GetServerPort())
}
