package main

import (
	"github.com/CesarDelgadoM/go-fiber-crm-basic/database"
	"github.com/CesarDelgadoM/go-fiber-crm-basic/models"
	"github.com/CesarDelgadoM/go-fiber-crm-basic/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{EnablePrintRoutes: true})
	database.InitDatabase()
	database.AutoMigrate(&models.Lead{})
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	leadServ := service.LeadServ{}
	app.Get("/api/v1/lead/getall", leadServ.GetLeads)
	app.Get("/api/v1/lead/get/:id", leadServ.GetLead)
	app.Post("/api/v1/lead/create", leadServ.NewLead)
	app.Delete("/api/v1/lead/delete/:id", leadServ.DeleteLead)
}
