package listener

import (
	"fmt"
	"log"

	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func Borrowed() {
	log.Println("WatchBorrowed...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFBorrowed)

	sub, err := ff.WatchBorrowed(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to borrowed events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Borrowed Event subscription failed: %v", err)
			case borrowedLog := <-ch:
				fmt.Print(borrowedLog)

				amount := borrowedLog.Amount.String()
				collateral := borrowedLog.Collateral.String()
				account := borrowedLog.User.Hex()

				fmt.Printf("New borrowed event with user: [%s], amount: [%s], collateral: [%s].", account, amount, collateral)

				borrowed := models.Borrow{
					Address:    account,
					Amount:     amount,
					Collateral: collateral,
				}

				database.DB.Create(&borrowed)
				fiberlog.Infof("New borrowed with ID: [%d].", borrowed.ID)

			}
		}
	}()
}

func Lent() {
	log.Println("WatchLent...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFLent)

	sub, err := ff.WatchLent(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to lent events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Lent Event subscription failed: %v", err)
			case lentLog := <-ch:
				fmt.Print(lentLog)

				amount := lentLog.Amount.String()
				account := lentLog.User.String()

				fmt.Printf("New lent event with user: [%s], amount: [%s].", account, amount)

				lent := models.Lent{
					Address: account,
					Amount:  amount,
				}

				database.DB.Create(&lent)
				fiberlog.Infof("New lent with ID: [%d].", lent.ID)
			}
		}
	}()
}

func Repaid() {
	log.Println("WatchRepaid...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFRepaid)

	sub, err := ff.WatchRepaid(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to repaid events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Repaid Event subscription failed: %v", err)
			case repaidLog := <-ch:
				fmt.Print(repaidLog)

				amount := repaidLog.Amount.String()
				account := repaidLog.User.String()

				fmt.Printf("New repaid event with user: [%s], amount: [%s].", account, amount)

				repay := models.Repay{
					Address: account,
					Amount:  amount,
				}

				database.DB.Create(&repay)
				fiberlog.Infof("New repay with ID: [%d].", repay.ID)
			}
		}
	}()
}

func RewardRepaid() {
	log.Println("WatchRewardRepaid...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFRewardPaid)

	sub, err := ff.WatchRewardPaid(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to rewardRepaid events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("RewardRepaid Event subscription failed: %v", err)
			case rewardRepaidLog := <-ch:
				fmt.Print(rewardRepaidLog)

				amount := rewardRepaidLog.Reward.String()
				account := rewardRepaidLog.User.String()

				fmt.Printf("New RewardRepaid event with user: [%s], amount: [%s].", account, amount)

				reward := models.Reward{
					Address: account,
					Reward:  amount,
				}

				database.DB.Create(&reward)
				fiberlog.Infof("New reward with ID: [%d].", reward.ID)
			}
		}
	}()
}

func Staked() {
	log.Println("WatchStaked...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFStaked)

	sub, err := ff.WatchStaked(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to staked events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Staked Event subscription failed: %v", err)
			case stakedLog := <-ch:
				fmt.Print(stakedLog)

				amount := stakedLog.Amount.String()
				account := stakedLog.User.String()

				fmt.Printf("New staked event with user: [%s], amount: [%s].", account, amount)

				stake := models.Stake{
					Address: account,
					Amount:  amount,
				}

				database.DB.Create(&stake)
				fiberlog.Infof("New stake with ID: [%d].", stake.ID)
			}
		}
	}()
}

func Withdrawn() {
	log.Println("WatchWithdrawn...")

	ff := contract.FFInstance

	ch := make(chan *contract.FFWithdrawn)

	sub, err := ff.WatchWithdrawn(nil, ch, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to withdrawn events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Withdrawn Event subscription failed: %v", err)
			case withdrawnLog := <-ch:
				fmt.Print(withdrawnLog)

				amount := withdrawnLog.Amount.String()
				account := withdrawnLog.User.String()

				fmt.Printf("New withdrawn event with user: [%s], amount: [%s].", account, amount)

				withdraw := models.Withdraw{
					Address: account,
					Amount:  amount,
				}

				database.DB.Create(&withdraw)
				fiberlog.Infof("New withdraw with ID: [%d].", withdraw.ID)
			}
		}
	}()
}
