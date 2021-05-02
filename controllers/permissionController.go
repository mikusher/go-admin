package controllers

import (
	"go-admin/database"
	"go-admin/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermission(c *fiber.Ctx) error {
	var permission []models.Permission

	database.DB.Find(&permission)
	return c.JSON(permission)
}
