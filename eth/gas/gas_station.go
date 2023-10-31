package gas

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// PolygonGasStationEndpoint is the endpoint for the Polygon gas station
	PolygonGasStationEndpoint = "https://gasstation-mainnet.matic.network"
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

// FetchGasPriceFromPolygon fetches the gas price from the Polygon gas station
func FetchGasPriceFromPolygon() (float64, float64, float64, error) {
	resp, err := http.Get(PolygonGasStationEndpoint)
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
