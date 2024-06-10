package listener

import (
	"fmt"
	"log"

	"github.com/bsaii/ff-server/contract"
	"github.com/bsaii/ff-server/database"
	"github.com/bsaii/ff-server/models"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func Approval() {
	log.Println("WatchApproval...")

	fftoken := contract.FFTokenInstance

	ch := make(chan *contract.FFTokenApproval)

	sub, err := fftoken.WatchApproval(nil, ch, nil, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to approval events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Approval Event subscription failed: %v", err)
			case approvalLog := <-ch:
				fmt.Print(approvalLog)

				owner := approvalLog.Owner.Hex()
				spender := approvalLog.Spender.Hex()
				value := approvalLog.Value.String()
				fmt.Printf("New approval event with owner: [%s], spender: [%s], value: [%s].", owner, spender, value)

				approval := models.Approval{
					Owner:   owner,
					Spender: spender,
					Value:   value,
				}

				database.DB.Create(&approval)
				fiberlog.Infof("New approval with ID: [%d].", approval.ID)
			}
		}
	}()
}

func Transfer() {
	log.Println("WatchTransfer...")

	fftoken := contract.FFTokenInstance

	ch := make(chan *contract.FFTokenTransfer)

	sub, err := fftoken.WatchTransfer(nil, ch, nil, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to transfer events: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatalf("Transfer Event subscription failed: %v", err)
			case transferLog := <-ch:
				fmt.Print(transferLog)

				from := transferLog.From.Hex()
				to := transferLog.To.Hex()
				value := transferLog.Value.String()
				fmt.Printf("New transfer event with from: [%s], to: [%s], value: [%s].", from, to, value)

				transfer := models.Transfer{
					From:  from,
					To:    to,
					Value: value,
				}

				database.DB.Create(&transfer)
				fiberlog.Info("New Transfer with ID: [%d].", transfer.ID)
			}
		}
	}()
}
