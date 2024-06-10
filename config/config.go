package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	EthNodeURL             string
	DbPath                 string
	FFContractAddress      string
	FFTokenContractAddress string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set in the environment")
	}

	nodeUrl := os.Getenv("ETH_NODE_URL")
	if nodeUrl == "" {
		log.Fatal("ETH_NODE_URL not set in the environment")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH not set in the environment")
	}

	ffAddress := os.Getenv("FF_ADDRESS")
	if ffAddress == "" {
		log.Fatal("FF_ADDRESS not set in the environment")
	}

	ffTAddress := os.Getenv("FFT_ADDRESS")
	if ffTAddress == "" {
		log.Fatal("FFT_ADDRESS not set in the environment")
	}

	return Config{
		EthNodeURL:             nodeUrl,
		DbPath:                 dbPath,
		FFContractAddress:      ffAddress,
		FFTokenContractAddress: ffTAddress,
		Port:                   port,
	}
}
