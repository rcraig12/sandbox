package main

import (
	"fmt"
	"os"
)



func main() {

	err := setEnvironmentVariables()
	if err != nil {
		panic("Could not set environment variables.")
	}

	fmt.Println( "WBNB:", os.Getenv("WBNB"))
	fmt.Println( "WETH:", os.Getenv("WETH"))

	provider := newProvider("BSCScan", "https://bsc-dataseed.binance.org/")
	token := newToken("0x19ae49b9f38dd836317363839a5f6bfbfa7e319a", provider)
	

	//fmt.Println(provider)
	//fmt.Println(token)
	outputTokenDetail(token)
	
}