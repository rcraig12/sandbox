package main

import (
	"fmt"
	"math"
	"math/big"
)

func outputTokenDetail(token *erc20token, lpool *liquiditypool) {

	// output the erc-20 or bep-20 information for a token

	ftotalSupply := new(big.Float)
	ftotalSupply.SetString(token.totalsupply.String())

	// Use token's decimals for the division
	divisor := math.Pow10(int(token.decimals))
	totalSupply := new(big.Float).Quo(ftotalSupply, big.NewFloat(divisor))

	// Convert totalSupply to string with desired precision
	// The 'f' format is for floating-point number, and -1 uses the smallest number of digits necessary
	totalSupplyStr := totalSupply.Text('f', -1)

	fmt.Printf("Contract Address: %s\nName: %s\nSymbol: %s\nSupply: %v\nDecimals: %d\n", token.contract, token.name, token.symbol, totalSupplyStr, token.decimals)

	// output the liquidiy information for the token
	// Calculate price based on reserve order
	tokenReserve := new(big.Float).SetInt(lpool.reserve0)
	stableCoinReserve := new(big.Float).SetInt(lpool.reserve1)

	//fmt.Println("tokenReserve:", tokenReserve)
	//fmt.Println("StableCoinReserve:", stableCoinReserve)
	fmt.Printf("Pair Address: %s\n", lpool.pair)

	fBalance := new(big.Float)
	fBalance.SetString(tokenReserve.String())
	balanceSTC := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(int(token.decimals))))
	fmt.Printf("Pooled %s: %f\n", token.symbol, balanceSTC)

	fBalance = new(big.Float)
	fBalance.SetString(stableCoinReserve.String())
	balanceWBNB := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("Pooled WBNB: %f\n", balanceWBNB)

	// First output the liquidity pool pair address

	// output the 
}