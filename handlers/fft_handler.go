package handlers

import (
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func Allowance(c *fiber.Ctx) error {
	q := c.Queries()
	owner := q["owner"]
	spender := q["spender"]

	if owner == "" || spender == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Owner or Spender not found",
		})
	}

	fftoken := contract.FFTokenInstance
	ownerAddr := common.HexToAddress(owner)
	spenderAddr := common.HexToAddress(spender)

	amt, err := fftoken.Allowance(nil, ownerAddr, spenderAddr)
	if err != nil {
		fiberlog.Errorf("Failed to retrieve remaining number of tokens: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve remaining number of tokens",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"allowance": amt.String(),
	})
}

func BalanceOf(c *fiber.Ctx) error {
	account := c.Params("account")

	if account == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	fftoken := contract.FFTokenInstance
	accountAddr := common.HexToAddress(account)

	amt, err := fftoken.BalanceOf(nil, accountAddr)
	if err != nil {
		fiberlog.Errorf("Failed to retrieve the value of tokens owned by account: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve the value of tokens owned by account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"balance": amt.String(),
	})
}

func Approvals(c *fiber.Ctx) error {
	approvals := []models.Approval{}

	if err := database.DB.Find(&approvals).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all approvals: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all approvals.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(approvals)
}

func Transfers(c *fiber.Ctx) error {
	transfers := []models.Transfer{}

	if err := database.DB.Find(&transfers).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all transfers: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all transfers.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(transfers)
}
