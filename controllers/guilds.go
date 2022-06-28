package controllers

import (
	"blacklist-service/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	GetGuildReportById = "SELECT id, total_reports, first_report, last_report FROM guilds_reports WHERE id = $1"
)

func GetGuildByID(c *fiber.Ctx) error {
	guildID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var id int64
	var totalReports int
	var firstReport time.Time
	var lastReport time.Time

	if err := database.QueryRow(GetGuildReportById, guildID).Scan(&id, &totalReports, &firstReport, &lastReport); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"id":           id,
		"totalReports": totalReports,
		"firstReport":  firstReport,
		"lastReport":   lastReport,
	})
}

func ReportGuildByID(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
