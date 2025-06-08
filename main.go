package main

import (
	"github.com/OceanWhisperer/Simple-CRM/database.db"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get(getleads)
	app.Get(getlead)
	app.Post(newlead)
	app.Delete(deletelead)
}

func main() {

	app := fiber.New()
	app.Listen(":3000")

}
