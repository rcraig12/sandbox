package main

import "log"

func newExchange(name string) *exchange {
	exchange := exchange{name: name}

	switch name {
	case "PancakeSwap":
		exchange.factoryABI = "./abi/panckakeswap.factory.abi.json"
		exchange.pairABI = "./abi/panckakeswap.pair.abi.json"
	default:
		log.Fatalf("Exchange named \"" + exchange.name + "\" is unknown.")
	}
	return &exchange
}
