package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
    
    "github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type SmartContract struct {
	contractapi.Contract
}

type Car struct {
	Model   string `json:"carModel"`
	ManufacturedOn string `json:"dateManufacture"`
	Status    string `json:"status"`
	Owner    string `json:"owner"`
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createCar" {
		return t.createCar(stub, args)
	} else if function == "getCar" { 
		return t.getCar(stub, args)
	} else if function == "transferCar" { 
		return t.transferCar(stub, args)
	} 
	
	return shim.Error("unknown function")
}

func (t *SmartContract) createCar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//Check sender
	
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

    val, err := cid.GetAttributeValue(stub, "role")
    
    if err != nil {
		return shim.Error(err.Error())
	}

    if val != "MANUFACTURER" {
        return "only manufacturer can create car", err
    }



    
	model := args[0]
	manufacturedOn := args[1]
	owner := "MANUFACTURER"
    status := "CREATED"
	
	rand.Seed(time.Now().UnixNano())
 
    carId := rand.Intn(10000000)
	
	Car := &Car{model, manufacturedOn, owner, carId}
	
	
	carJSONasBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error(err.Error())
	}

	 
	err = stub.PutState(carId, carJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}

func (t *SmartContract) getCar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carId = args[0]
	valAsbytes, err := stub.GetState(carId)  
	if err != nil {
		return shim.Error(err.Error())
	} else if valAsbytes == nil {
		return shim.Error("Car does not exists")
	}

	return shim.Success(valAsbytes)
}

func (t *SmartContract) transferCar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	 
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
    
    role, err := cid.GetAttributeValue(stub, "role")
    
    if err != nil {
		return shim.Error(err.Error())
	}

	if currentOwner == "MANUFACTURER" && (newOwner == "CUSTOMER" || newOwner == "MANUFACTURER") {
		return "This car cannot be transferred from MANUFACTURER to" + newOwner, err
	}
	if currentOwner == "DEALER" && (newOwner == "DEALER" || newOwner == "MANUFACTURER") {
		return "This car cannot be transferred from DEALER to" + newOwner, err
	}
	
	
	carId := args[0]
	currentOwner := args[1]
	newOwner := args[2]
	 
	valAsbytes, err := stub.GetState(c)  
	if err != nil {
		return shim.Error(err.Error())
	} else if valAsbytes == nil {
		return shim.Error("Car does not exists")
	}
	
	
	carToTransfer := Car{}
	err = json.Unmarshal(valAsbytes, &carToTransfer)
	if err != nil {
		return "", err
	}
	
	if currentOwner != carToTransfer.Owner {
		return "This car is not owned by" + currentOwner, err
	}
	
	carToTransfer.Owner = newOwner

    if role = "MANUFACTURER"  && newOwner == "DEALER" {
        carToTransfer.Status = "READY_FOR_SALE"
    }
    if role = "DELAER"  && newOwner == "CUSTOMER" {
        carToTransfer.Status = "SOLD"
    }
	
	carJSONasBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error(err.Error())
	}

	 
	err = stub.PutState(carId, carJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}