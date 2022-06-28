package controllers

import "github.com/gofiber/fiber/v2"

func GetGuildByID(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func ReportGuildByID(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
