package main

import (
	"log"

	reporting "github.com/hyperledger/BMR1/asset-transfer-reporting/chaincode-go/smart-contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	reportingSmartContract, err := contractapi.NewChaincode(&reporting.SmartContract{})
	if err != nil {
		log.Panicf("Error creating reporting chaincode: %v", err)
	}

	if err := reportingSmartContract.Start(); err != nil {
		log.Panicf("Error starting reporting chaincode: %v", err)
	}
}
