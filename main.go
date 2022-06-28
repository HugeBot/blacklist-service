package main

import (
	"blacklist-service/config"
	"blacklist-service/controllers"
	"blacklist-service/database"
	"blacklist-service/middlewares"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var (
	cfg *config.Config
	app *fiber.App
)

func init() {
	var path string
	flag.StringVar(&path, "config", "config.yml", "Path to config file")
	flag.Parse()

	if newCfg, err := config.Init(path); err != nil {
		log.Fatal(err)
	} else {
		cfg = newCfg
	}

	if err := database.Init(cfg.Database); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app = fiber.New()

	app.Get("/guild/:id", middlewares.ClientAuthorization, controllers.GetGuildByID)

	app.Post("/report/guild/:id", middlewares.ClientAuthorization, controllers.ReportGuildByID)

	app.Get("/user/:id", middlewares.ClientAuthorization, controllers.GetGuildByID)

	app.Post("/report/user/:id", middlewares.ClientAuthorization, controllers.ReportGuildByID)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)))
}
