package main

import (
	"log"
	"os"
)

func setEnvironmentVariables() (err error) {

	err = os.Setenv("WBNB", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	if err != nil {
		log.Fatal(err)
	}
	
	err = os.Setenv("WETH", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	if err != nil {
		log.Fatal(err)
	}
	
	return err

}