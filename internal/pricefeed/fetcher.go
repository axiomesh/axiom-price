package pricefeed

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"net/url"
	"price-aggregator/pkg/config"
	"strconv"
	"strings"
	"time"
)

type okxResponse struct {
	Data []struct {
		InstId string `json:"instId"`
		IdxPx  string `json:"idxPx"`
	}
}

// Define a struct to unmarshal the response from exchange APIs
type binanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type bybitResponse struct {
	Result []struct {
		Symbol    string `json:"symbol"`
		LastPrice string `json:"last_price"`
	} `json:"result"`
}

const (
	// BINANCE_API_URL api rate limit: 6000 /min
	BINANCE_API_URL = "https://api.binance.com/api/v3/ticker/price"
	BYBIT_API_URL   = "https://api.bybit.com/v2/public/tickers"
	// OKX_API_URL api rate limit: 20 /2s
	OKX_API_URL      = "https://www.okx.com/api/v5/market/index-tickers?quoteCcy=USDT"
	OKX_API_USDT_URL = "https://www.okx.com/api/v5/market/index-tickers?instId=USDT-USD"
	OKX_API_USDC_URL = "https://www.okx.com/api/v5/market/index-tickers?instId=USDC-USD"

	BINANCE_SOURCE = "binance"
	BYBIT_SOURCE   = "bybit"
	OKX_SOURCE     = "okx"
	MIXED_SOURCE   = "mixed"
	TEST_SOURCE    = "test"
	CONFIG_SOURCE  = "config"

	HTTP_CLIENT_TIMEOUT = 5 * time.Second
)

func fetchUSDPrices(pricesMap map[string]float64, cfg *config.Config) error {
	var okxRes okxResponse

	httpClient := NewHTTPClient(cfg)
	urls := []string{OKX_API_USDC_URL, OKX_API_USDT_URL}

	for _, u := range urls {
		resp, err := httpClient.Get(u)
		if err != nil {
			return err
		}

		if err := json.NewDecoder(resp.Body).Decode(&okxRes); err != nil {
			return err
		}

		for _, pair := range okxRes.Data {
			if strings.HasSuffix(pair.InstId, "USD") {
				symbol := strings.Replace(pair.InstId, "-", "", -1)
				price, parseErr := strconv.ParseFloat(pair.IdxPx, 64)
				if parseErr != nil {
					log.Printf("Error parsing price for %s: %v", pair.InstId, parseErr)
					continue
				}

				roundedPrice := math.Round(price*1e8) / 1e8
				pricesMap[symbol] = roundedPrice
			}
		}
		resp.Body.Close()
	}
	rate, err := strconv.ParseFloat(cfg.AxcUsd, 64)
	if err != nil {
		return err
	}
	pricesMap["AXCUSD"] = rate

	return nil
}

func fetchFromOKX(pricesMap map[string][]*PriceInfo, cfg *config.Config) error {
	var okxRes okxResponse

	httpClint := NewHTTPClient(cfg)
	resp, err := httpClint.Get(OKX_API_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&okxRes); err != nil {
		return err
	}

	// Convert the API response to our PriceInfo struct
	for _, pair := range okxRes.Data {
		if strings.HasSuffix(pair.InstId, "USDT") {
			symbol := strings.Replace(pair.InstId, "-", "", -1)
			price, parseErr := strconv.ParseFloat(pair.IdxPx, 64)
			if parseErr != nil {
				log.Printf("Error parsing price for %s: %v", pair.InstId, parseErr)
				continue
			}
			roundedPrice := math.Round(price*1e8) / 1e8
			pricesMap[symbol] = append(pricesMap[symbol], &PriceInfo{
				Price:  roundedPrice,
				Source: OKX_SOURCE,
			})
		}
	}
	return nil
}

// fetchFromBinance fetches the price for a specific symbol from an exchange API.
func fetchFromBinance(pricesMap map[string][]*PriceInfo, cfg *config.Config) error {
	var bRes []binanceResponse

	httpClient := NewHTTPClient(cfg)
	resp, err := httpClient.Get(BINANCE_API_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&bRes); err != nil {
		return err
	}

	// Convert the API response to our PriceInfo struct
	for _, pair := range bRes {
		if strings.HasSuffix(pair.Symbol, "USDT") {
			price, parseErr := strconv.ParseFloat(pair.Price, 64)
			if parseErr != nil {
				log.Printf("Error parsing price for %s: %v", pair.Symbol, parseErr)
				continue
			}
			roundedPrice := math.Round(price*1e8) / 1e8
			pricesMap[pair.Symbol] = []*PriceInfo{
				{
					Price:  roundedPrice,
					Source: BINANCE_SOURCE,
				},
			}
		}
	}

	return nil
}

func fetchFromBybit(pricesMap map[string][]*PriceInfo, cfg *config.Config) error {
	var bRes bybitResponse

	httpClient := NewHTTPClient(cfg)
	resp, err := httpClient.Get(BYBIT_API_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&bRes); err != nil {
		return err
	}

	// Process each pair in the response
	for _, pair := range bRes.Result {
		if strings.HasSuffix(pair.Symbol, "USDT") {
			price, parseErr := strconv.ParseFloat(pair.LastPrice, 64)
			if parseErr != nil {
				log.Printf("Error parsing price for %s: %v", pair.Symbol, parseErr)
				continue
			}
			roundedPrice := math.Round(price*1e8) / 1e8
			pricesMap[pair.Symbol] = append(pricesMap[pair.Symbol], &PriceInfo{
				Price:  roundedPrice,
				Source: BYBIT_SOURCE,
			})
		}
	}

	return nil
}

// FetchPrices is the main function to fetch prices from exchanges and update the store.
func FetchPrices(store *PricesStore, errChan chan<- error, cfg *config.Config) {
	pricesMap := make(map[string][]*PriceInfo)
	// fetch from binance
	binanceErr := fetchFromBinance(pricesMap, cfg)
	if binanceErr != nil {
		errChan <- binanceErr
	}
	// fetch from bybit
	bybitErr := fetchFromBybit(pricesMap, cfg)
	if bybitErr != nil {
		errChan <- bybitErr
	}
	// fetch from okx
	okxErr := fetchFromOKX(pricesMap, cfg)
	if okxErr != nil {
		errChan <- okxErr
	}
	// fetch usd
	rateMap := make(map[string]float64)
	usdErr := fetchUSDPrices(rateMap, cfg)
	if usdErr != nil {
		errChan <- usdErr
	}
	if usdErr != nil {
		log.Printf("Error fetching USD price: %v", usdErr)
		return
	}
	if binanceErr != nil && bybitErr != nil && okxErr != nil {
		log.Printf("Error fetching prices from all exchanges: %v, %v, %v", binanceErr, bybitErr, okxErr)
		return
	}
	usdtRate := rateMap["USDTUSD"]
	for _, s := range cfg.WalletTickerList {
		if s.Ticker == "AXC" {
			price := rateMap[s.Ticker+"USD"]
			store.UpdatePrice(PriceInfo{
				Symbol: s.Ticker + "USD",
				Price:  price,
				Source: CONFIG_SOURCE,
			})
			store.UpdatePrice(PriceInfo{
				Symbol: s.Alias + "USD",
				Price:  price,
				Source: CONFIG_SOURCE,
			})
		} else if s.Ticker == "USDT" {
			store.UpdatePrice(PriceInfo{
				Symbol: s.Ticker + "USD",
				Price:  usdtRate,
				Source: OKX_SOURCE,
			})
			store.UpdatePrice(PriceInfo{
				Symbol: s.Alias + "USD",
				Price:  usdtRate,
				Source: OKX_SOURCE,
			})
		} else {
			key := s.Ticker + "USDT"
			priceInfos := pricesMap[key]
			length := len(priceInfos)
			var mixedPrice float64
			for _, info := range priceInfos {
				mixedPrice += info.Price
			}
			mixedPrice /= float64(length)
			mixedPrice /= usdtRate
			roundedPrice := math.Round(mixedPrice*1e8) / 1e8
			// update store
			if length > 1 {
				store.UpdatePrice(PriceInfo{
					Symbol: s.Ticker + "USD",
					Price:  roundedPrice,
					Source: MIXED_SOURCE,
				})
				store.UpdatePrice(PriceInfo{
					Symbol: s.Alias + "USD",
					Price:  roundedPrice,
					Source: MIXED_SOURCE,
				})
			} else {
				store.UpdatePrice(PriceInfo{
					Symbol: s.Ticker + "USD",
					Price:  roundedPrice,
					Source: priceInfos[0].Source,
				})
				store.UpdatePrice(PriceInfo{
					Symbol: s.Alias + "USD",
					Price:  roundedPrice,
					Source: priceInfos[0].Source,
				})
			}
		}
	}
	store.SwapMaps()
	go PushDataToContract(cfg, store)
}

func NewHTTPClient(cfg *config.Config) *http.Client {
	var httpClient *http.Client
	if cfg.UseProxy {
		proxyURL, err := url.Parse(cfg.ProxyURL)
		if err != nil {
			log.Fatal(err)
		}
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
			Timeout: HTTP_CLIENT_TIMEOUT,
		}
	} else {
		httpClient = &http.Client{
			Timeout: HTTP_CLIENT_TIMEOUT,
		}
	}
	return httpClient
}
