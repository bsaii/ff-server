package config

import (
	"os"
)

type Config struct {
	EthNodeURL string
	DbPath     string
}

func LoadConfig() Config {
	return Config{
		EthNodeURL: os.Getenv("ETH_NODE_URL"),
		DbPath:     os.Getenv("DB_PATH"),
	}
}
