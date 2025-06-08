package lead

import (
	db "github.com/OceanWhisperer/Simple-CRM/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func Getleads(c *fiber.Ctx) error{
	database := db.DbConn
	var leads []Lead
	database.Find(&leads)
	return c.JSON(leads)
}

func Getlead(c *fiber.Ctx) error{
	id := c.Params("id")
	database := db.DbConn
	var lead Lead
	database.Find(&lead, id)
	return c.JSON(lead)

}

func Newlead(c *fiber.Ctx) error{
	database := db.DbConn
	var lead Lead
	// parses body from context and sends error if issues
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	database.Create(&lead)
	return c.JSON(lead)

}

func Deletelead(c *fiber.Ctx) error {
	id := c.Params("id")
	database := db.DbConn
	var lead Lead
	database.First(&lead, id)
	if lead.Name == "" {
		return c.Status(404).SendString("No lead found")
	}
	database.Delete(&lead)
	return c.SendString("Lead deleted")
}
