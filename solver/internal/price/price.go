package price

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type Cache struct {
	ETH       float64
	WETH      float64
	USDC      float64
	DAI       float64
	UpdatedAt time.Time
	mu        sync.RWMutex
}

var cache = &Cache{}

type CoinGeckoResponse struct {
	Ethereum map[string]float64 `json:"ethereum"`
	USDCoin  map[string]float64 `json:"usd-coin"`
	Dai      map[string]float64 `json:"dai"`
}

func FetchPrices() error {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	url := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum,usd-coin,dai&vs_currencies=usd"
	if apiKey != "" {
		url = "https://pro-api.coingecko.com/api/v3/simple/price?ids=ethereum,usd-coin,dai&vs_currencies=usd&x_cg_pro_api_key=" + apiKey
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var data CoinGeckoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("failed to parse prices: %w", err)
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.ETH = data.Ethereum["usd"]
	cache.WETH = data.Ethereum["usd"]
	cache.USDC = data.USDCoin["usd"]
	cache.DAI = data.Dai["usd"]
	cache.UpdatedAt = time.Now()

	fmt.Printf("Prices updated: ETH=$%.2f, USDC=$%.2f, DAI=$%.2f\n", cache.ETH, cache.USDC, cache.DAI)

	return nil
}

func GetPrices() map[string]float64 {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	return map[string]float64{
		"eth":  cache.ETH,
		"weth": cache.WETH,
		"usdc": cache.USDC,
		"dai":  cache.DAI,
	}
}

func StartPriceUpdater(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		FetchPrices()

		for range ticker.C {
			if err := FetchPrices(); err != nil {
				fmt.Printf("Error updating prices: %v\n", err)
			}
		}
	}()
}
