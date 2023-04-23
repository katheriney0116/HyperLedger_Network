package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"project_contract/contract"
	"fmt"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&contract.UserContract{}, &contract.ProjectContract{})
	fmt.Println("Main is called")
	if err != nil {
		panic(err)
	}

	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
