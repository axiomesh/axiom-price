package pricefeed

import "sync"

type PriceInfo struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Source string  `json:"source"`
}

type PricesStore struct {
	sync.RWMutex
	ReadMap  map[string]PriceInfo
	writeMap map[string]PriceInfo
}

func NewPricesStore() *PricesStore {
	return &PricesStore{
		ReadMap:  make(map[string]PriceInfo),
		writeMap: make(map[string]PriceInfo),
	}
}

func (ps *PricesStore) SwapMaps() {
	ps.Lock()
	defer ps.Unlock()
	ps.ReadMap, ps.writeMap = ps.writeMap, make(map[string]PriceInfo)
}

func (ps *PricesStore) UpdatePrice(info PriceInfo) {
	ps.Lock() // Lock for write operations
	ps.writeMap[info.Symbol] = info
	ps.Unlock()
}

func (ps *PricesStore) GetPrice(symbol string) (PriceInfo, bool) {
	ps.RLock() // Read lock for accessing ReadMap
	info, exists := ps.ReadMap[symbol]
	ps.RUnlock()
	return info, exists
}

func (ps *PricesStore) GetAllPrices() []PriceInfo {
	ps.RLock() // Ensure thread-safe read access
	defer ps.RUnlock()

	var prices []PriceInfo
	for _, price := range ps.ReadMap {
		prices = append(prices, price)
	}
	return prices
}
