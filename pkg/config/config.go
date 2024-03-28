package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	UseProxy             bool   `yaml:"use_proxy"`
	ProxyURL             string `yaml:"proxy_url"`
	HTTPPort             string `yaml:"http_port"`
	PrivateKey           string `yaml:"private_key"`
	PriceFeedsContract   string `yaml:"price_feeds_contract"`
	RpcUrl               string `yaml:"rpc_url"`
	ChainId              int    `yaml:"chain_id"`
	AxcUsd               string `yaml:"axc_usd"`
	FetchSecondsDuration int    `yaml:"fetch_seconds_duration"`
	EnvPrivateKeyName    string `yaml:"env_private_key_name"`
	AATickerList         []struct {
		Ticker  string `yaml:"ticker"`
		Symbol  string `yaml:"symbol"`
		Address string `yaml:"address"`
	} `yaml:"aa_ticker_list"`
	WalletTickerList []struct {
		Ticker string `yaml:"ticker"`
		Alias  string `yaml:"alias"`
	} `yaml:"wallet_ticker_list"`
}

const (
	configPath = "./config.yml"
)

func LoadConfig() (*Config, error) {
	config := &Config{}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("yaml config get err #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}
	log.Printf("load config successfully: %+v", config)

	// read private key from environment
	envPrivateKey := os.Getenv(config.EnvPrivateKeyName)
	if envPrivateKey != "" {
		config.PrivateKey = envPrivateKey
	} else if config.PrivateKey == "" {
		log.Fatalf("private key is required")
		return nil, err
	}
	return config, nil
}
