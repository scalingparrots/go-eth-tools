package gas

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// MainnetEndpoint is the mainnet endpoint for the Polygon gas station
	MainnetEndpoint = "https://gasstation-mainnet.matic.network"
	// TestnetEndpoint is the testnet endpoint for the Polygon gas station
	TestnetEndpoint = "https://gasstation-mumbai.matic.today"
)

// Response represents the response from the Polygon gas station
type Response struct {
	SafeLow struct {
		MaxPriorityFee float64 `json:"maxPriorityFee"`
		MaxFee         float64 `json:"maxFee"`
	} `json:"safeLow"`
	Standard struct {
		MaxPriorityFee float64 `json:"maxPriorityFee"`
		MaxFee         float64 `json:"maxFee"`
	} `json:"standard"`
	Fast struct {
		MaxPriorityFee float64 `json:"maxPriorityFee"`
		MaxFee         float64 `json:"maxFee"`
	} `json:"fast"`
	EstimatedBaseFee float64 `json:"estimatedBaseFee"`
	BlockTime        int     `json:"blockTime"`
	BlockNumber      int     `json:"blockNumber"`
}

type PolygonGasStation struct {
	Endpoint string
}

func NewPolygonGasStation(endpoint string) *PolygonGasStation {
	if endpoint == "" {
		endpoint = TestnetEndpoint
	}
	return &PolygonGasStation{Endpoint: endpoint}
}

// FetchGasPriceFromPolygon fetches the gas price from the Polygon gas station
func (gasStation PolygonGasStation) FetchGasPriceFromPolygon() (float64, float64, float64, error) {
	resp, err := http.Get(gasStation.Endpoint)
	if err != nil {
		return 0, 0, 0, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Println("Error closing response body:", closeErr)
		}
	}()

	var gasResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&gasResponse); err != nil {
		return 0, 0, 0, err
	}

	return gasResponse.Fast.MaxFee, gasResponse.Fast.MaxPriorityFee, gasResponse.EstimatedBaseFee, nil
}
