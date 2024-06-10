package handlers

import (
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

type AllowanceResponse struct {
	Allowance string `json:"allowance"`
}

type BalanceResponse struct {
	Balance string `json:"balance"`
}

// @Summary Retrieve remaining token allowance
// @Tags token
// @Description Returns the remaining number of tokens that the spender is allowed to spend on behalf of the owner
// @Param owner query string true "owner address"
// @Param spender query string true "spender address"
// @Success 200 {object} AllowanceResponse "Successful retrieval"
// @Failure 404 {object} ErrorResponse "Allowance not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/token/allowance [get]
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

	return c.Status(fiber.StatusOK).JSON(AllowanceResponse{
		Allowance: amt.String(),
	})
}

// @Summary Retrieve account balance
// @Tags token
// @Description Returns the value of tokens owned by the specified account
// @Param account path string true "Account address"
// @Success 200 {object} BalanceResponse "Successful retrieval"
// @Failure 404 {object} ErrorResponse "Account not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/token/balance/{account} [get]
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

	return c.Status(fiber.StatusOK).JSON(BalanceResponse{
		Balance: amt.String(),
	})
}

// @Summary Retrieve all approvals
// @Tags token
// @Description Returns all the approvals that have happened on the blockchain and captured by the server
// @Success 200 {array} models.Approval "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/token/approvals [get]
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

// @Summary Retrieve all transfers
// @Tags token
// @Description Returns all the transfers that have happened on the blockchain and captured by the server
// @Success 200 {array} models.Transfer "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/token/transfers [get]
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
