package service

import (
	"github.com/CesarDelgadoM/go-fiber-crm-basic/database"
	"github.com/CesarDelgadoM/go-fiber-crm-basic/models"
	"github.com/gofiber/fiber/v2"
)

type LeadServ struct{}

func (ls *LeadServ) GetLeads(ctx *fiber.Ctx) error {
	var leads []models.Lead
	err := database.DB.Debug().Model(&leads).Limit(100).Find(&leads).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	ctx.JSON(leads)
	return ctx.SendStatus(fiber.StatusOK)
}

func (ls *LeadServ) GetLead(ctx *fiber.Ctx) error {
	var lead models.Lead
	id := ctx.Params("id")
	err := database.DB.Debug().Model(&lead).Where(("ID = ?"), id).Take(&lead).Error
	if err != nil {
		return ctx.Status(404).SendString("Not found")
	}
	ctx.JSON(lead)
	return ctx.SendStatus(200)
}

func (ls *LeadServ) NewLead(ctx *fiber.Ctx) error {
	lead := new(models.Lead)
	err := ctx.BodyParser(lead)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).SendString("Unprocesable Body")
	}
	err = database.DB.Debug().Model(&lead).Create(&lead).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return ctx.Status(fiber.StatusCreated).SendString("Created Succesful")
}

func (ls *LeadServ) DeleteLead(ctx *fiber.Ctx) error {
	var lead models.Lead
	id := ctx.Params("id")
	err := database.DB.Debug().Model(&lead).Where("id = ?", id).Delete(&lead).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Not found")
	}
	return ctx.Status(fiber.StatusOK).SendString("Delete Succesful")
}
