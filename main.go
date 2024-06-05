package main

import (
	"log"

	"github.com/bsaii/ff-server/config"
	"github.com/bsaii/ff-server/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize database
	err := database.InitDatabase(cfg.DbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Connect to Ethereum node
	//   client, err := ethclient.Dial(cfg.EthNodeURL)
	//   if err != nil {
	// 	  log.Fatalf("Failed to connect to Ethereum node: %v", err)
	//   }

	app := fiber.New(fiber.Config{
		AppName: "Finance Forget Server",
	})
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}
