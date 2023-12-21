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
	factoryABI string
	pairABI string
}