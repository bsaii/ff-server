package handlers

import (
	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type BorrowedAmountResponse struct {
	Borrowed string `json:"borrowed"`
}

type EarnedResponse struct {
	Earned string `json:"earned"`
}

type RateResponse struct {
	Rate string `json:"rate"`
}

type UserResponse struct {
	BorrowedBalance   string `json:"borrowedBalance"`
	CollateralBalance string `json:"collateralBalance"`
	LastUpdated       string `json:"lastUpdated"`
	LendingBalance    string `json:"lendingBalance"`
	Rewards           string `json:"rewards"`
	StakingBalance    string `json:"stakingBalance"`
}

// @Summary		Retrieve borrowed amount
// @Tags			contract
// @Description	Retrieves the amount borrowed by an account from the blockchain
// @Param			account	path		string			true	"Account address"
// @Success		200		{object}	BorrowedAmountResponse	"Successful retrieval"
// @Failure		404		{object}	ErrorResponse	"Account not found"
// @Failure		500		{object}	ErrorResponse	"Internal server error"
// @Router			/api/contract/borrowed/{account} [get]
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

	return c.Status(fiber.StatusOK).JSON(BorrowedAmountResponse{
		Borrowed: value.String(),
	})
}

// @Summary Retrieve earned amount
// @Tags contract
// @Description Retrieves the amount earned by an account from the blockchain
// @Param account path string true "Account address"
// @Success 200 {object} EarnedResponse "Successful retrieval"
// @Failure 404 {object} ErrorResponse "Account not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/earned/{account} [get]
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

	return c.Status(fiber.StatusOK).JSON(EarnedResponse{
		Earned: value.String(),
	})
}

// @Summary Retrieve interest rate
// @Tags contract
// @Description Retrieves the general interest rate from the blockchain
// @Success 200 {object} RateResponse "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/rate/interest [get]
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

// @Summary Retrieve reward rate
// @Tags contract
// @Description Retrieves the general reward rate from the blockchain
// @Success 200 {object} RateResponse "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/rate/reward [get]
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

// @Summary Retrieve user information
// @Tags contract
// @Description Retrieves user information from the blockchain using the address
// @Param account path string true "User address"
// @Success 200 {object} UserResponse "Successful retrieval"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/user/{account} [get]
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

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		BorrowedBalance:   user.BorrowingBalance.String(),
		CollateralBalance: user.CollateralBalance.String(),
		LastUpdated:       user.LastUpdated.String(),
		LendingBalance:    user.LendingBalance.String(),
		Rewards:           user.Rewards.String(),
		StakingBalance:    user.StakingBalance.String(),
	})
}

// @Summary Retrieve total borrowed amount
// @Tags contract
// @Description Retrieves the total borrowed amount from the blockchain
// @Success 200 {array} models.Borrow "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/borrowed [get]
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

// @Summary Retrieve total lent amount
// @Tags contract
// @Description Retrieves the total lent amount from the blockchain
// @Success 200 {object} []models.Lent "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/lent [get]
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

// @Summary Retrieve total repaid amount
// @Tags contract
// @Description Retrieves the total repaid amount from the blockchain
// @Success 200 {array} models.Repay "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/repaid [get]
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

// @Summary Retrieve total reward amount
// @Tags contract
// @Description Retrieves the total reward amount from the blockchain
// @Success 200 {array} models.Reward "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/reward [get]
func AllReward(c *fiber.Ctx) error {
	rewards := []models.Reward{}

	if err := database.DB.Find(&rewards).Error; err != nil {
		fiberlog.Errorf("Failed to retrieve all rewards: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve all rewards.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(rewards)
}

// @Summary Retrieve total staked amount
// @Tags contract
// @Description Retrieves the total staked amount from the blockchain
// @Success 200 {array} models.Stake "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/staked [get]
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

// @Summary Retrieve total withdrawn amount
// @Tags contract
// @Description Retrieves the total withdrawn amount from the blockchain
// @Success 200 {array} models.Withdraw "Successful retrieval"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/contract/withdrawn [get]
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
