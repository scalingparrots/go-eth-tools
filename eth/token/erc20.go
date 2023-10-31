package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type ERC20 struct {
	instance *bind.BoundContract
	Address  common.Address
	ABI      *abi.ABI
}

func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	parsedABI, err := abi.JSON(strings.NewReader(Erc20AbiJson))
	if err != nil {
		return nil, err
	}

	return &ERC20{
		Address:  address,
		ABI:      &parsedABI,
		instance: bind.NewBoundContract(address, parsedABI, backend, backend, backend),
	}, nil
}

// Name gets the name of the token.
func (e *ERC20) Name(opts *bind.CallOpts) (string, error) {
	var result []interface{}
	err := e.instance.Call(opts, &result, "name")
	if err != nil {
		return "", err
	}

	return result[0].(string), nil
}

// Symbol gets the symbol of the token.
func (e *ERC20) Symbol(opts *bind.CallOpts) (string, error) {
	var result []interface{}
	err := e.instance.Call(opts, &result, "symbol")
	if err != nil {
		return "", err
	}

	return result[0].(string), nil
}

// Decimals gets the number of decimals of the token.
func (e *ERC20) Decimals(opts *bind.CallOpts) (uint8, error) {
	var result []interface{}
	err := e.instance.Call(opts, &result, "decimals")
	if err != nil {
		return 0, err
	}

	return result[0].(uint8), nil
}

// TotalSupply gets the total token supply.
func (e *ERC20) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var result []interface{}
	err := e.instance.Call(opts, &result, "totalSupply")
	if err != nil {
		return nil, err
	}

	return result[0].(*big.Int), nil
}

// BalanceOf gets the account's balance of the specified token.
func (e *ERC20) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var result []interface{}
	err := e.instance.Call(opts, &result, "balanceOf", account)
	if err != nil {
		return nil, err
	}

	return result[0].(*big.Int), nil
}

// Transfer transfers the specified amount of tokens to the specified address.

// Approve approves the passed address to spend the specified amount of tokens on behalf of the owner.

var Erc20AbiJson = `
	[
    {
        "constant": true,
        "inputs": [],
        "name": "name",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_spender",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "totalSupply",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_from",
                "type": "address"
            },
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "transferFrom",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "decimals",
        "outputs": [
            {
                "name": "",
                "type": "uint8"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [
            {
                "name": "_owner",
                "type": "address"
            }
        ],
        "name": "balanceOf",
        "outputs": [
            {
                "name": "balance",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [],
        "name": "symbol",
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "_to",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "transfer",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "constant": true,
        "inputs": [
            {
                "name": "_owner",
                "type": "address"
            },
            {
                "name": "_spender",
                "type": "address"
            }
        ],
        "name": "allowance",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "payable": true,
        "stateMutability": "payable",
        "type": "fallback"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": true,
                "name": "spender",
                "type": "address"
            },
            {
                "indexed": false,
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Approval",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "name": "from",
                "type": "address"
            },
            {
                "indexed": true,
                "name": "to",
                "type": "address"
            },
            {
                "indexed": false,
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Transfer",
        "type": "event"
    }
]
	`
