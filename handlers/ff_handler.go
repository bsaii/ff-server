package handlers

import (
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func BorrowedAmt(c *fiber.Ctx) error {
	account := c.Params("account")

	if account == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	ff := contract.FFInstance
	accountAddr := common.HexToAddress(account)

	value, err := ff.BorrowedAmount(nil, accountAddr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve amount borrowed for account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"borrowed": value.String(),
	})
}

func Earned(c *fiber.Ctx) error {
	account := c.Params("account")

	if account == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	ff := contract.FFInstance
	accountAddr := common.HexToAddress(account)

	value, err := ff.Earned(nil, accountAddr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve amount earned for account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"earned": value.String(),
	})
}

func InterestRate(c *fiber.Ctx) error {
	ff := contract.FFInstance

	value, err := ff.InterestRate(nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve the interest rate amount",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rate": value.String(),
	})
}

func RewardRate(c *fiber.Ctx) error {
	ff := contract.FFInstance

	value, err := ff.RewardRate(nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve the reward rate amount",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rate": value.String(),
	})
}

func User(c *fiber.Ctx) error {
	account := c.Params("account")

	if account == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	ff := contract.FFInstance
	accountAddr := common.HexToAddress(account)

	user, err := ff.Users(nil, accountAddr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve the user for an account",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"borrowingBalance":  user.BorrowingBalance.String(),
		"collateralBalance": user.CollateralBalance.String(),
		"lastUpdated":       user.LastUpdated.String(),
		"lendingBalance":    user.LendingBalance.String(),
		"rewards":           user.Rewards.String(),
		"stakingBalance":    user.StakingBalance.String(),
	})
}

func AllBorrowed(c *fiber.Ctx) error {
	borrows := []models.Borrow{}

	if err := database.DB.Find(&borrows).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all borrowed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all borrowed.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(borrows)
}

func AllLent(c *fiber.Ctx) error {
	lents := []models.Lent{}

	if err := database.DB.Find(&lents).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all lent: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all lent.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(lents)
}

func AllRepaid(c *fiber.Ctx) error {
	repays := []models.Repay{}

	if err := database.DB.Find(&repays).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all repaid: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all repaid.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(repays)
}

func AllStaked(c *fiber.Ctx) error {
	stakes := []models.Stake{}

	if err := database.DB.Find(&stakes).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all staked: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all staked.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(stakes)
}

func AllWithdrawn(c *fiber.Ctx) error {
	withdraws := []models.Withdraw{}

	if err := database.DB.Find(&withdraws).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all withdrawn: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all withdrawn.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(withdraws)
}
