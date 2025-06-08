package main

import (
	"fmt"
	"log"

	"github.com/OceanWhisperer/Simple-CRM/database"
	lead "github.com/OceanWhisperer/Simple-CRM/leaad"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1", lead.Getleads)
	app.Get("/api/v1/:id", lead.Getleads)
	app.Post("/api/v1", lead.Newlead)
	app.Delete("/api/v1/:id", lead.Deletelead)
}

func dbInit() {
     var err error
	db.DbConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	 if err != nil {
		log.Fatal("failed to connect to db");
	 }
	 fmt.Printf("Connection opened to Database successfully!")
	 
	 db.DbConn.AutoMigrate(&lead.Lead{})
	 fmt.Println("DataBase Migrated to type lead")	
}

func main() {
    dbInit()	
	app := fiber.New()
	app.Listen(":3000")
    sqldb , err := db.DbConn.DB();
	if err != nil {
		log.Fatal(err);
	}
	defer sqldb.Close()
}
