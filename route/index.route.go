package route

import (
	"go-fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controller.UserHandlerGetAll)
	app.Get("/users/:id", controller.UserHandlerGetById)
	app.Post("/users", controller.UserHandlerCreate)
	app.Put("/users/:id", controller.UserhandlerUpdate)
	app.Put("/users/:id/update-email", controller.UserhandlerUpdateEmail)
	app.Delete("/users/:id", controller.UserHandlerDelete)
}
