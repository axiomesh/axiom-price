package pricefeed

import (
	"log"
	"math/big"
	"price-aggregator/contract/pricefeeds"
	"price-aggregator/pkg/config"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PushDataToContract(config *config.Config, store *PricesStore) {
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:])
	sender := crypto.PubkeyToAddress(privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(config.ChainId))) // 使用正确的链ID
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contract, err := pricefeeds.NewPricefeeds(common.HexToAddress(config.PriceFeedsContract), client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	timestamp := time.Now().UTC().Unix()
	readMap := store.ReadMap
	if readMap == nil {
		log.Printf("No data to push")
		return
	}
	axcInfo, ok := readMap["AXCUSD"]
	if !ok {
		log.Fatalf("AXCUSD not found in store")
		return
	}
	axcPrice := axcInfo.Price
	push := make([]pricefeeds.PriceFeedsTickerInfo, len(config.AATickerList))
	for i, s := range config.AATickerList {
		info, ok := readMap[s.Ticker+"USD"]
		if !ok {
			push[i] = pricefeeds.PriceFeedsTickerInfo{}
			continue
		}
		priceScaled := new(big.Float).Mul(big.NewFloat(info.Price), big.NewFloat(1e8))
		priceFinal := new(big.Int)
		priceScaled.Quo(priceScaled, big.NewFloat(axcPrice)).Int(priceFinal)
		push[i] = pricefeeds.PriceFeedsTickerInfo{
			Price:         priceFinal,
			Decimals:      big.NewInt(100000000),
			Timestamp:     big.NewInt(timestamp),
			Source:        info.Source,
			Sender:        sender,
			Currency:      "AXC",
			Ticker:        s.Ticker,
			TickerAddress: common.HexToAddress(s.Address),
		}
	}
	tx, err := contract.UpdatePrices(auth, push)
	if err != nil {
		log.Printf("Failed to push data to contract: %v", err)
		return
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
}
