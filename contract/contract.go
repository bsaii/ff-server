package contract

import (
	"log"

	"github.com/bsaii/ff-server/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var FFInstance *FF
var FFTokenInstance *FFToken

func InitContract(client *ethclient.Client, cfg config.Config) {
	var err error

	FFInstance, err = NewFF(common.HexToAddress(cfg.FFContractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate FinanceForge contract: %v", err)
	}

	FFTokenInstance, err = NewFFToken(common.HexToAddress(cfg.FFTokenContractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate FinanceForgeToken contract: %v", err)
	}

}
