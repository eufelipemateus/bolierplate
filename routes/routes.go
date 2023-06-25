package routes

import (
	"github.com/eufelipemateus/bolierplate-go/controllers"
	"github.com/gofiber/fiber/v2"
)

//Setup configura las rutas de la api
func Setup(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/eample", controllers.ExamplePage)
	
}