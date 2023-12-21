package main

import (
	"fmt"
)



func main() {

	BuildStableCoinTable()
	newExchange("PancakeSwap")
	provider := newProvider("BSCScan", "https://bsc-dataseed.binance.org/")
	//exchange := 
	token := newToken("0x19ae49b9f38dd836317363839a5f6bfbfa7e319a", provider)
	

	//fmt.Println(provider)
	//fmt.Println(token)
	outputTokenDetail(token)
	//outputLiquidity(token, exchange, provider)

	fmt.Println( "WBNB::", stableCoin["WBNB"])
	
}