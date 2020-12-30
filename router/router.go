package router

import (
	"race_condition/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/transferMoney", handler.TransferMoney)

}
