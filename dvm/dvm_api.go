/*
 *    This file is part of DVM library.
 *
 *    The DVM library is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    The DVM library is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with the DVM library.  If not, see <http://www.gnu.org/licenses/>.
 */
package dvm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/dispatchlabs/disgo/commons/crypto"
	commonTypes "github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/dvm/ethereum/abi"
	ethTypes "github.com/dispatchlabs/disgo/dvm/ethereum/types"
)

// DeploySmartContract -
func (dvm *DVMService) DeploySmartContract(tx *commonTypes.Transaction) (*DVMResult, error) {
	utils.Info(fmt.Sprintf("DVMServices-DeploySmartContract: %s", tx))

	// Load the TRIE state for [FROM:TO] combo
	stateHelper, err := NewVMStateHelper(crypto.GetAddressBytes(tx.From), crypto.GetAddressBytes(tx.To))
	if err != nil {
		return nil, err
	}

	if err := dvm.applyTransaction(tx, stateHelper); err != nil {
		utils.Error(err)
		return nil, err
	}

	// Commit the change
	_, err = stateHelper.Commit() // returns `hashOfTrieRootNode` but don't use below
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	// Get info about the TX
	bytes, _ := hex.DecodeString(tx.Hash)
	receipt, err := dvm.getReceipt(bytes)

	return &DVMResult{
		From:               crypto.GetAddressBytes(tx.From),
		To:                 receipt.ContractAddress,
		ABI:                "",
		StorageState:       stateHelper,
		ContractExecError:  nil,
		ContractExecResult: nil,

		Divvy:               _defaultDivvy,
		Status:              receipt.Status,
		HertzCost:           receipt.GasUsed,
		CumulativeHertzUsed: receipt.CumulativeGasUsed,
		Bloom:               receipt.Bloom,
		Logs:                receipt.Logs,
	}, nil
}

// ExecuteSmartContract -
func (dvm *DVMService) ExecuteSmartContract(tx *commonTypes.Transaction) (*DVMResult, error) {
	utils.Info(fmt.Sprintf("DVMServices-ExecuteSmartContract: %s", tx))

	// Load the TRIE state for [FROM:TO] combo
	stateHelper, err := NewVMStateHelper(crypto.GetAddressBytes(tx.From), crypto.GetAddressBytes(tx.To))
	if err != nil {
		return nil, err
	}

	// Prepare the method params from ABI
	fromHexAsByteArray, _ := hex.DecodeString(tx.Abi)
	abiAsString := string(fromHexAsByteArray)
	jsonABI, err := abi.JSON(strings.NewReader(abiAsString))
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	callData, err := jsonABI.Pack(tx.Method, tx.Params...)
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	// Execte the Smart-Contract Method
	toAsBytes := crypto.GetAddressBytes(tx.To)
	callMsg := ethTypes.NewMessage(
		crypto.GetAddressBytes(tx.From),
		&toAsBytes,
		0, // nonce
		_defaultValue,
		_defaultGas.Uint64(),
		_defaultGasPrice,
		callData,
		false,
	)

	execResult, execError := dvm.call(callMsg, stateHelper)
	if execError != nil {
		utils.Error(execError)
		return nil, execError
	}

	// Commit the change
	_, err = stateHelper.Commit() // returns `hashOfTrieRootNode` but don't use below
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	// Return the state of the storage and the execution result
	return &DVMResult{
		From:               crypto.GetAddressBytes(tx.From),
		To:                 crypto.GetAddressBytes(tx.To),
		ABI:                tx.Abi,
		StorageState:       stateHelper,
		ContractExecError:  execError,
		ContractExecResult: execResult,

		// Status:              receipt.Status,
		// HertzCost:           receipt.GasUsed,
		// CumulativeHertzUsed: receipt.CumulativeGasUsed,
		// Bloom:               receipt.Bloom,
		// Logs:                receipt.Logs,
	}, nil
}
