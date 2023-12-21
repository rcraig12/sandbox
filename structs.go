package main

import (
	"math/big"
)

type provider struct {
	name string
	url string
}

type erc20token struct {
	contract string
	name string
	symbol string
	totalsupply *big.Int
	decimals uint8
}

type exchange struct {
	name string
	address string
	factoryABI string
	pairABI string
}

type liquiditypool struct {
	token string
	pair string
	token0 string
	reserve0 *big.Int
	token1 string
	reserve1 *big.Int
	BlockTimestampLast uint32
}