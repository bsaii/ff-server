package main

import (
	"log"

	"github.com/bsaii/ff-server/config"
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/handlers"
	"github.com/bsaii/ff-server/listener"
	"github.com/ethereum/go-ethereum/ethclient"
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
	ec, err := ethclient.Dial(cfg.EthNodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	contract.InitContract(ec, cfg)

	// Events
	listener.Approval()
	listener.Transfer()

	app := fiber.New(fiber.Config{
		AppName: "Finance Forget Server",
	})
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Finance Forge")
	})

	token := app.Group("/api/token")
	token.Get("/allowance", handlers.Allowance)
	token.Get("/balance/:account", handlers.BalanceOf)
	token.Get("/approvals", handlers.Approvals)
	token.Get("/transfers", handlers.Transfers)

	// contract := app.Group("/api/contract")
	// contract.Get("/stake/:eth_address")

	log.Printf("Server running on port: %v", cfg.Port)
	app.Listen(cfg.Port)
}
