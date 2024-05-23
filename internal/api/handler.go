package api

import (
	"encoding/json"
	"log"
	"net/http"
	"price-aggregator/internal/pricefeed"
	"strings"
)

// PriceResponse defines the structure for the price response JSON
type PriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,omitempty"`
	Error  string  `json:"error,omitempty"`
}

type Response struct {
	Data    []pricefeed.PriceInfo `json:"data"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
}

// HandleTickerPrice handles requests for /api/ticker/price?symbol=XXX
func HandleTickerPrice(w http.ResponseWriter, r *http.Request, store *pricefeed.PricesStore) {
	querySymbols := r.URL.Query().Get("symbol")

	var prices []pricefeed.PriceInfo
	if querySymbols == "" {
		prices = store.GetAllPrices()
	} else {
		// Parse the query string for specific symbols
		symbols := strings.Split(querySymbols, ",")
		for _, symbol := range symbols {
			if priceInfo, exists := store.GetPrice(symbol); exists {
				prices = append(prices, priceInfo)
			}
		}
	}

	// Check if prices are empty and return test prices for wallet tokens
	if len(prices) == 0 && len(store.ReadMap) == 0 {
		prices = []pricefeed.PriceInfo{
			{Symbol: "wETHUSD", Price: 3500.00, Source: pricefeed.TEST_SOURCE},
			{Symbol: "ETHUSD", Price: 3500.00, Source: pricefeed.TEST_SOURCE},
			{Symbol: "USDCUSD", Price: 1.0004, Source: pricefeed.TEST_SOURCE},
			{Symbol: "wUSDCUSD", Price: 1.0004, Source: pricefeed.TEST_SOURCE},
			{Symbol: "USDTUSD", Price: 0.9996, Source: pricefeed.TEST_SOURCE},
			{Symbol: "wUSDTUSD", Price: 0.9996, Source: pricefeed.TEST_SOURCE},
			{Symbol: "AXCUSD", Price: 0.1300, Source: pricefeed.TEST_SOURCE},
			{Symbol: "wAXCUSD", Price: 0.1300, Source: pricefeed.TEST_SOURCE},
		}
	}

	res := Response{
		Data:    prices,
		Code:    0,
		Message: "success",
	}
	jsonResponse(w, res, http.StatusOK)
}

// jsonResponse writes a JSON response
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}
