package main

func main() {

	BuildStableCoinTable()
	exchange := SetExchange("PancakeSwap")
	provider := SetProvider("BSCScan", "https://bsc-dataseed.binance.org/")
	token := SetToken("0x19ae49b9f38dd836317363839a5f6bfbfa7e319a", provider)
	lpool := SetLiquidity(token, exchange, provider)

	outputTokenDetail(token, lpool)
	
}