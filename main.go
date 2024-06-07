package main

import (
	"log"

	"github.com/bsaii/ff-server/config"
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	_ "github.com/bsaii/ff-server/docs"
	"github.com/bsaii/ff-server/handlers"
	"github.com/bsaii/ff-server/listener"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Finance Forge Server
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
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
	listener.Borrowed()
	listener.Lent()
	listener.Repaid()
	listener.RewardRepaid()
	listener.Staked()
	listener.Withdrawn()

	app := fiber.New(fiber.Config{
		AppName: "Finance Forget Server",
	})
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Finance Forge")
	})

	token := app.Group("/api/token")
	token.Get("/allowance", handlers.Allowance)
	token.Get("/balance/:account", handlers.BalanceOf)
	token.Get("/approvals", handlers.Approvals)
	token.Get("/transfers", handlers.Transfers)

	contract := app.Group("/api/contract")
	contract.Get("/borrowed/:account", handlers.BorrowedAmt)
	contract.Get("/earned/:account", handlers.Earned)
	contract.Get("/rate/interest", handlers.InterestRate)
	contract.Get("/rate/reward", handlers.RewardRate)
	contract.Get("/user/:account", handlers.User)
	contract.Get("/borrowed", handlers.AllBorrowed)
	contract.Get("/lent", handlers.AllLent)
	contract.Get("/repaid", handlers.AllRepaid)
	contract.Get("/staked", handlers.AllStaked)
	contract.Get("/withdrawn", handlers.AllWithdrawn)

	log.Printf("Server running on port: %v", cfg.Port)
	app.Listen(cfg.Port)
}
