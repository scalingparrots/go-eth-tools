package explorer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Explorer represents the explorer
type Explorer struct {
	Endpoint string
	ApiKey   string
}

// LogResponse is the structure for the logs returned from the explorer API.
type LogResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []EventLog `json:"result"`
}

// EventLog is the structure for the logs returned from the explorer API.
type EventLog struct {
	Address         string       `json:"address"`
	Topics          []string     `json:"topics"`
	Data            string       `json:"data"`
	BlockHash       string       `json:"blockHash"`
	TimeStamp       HexTimestamp `json:"timeStamp"`
	TransactionHash string       `json:"transactionHash"`
}

// HexTimestamp is a timestamp in hex format
type HexTimestamp string

// NewExplorer returns a new explorer
func NewExplorer(endpoint string, apiKey string) *Explorer {
	return &Explorer{Endpoint: strings.TrimSuffix(endpoint, "/"), ApiKey: apiKey}
}

// constructURL constructs the final URL to make the request
func (e Explorer) constructURL(path string) string {
	separator := "?"
	if strings.Contains(path, "?") {
		separator = "&"
	}
	return fmt.Sprintf("%s%s%sapiKey=%s", e.Endpoint, path, separator, e.ApiKey)
}

// HttpGet performs a GET request to the specified path
func (e Explorer) HttpGet(path string) ([]byte, error) {
	url := e.constructURL(path)
	log.Println("URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Println("Error closing response body:", closeErr)
		}
	}()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetEventLogsByAddress is a generic function to get swap data given a topic.
func (e Explorer) GetEventLogsByAddress(address, topic string, from, to uint64) ([]EventLog, error) {
	module := "logs"
	action := "getLogs"
	fromBlock := from
	toBlock := to
	offset := "1000"

	url := fmt.Sprintf("?module=%s&action=%s&fromBlock=%d&toBlock=%d&address=%s&offset=%s&topic0=%s", module, action, fromBlock, toBlock, address, offset, topic)
	res, err := e.HttpGet(url)
	if err != nil {
		return nil, fmt.Errorf("error getting swap data: %w", err)
	}

	var response LogResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling swap data: %w", err)
	}

	logs := response.Result

	return logs, nil
}

// GetEventLogsByTopic is a generic function to get event logs given a topic.
func (e Explorer) GetEventLogsByTopic(address, topic string, from, to uint64) ([]EventLog, error) {
	module := "logs"
	action := "getLogs"
	fromBlock := from
	toBlock := to
	topic0 := topic
	page := "1"
	offset := "1000"

	url := fmt.Sprintf("?module=%s&action=%s&fromBlock=%d&toBlock=%d&address=%s&topic0=%s&page=%s&offset=%s", module, action, fromBlock, toBlock, address, topic0, page, offset)
	res, err := e.HttpGet(url)
	if err != nil {
		return nil, fmt.Errorf("error getting event logs: %w", err)
	}

	var response LogResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling event logs: %w", err)
	}

	logs := response.Result

	return logs, nil
}
