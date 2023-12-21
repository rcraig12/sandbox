package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getTokenName(client *ethclient.Client, tokenAddress common.Address) (string, error) {
	// Read the ABI from the file
	abiData, err := os.ReadFile("abi/erc20.abi.json") // Specify the path to your ABI file
	if err != nil {
		return "", fmt.Errorf("failed to read ABI file: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Pack the data to send a transaction
	data, err := parsedABI.Pack("name")
	if err != nil {
		return "", fmt.Errorf("failed to pack data for 'name': %v", err)
	}

	// Call the contract
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}
	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %v", err)
	}

	// Unpack the data
	var result interface{}
	err = parsedABI.UnpackIntoInterface(&result, "name", res)
	if err != nil {
		return "", fmt.Errorf("failed to unpack response: %v", err)
	}

	// Type assert the result to string
	tokenName, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("result is not a string")
	}

	return tokenName, nil

}

func getTokenSymbol(client *ethclient.Client, tokenAddress common.Address) (string, error) {
	// Read the ABI from the file
	abiData, err := os.ReadFile("abi/erc20.abi.json") // Specify the path to your ABI file
	if err != nil {
		return "", fmt.Errorf("failed to read ABI file: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Pack the data to send a transaction
	data, err := parsedABI.Pack("symbol")
	if err != nil {
		return "", fmt.Errorf("failed to pack data for 'name': %v", err)
	}

	// Call the contract
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}
	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %v", err)
	}

	// Unpack the data
	var result interface{}
	err = parsedABI.UnpackIntoInterface(&result, "name", res)
	if err != nil {
		return "", fmt.Errorf("failed to unpack response: %v", err)
	}

	// Type assert the result to string
	tokenName, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("result is not a string")
	}

	return tokenName, nil

}

func getTokenDecimals(client *ethclient.Client, tokenAddress common.Address) (uint8, error) {
	// Read the ABI from the file
	abiData, err := os.ReadFile("abi/erc20.abi.json") // Specify the path to your ABI file
	if err != nil {
		return 0, fmt.Errorf("failed to read ABI file: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		return 0, fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Pack the data to send a transaction
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		return 0, fmt.Errorf("failed to pack data for 'decimals': %v", err)
	}

	// Call the contract
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}
	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to call contract: %v", err)
	}

	// Unpack the data
	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", res)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack response: %v", err)
	}

	return decimals, nil
}

func getTokenTotalSupply(client *ethclient.Client, tokenAddress common.Address) (*big.Int, error) {
	// Read the ABI from the file
	abiData, err := os.ReadFile("abi/erc20.abi.json") // Specify the path to your ABI file
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	// Pack the data to send a transaction
	data, err := parsedABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for 'totalSupply': %v", err)
	}

	// Call the contract
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}
	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %v", err)
	}

	// Unpack the data
	var totalSupply *big.Int
	err = parsedABI.UnpackIntoInterface(&totalSupply, "totalSupply", res)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack response: %v", err)
	}

	return totalSupply, nil
}

func SetToken( contract string, provider *provider) *erc20token {
	erc20 := erc20token{ contract: contract}

	client, err := ethclient.Dial(provider.url)
	if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	erc20.name, err = getTokenName(client, common.HexToAddress(erc20.contract))
	if err != nil {
		log.Fatal(err)
	}
	erc20.symbol, err = getTokenSymbol(client, common.HexToAddress(erc20.contract))
	if err != nil {
		log.Fatal(err)
	}
	erc20.totalsupply, err = getTokenTotalSupply(client, common.HexToAddress(erc20.contract))
	if err != nil {
		log.Fatal(err)
	}
	erc20.decimals, err = getTokenDecimals(client, common.HexToAddress(erc20.contract))
	if err != nil {
		log.Fatal(err)
	}
	return &erc20
}