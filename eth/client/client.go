package client

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scalingparrots/go-eth-tools/eth/gas"
	"math/big"
	"strings"
)

// RPCClient is the ethereum client
type RPCClient struct {
	client     *ethclient.Client
	chainId    *big.Int
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	gasStation string
}

// Params is the ethereum parameters
type Params struct {
	ChainId    int
	RPC        string
	privateKey string
	gasStation string
}

// NewClient initializes the ethereum client
func NewClient(params *Params) (*RPCClient, error) {
	gasStation := params.gasStation

	client, err := ethclient.Dial(params.RPC)
	if err != nil {
		return nil, err
	}

	privateKey := params.privateKey

	if strings.HasPrefix(privateKey, "0x") {
		privateKey = privateKey[2:]
	}

	// Convert the private key to an ECDSA private key
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKeyECDSA := privateKeyECDSA.PublicKey

	return &RPCClient{client: client, chainId: big.NewInt(int64(params.ChainId)), privateKey: privateKeyECDSA, publicKey: &publicKeyECDSA, gasStation: gasStation}, nil
}

// Common transaction preparation logic
func (e *RPCClient) prepareTransaction(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, e.chainId)
	if err != nil {
		return nil, err
	}

	sender := crypto.PubkeyToAddress(*e.publicKey)
	nonce, err := e.client.NonceAt(ctx, sender, nil)
	if err != nil {
		return nil, err
	}

	// Get the current gas price from polygon station
	e.gasStation = "https://gasstation-mainnet.matic.network"
	fastGasFee, fastGasPriorityFee, _, err := gas.NewPolygonGasStation(e.gasStation).FetchGasPriceFromPolygon()
	if err != nil {
		return nil, err
	}

	// GasTipCap = fastGasPriorityFee
	gasTipCap := big.NewInt(int64(fastGasPriorityFee * 1e9))

	// GasFeeCap = fastGasFee
	gasFeeCap := big.NewInt(int64(fastGasFee * 1e9))

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasFeeCap = gasFeeCap
	auth.GasTipCap = gasTipCap
	auth.GasLimit = uint64(1000000)

	return auth, nil
}
