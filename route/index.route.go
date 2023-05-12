package route

import (
	"go-fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controller.UserHandlerGetAll)
	app.Get("/users/:id", controller.UserHandlerGetById)
	app.Post("/users", controller.UserHandlerCreate)
	app.Patch("/users/:id", controller.UserhandlerUpdate)
}
