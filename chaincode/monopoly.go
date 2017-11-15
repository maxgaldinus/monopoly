/**
 * BLOCKCHAIN EXPERIENCE
 * Banco Imobiliário
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, formatting, and string manipulation
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
//	"bytes"
	"encoding/json"
	"fmt"
//	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Propriedade struct {
	Location string `json:"Location"`
	InitialValue string `json:"InitialValue"`
	Holder string `json:"Holder"`
}

type Wallet struct {
	Value string `json:"InitialValue"`
	Holder string `json:"Holder"`
}


type ReservaDeTransacao struct {
	Value string `json:"InitialValue"`
	Buyer string `json:"Buyer"`
	Seller string `json:"Seller"`
	Location string `json:"Location"`
	// Datetime
	// expired ?
}



/*
 * The Init method
 * called when the Smart Contract is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
          return shim.Success(nil)
}


/*
 * The Invoke method
 * called when an application requests to run the Smart Contract 
 * The app also specifies the specific smart contract function to call with args
 */
//func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

//func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    // Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	function = function
	args = args
	// Route to the appropriate handler function to interact with the ledger appropriately
/*
	if function == "consultarPropriedades" {
		return s.consultarPropriedades(stub, args)
	} else if function == "initLedger" {
		return s.initLedger(stub)
	} else if function == "transferirPropriedade" {
		return s.transferirPropriedade(stub, args)
	} else if function == "realizarPagamento" {
		return s.realizarPagamento(stub, args)
	}
*/
	//else if function == "consultarTodasPropriedades" {
	//	return s.consultarTodasPropriedades(stub)
	//}
	return shim.Error("Invalid Smart Contract function name.")
}



func (s *SmartContract) consultarPropriedades(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	propriedadeAsBytes, _ := stub.GetState(args[0])
	if propriedadeAsBytes == nil {
		return shim.Error("Could not locate property")
	}
	return shim.Success(propriedadeAsBytes)
}


func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) sc.Response {
	propriedades := []Propriedade{
		Propriedade{Location: "Jardim Botânico", InitialValue: "100000", Holder: "Banco"},
		Propriedade{Location: "Avenida Niemeyer", InitialValue: "75000", Holder: "Banco"},
	}
	i := 0
	for i < len(propriedades) {
		fmt.Println("i is ", i)
		propriedadeAsBytes, _ := json.Marshal(propriedades[i])
		stub.PutState(propriedades[i].Location, propriedadeAsBytes)
		fmt.Println("Added", propriedades[i])
		i = i + 1
	}
	return shim.Success(nil)
}

func (s *SmartContract) initWallets(stub shim.ChaincodeStubInterface) sc.Response {
	wallets := []Wallet{
		Wallet{Value: "100000", Holder: "Player 1"},
		Wallet{Value: "75000", Holder: "Player 2"},
		Wallet{Value: "999999999999", Holder: "Banco"},
	}
	i := 0
	for i < len(wallets) {
	fmt.Println("i is ", i)
	walletAsBytes, _ := json.Marshal(wallets[i])
	stub.PutState(wallets[i].Holder, walletAsBytes)
	fmt.Println("Added", wallets[i])
	i = i + 1
	}
	return shim.Success(nil)
}

func (s *SmartContract) transferirPropriedade(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	propriedadeAsBytes, _ := stub.GetState(args[0])
	if propriedadeAsBytes != nil {
		return shim.Error("Could not locate property")
	}
	propriedade := Propriedade{}
	json.Unmarshal(propriedadeAsBytes, &propriedade)

	propriedade.Holder = args[1]
	propriedadeAsBytes, _ = json.Marshal(propriedade)
	err := stub.PutState(args[0], propriedadeAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change property holder: %s", args[0]))
	}
	return shim.Success(nil)
}

/*
func (s *SmartContract) realizarPagamento(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	
	// pegar carteira ORIGEM
	walletFromAsBytes, _ := stub.GetState(args[0])
	if walletFromAsBytes != nil {
		return shim.Error("Could not locate source wallet")
	}
	walletFrom := Wallet{}
	json.Unmarshal(walletFromAsBytes, &walletFrom)
	
	// pegar carteira DESTINO
	walletToAsBytes, _ := stub.GetState(args[1])
	if walletToAsBytes != nil {
		return shim.Error("Could not locate target wallet")
	}
	walletTo := Wallet{}
	json.Unmarshal(walletToAsBytes, &walletTo)
	
	// transfere recursos
	Value := args[2]
	
	if walletTo.Value < Value {
		return shim.Error(fmt.Sprintf("Insufficient funds to transfer %s from %s to %s", args[2], args[0], args[1]))
	}
	
	/**
	walletFrom.Value := walletFrom.Value - Value
	walletTo.Value := walletTo.Value + Value

	// salva estado
	walletFromAsBytes, _ := json.Marshal(walletFrom)
	walletToAsBytes, _ := json.Marshal(walletTo)
	err := stub.PutState(walletFrom.Holder, walletFromAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to transfer value from %s to %s", args[0], args[1]))
	}
	err := stub.PutState(walletTo.Holder, walletToAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to transfer value from %s to %s", args[0], args[1]))
	}
	
	return shim.Success(nil)
}
*/

/*
 * main function
 * calls the Start function 
 * The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	} else {
                fmt.Printf("Success creating new Smart Contract")
        }
//	shim.Invoke()
}
