package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
    // PancakeSwap Factory Contract Address
    factoryAddress = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
    // Binance-Pegged USD (like BUSD) Address
    stableCoinAddress = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
		//stableCoinAddress = "0x4fabb145d64652a948d72533023f6e7a623c7c53"
		
    // Your Token Address
    tokenAddress = "0x19ae49b9f38dd836317363839a5f6bfbfa7e319a"
)

// func outputLiquidity(token, exchange, provider) {
// 	client, err := ethclient.Dial(provider.url)
// 	if err != nil {
// 			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
// 	}

// 	tokenUSDPrice, err := getTokenUSDPrice(client)
// 	if err != nil {
// 			log.Fatalf("Failed to get token USD price: %v", err)
// 	}

// 	fmt.Printf("Token USD Price: %s\n", tokenUSDPrice.Text('f', 6))
// }

func getTokenUSDPrice(client *ethclient.Client) (*big.Float, error) {
	// Read the ABI from the file
	abiData, err := os.ReadFile("factory_abi.json") // Specify the correct path
	if err != nil {
			return nil, fmt.Errorf("failed to read ABI file: %v", err)
	}

	factory, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
			return nil, fmt.Errorf("failed to parse factory ABI: %v", err)
	}


	token := common.HexToAddress(tokenAddress)
	stableCoin := common.HexToAddress(stableCoinAddress)
	fAddress := common.HexToAddress(factoryAddress)

	// Get the pair address
	data, err := factory.Pack("getPair", token, stableCoin)
	if err != nil {
			return nil, fmt.Errorf("failed to pack data for getPair: %v", err)
	}

	msg := ethereum.CallMsg{To: &fAddress, Data: data}
	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil || len(res) == 0 {
			return nil, fmt.Errorf("failed to call contract for getPair: %v", err)
	}

	// Read the Pair ABI from the file
	pairAbiData, err := os.ReadFile("pair_abi.json") // Specify the correct path
	if err != nil {
			return nil, fmt.Errorf("failed to read Pair ABI file: %v", err)
	}

	pairAddress := common.BytesToAddress(res)
	pair, err := abi.JSON(strings.NewReader(string(pairAbiData)))
    if err != nil {
        return nil, fmt.Errorf("failed to parse pair ABI: %v", err)
    }

	// Get the reserves
	data, err = pair.Pack("getReserves")
	if err != nil {
			return nil, fmt.Errorf("failed to pack data for getReserves: %v", err)
	}

	msg = ethereum.CallMsg{To: &pairAddress, Data: data}
	res, err = client.CallContract(context.Background(), msg, nil)
	if err != nil {
			return nil, fmt.Errorf("failed to call contract for getReserves: %v", err)
	}

	fmt.Printf("pair:%v\n",msg)

	var reserves struct {
			Reserve0           *big.Int
			Reserve1           *big.Int
			BlockTimestampLast uint32
	}

	err = pair.UnpackIntoInterface(&reserves, "getReserves", res)
	if err != nil {
			return nil, fmt.Errorf("failed to unpack reserves: %v", err)
	}

	fmt.Printf("reserves:%v",reserves)

	// Calculate price based on reserve order
	tokenReserve := new(big.Float).SetInt(reserves.Reserve0)
	stableCoinReserve := new(big.Float).SetInt(reserves.Reserve1)

	fmt.Println("tokenReserve:", tokenReserve)
	fmt.Println("StableCoinReserve:", stableCoinReserve)

	fBalance := new(big.Float)
	fBalance.SetString(tokenReserve.String())
	balanceSTC := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(9)))
	fmt.Printf("STC Balance: %f\n", balanceSTC)

	fBalance = new(big.Float)
	fBalance.SetString(stableCoinReserve.String())
	balanceWBNB := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Printf("WBNB Balance: %f\n", balanceWBNB)

	/*

		To get the price of a token the function is:

		BNB USD Price * (Paired Currency / token) if bscscan mainnet
		ETH USD Price * (Paired Currency / token) if ethereum mainnet

		TODO: Work out how to get Paired Currency information dynamically and how to get token price from it.

	*/

	// Adjust the order if necessary
	// You might need additional logic here to ensure that tokenReserve and stableCoinReserve
	// correspond to the correct tokens

	return new(big.Float).Quo(stableCoinReserve, tokenReserve), nil
}