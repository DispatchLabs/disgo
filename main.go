/*
 *    This file is part of Disgo library.
 *
 *    The Disgo library is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    The Disgo library is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with the Disgo library.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"github.com/dispatchlabs/disgo/bootstrap"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/commons/types"
	"time"
	"encoding/hex"
)

func main() {
	utils.InitMainPackagePath()
	utils.InitializeLogger()

	privateKey := "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	from := "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"
	now := utils.ToMilliSeconds(time.Now())
	code := "608060405234801561001057600080fd5b506040805190810160405280600d81526020017f61616161616161616161616161000000000000000000000000000000000000008152506000908051906020019061005c9291906100f7565b50600060016000018190555060006001800160006101000a81548160ff02191690831515021790555060018060010160016101000a81548160ff021916908360ff1602179055506040805190810160405280600b81526020017f6262626262626262626262000000000000000000000000000000000000000000815250600160020190805190602001906100f19291906100f7565b5061019c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013857805160ff1916838001178555610166565b82800160010185558215610166579182015b8281111561016557825182559160200191906001019061014a565b5b5090506101739190610177565b5090565b61019991905b8082111561019557600081600090555060010161017d565b5090565b90565b610664806101ab6000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806333e538e91461007d57806334e45f531461010d5780633a458b1f1461017657806378d8866e1461022557806379af6473146102b5578063cb69e300146102cc575b600080fd5b34801561008957600080fd5b50610092610335565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100d25780820151818401526020810190506100b7565b50505050905090810190601f1680156100ff5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011957600080fd5b50610174600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506103d7565b005b34801561018257600080fd5b5061018b6103f4565b60405180858152602001841515151581526020018360ff1660ff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101e75780820151818401526020810190506101cc565b50505050905090810190601f1680156102145780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561023157600080fd5b5061023a6104c4565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561027a57808201518184015260208101905061025f565b50505050905090810190601f1680156102a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102c157600080fd5b506102ca610562565b005b3480156102d857600080fd5b50610333600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610579565b005b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103cd5780601f106103a2576101008083540402835291602001916103cd565b820191906000526020600020905b8154815290600101906020018083116103b057829003601f168201915b5050505050905090565b80600160020190805190602001906103f0929190610593565b5050565b60018060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104ba5780601f1061048f576101008083540402835291602001916104ba565b820191906000526020600020905b81548152906001019060200180831161049d57829003601f168201915b5050505050905084565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561055a5780601f1061052f5761010080835404028352916020019161055a565b820191906000526020600020905b81548152906001019060200180831161053d57829003601f168201915b505050505081565b600160000160008154809291906001019190505550565b806000908051906020019061058f929190610593565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105d457805160ff1916838001178555610602565b82800160010185558215610602579182015b828111156106015782518255916020019190600101906105e6565b5b50905061060f9190610613565b5090565b61063591905b80821115610631576000816000905550600101610619565b5090565b905600a165627a7a7230582026a289af0b033267e3fc13869c446b0552e2c62b9b3fa46bc626be2c683528680029"
 	abi :=  `[{
    "constant": true,
    "inputs": [],
    "name": "getVar5",
    "outputs": [
      {
        "name": "",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "value",
        "type": "string"
      }
    ],
    "name": "setVar6Var4",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "var6",
    "outputs": [
      {
        "name": "var1",
        "type": "uint256"
      },
      {
        "name": "var2",
        "type": "bool"
      },
      {
        "name": "var3",
        "type": "uint8"
      },
      {
        "name": "var4",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "var5",
    "outputs": [
      {
        "name": "",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "incVar6Var1",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "value",
        "type": "string"
      }
    ],
    "name": "setVar5",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`


	t, err := types.NewDeployContractTransaction(privateKey, from, code, now)
	if err != nil {
		utils.Error(err)
	}
	s := t.String()
	utils.Info(s)


	to := "95d7b0f7dd388e37b2c8eebd2ced116817c8bf4e"
	var method = "setVar5"
	var params = make([]interface{}, 1)
	params[0] = "5555"

	t, err = types.NewExecuteContractTransaction(privateKey, from, to, hex.EncodeToString([]byte(abi)), method, params, now)
	if err != nil {
		utils.Error(err)
	}
	s = t.String()
	utils.Info(s)

	method = "getVar5"
	params = make([]interface{}, 0)

	t, err = types.NewExecuteContractTransaction(privateKey, from, to, hex.EncodeToString([]byte(abi)), method, params, now)
	if err != nil {
		utils.Error(err)
	}
	s = t.String()
	utils.Info(s)

	if !t.Verify() {
		utils.Error("FOOK ME")
	}

	server := bootstrap.NewServer()
	server.Go()
}
