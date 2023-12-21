package main

import "log"

func SetExchange(name string) *exchange {
	exchange := exchange{name: name}

	switch name {
	case "PancakeSwap":
		exchange.address = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
		exchange.factoryABI = "./abi/panckakeswap.factory.abi.json"
		exchange.pairABI = "./abi/panckakeswap.pair.abi.json"
	default:
		log.Fatalf("Exchange named \"" + exchange.name + "\" is unknown.")
	}
	return &exchange
}
