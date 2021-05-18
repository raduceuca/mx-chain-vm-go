package testcommon

import (
	"fmt"
	"math/big"
	"testing"
	"unicode"

	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
	"github.com/stretchr/testify/require"
)

// VMOutputVerifier holds the output to be verified
type VMOutputVerifier struct {
	VmOutput *vmcommon.VMOutput
	T        testing.TB
}

// NewVMOutputVerifier builds a new verifier
func NewVMOutputVerifier(t testing.TB, vmOutput *vmcommon.VMOutput, err error) *VMOutputVerifier {
	require.Nil(t, err)
	require.NotNil(t, vmOutput)

	return &VMOutputVerifier{
		VmOutput: vmOutput,
		T:        t,
	}
}

// Ok verifies if return code is vmcommon.Ok
func (v *VMOutputVerifier) Ok() *VMOutputVerifier {
	return v.ReturnCode(vmcommon.Ok)
}

// ReturnCode verifies if ReturnCode of output is the same as the provided one
func (v *VMOutputVerifier) ReturnCode(code vmcommon.ReturnCode) *VMOutputVerifier {
	require.Equal(v.T, code, v.VmOutput.ReturnCode, "ReturnCode")
	return v
}

// ReturnMessage verifies if ReturnMessage of output is the same as the provided one
func (v *VMOutputVerifier) ReturnMessage(message string) *VMOutputVerifier {
	require.Equal(v.T, message, v.VmOutput.ReturnMessage, "ReturnMessage")
	return v
}

// NoMsg verifies that ReturnMessage is empty
func (v *VMOutputVerifier) NoMsg() *VMOutputVerifier {
	require.Equal(v.T, "", v.VmOutput.ReturnMessage, "ReturnMessage")
	return v
}

// Msg verifies if ReturnMessage of output is the same as the provided one
func (v *VMOutputVerifier) Msg(message string) *VMOutputVerifier {
	require.Equal(v.T, message, v.VmOutput.ReturnMessage, "ReturnMessage")
	return v
}

// GasUsed verifies if GasUsed of the specified account is the same as the provided one
func (v *VMOutputVerifier) GasUsed(address []byte, gas uint64) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("GasUsed", address)
	require.NotNil(v.T, account, errMsg)
	require.Equal(v.T, int(gas), int(account.GasUsed), errMsg)
	return v
}

// GasRemaining verifies if GasRemaining of the specified account is the same as the provided one
func (v *VMOutputVerifier) GasRemaining(gas uint64) *VMOutputVerifier {
	require.Equal(v.T, int(gas), int(v.VmOutput.GasRemaining), "GasRemaining")
	return v
}

// Balance verifies if Balance of the specified account is the same as the provided one
func (v *VMOutputVerifier) Balance(address []byte, balance int64) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("Balance", address)
	require.NotNil(v.T, account, errMsg)
	require.NotNil(v.T, account.Balance, errMsg)
	require.Equal(v.T, balance, account.Balance.Int64(), errMsg)
	return v
}

// BalanceDelta verifies if BalanceDelta of the specified account is the same as the provided one
func (v *VMOutputVerifier) BalanceDelta(address []byte, balanceDelta int64) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("BalanceDelta", address)
	require.NotNil(v.T, account, errMsg)
	require.NotNil(v.T, account.BalanceDelta, errMsg)
	require.Equal(v.T, balanceDelta, account.BalanceDelta.Int64(), errMsg)
	return v
}

// Nonce verifies if Nonce of the specified account is the same as the provided one
func (v *VMOutputVerifier) Nonce(address []byte, nonce uint64) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("Nonce", address)
	require.NotNil(v.T, account, errMsg)
	require.Equal(v.T, nonce, account.Nonce, errMsg)
	return v
}

// Code verifies if Code of the specified account is the same as the provided one
func (v *VMOutputVerifier) Code(address []byte, code []byte) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("Code", address)
	require.NotNil(v.T, account, errMsg)
	require.Equal(v.T, code, account.Code, errMsg)
	return v
}

// CodeMetadata if CodeMetadata of the specified account is the same as the provided one
func (v *VMOutputVerifier) CodeMetadata(address []byte, codeMetadata []byte) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("CodeMetadata", address)
	require.NotNil(v.T, account, errMsg)
	require.Equal(v.T, codeMetadata, account.CodeMetadata, errMsg)
	return v
}

// CodeDeployerAddress if CodeDeployerAddress of the specified account is the same as the provided one
func (v *VMOutputVerifier) CodeDeployerAddress(address []byte, codeDeployerAddress []byte) *VMOutputVerifier {
	account := v.VmOutput.OutputAccounts[string(address)]
	errMsg := getErrorForAccount("CodeDeployerAddress", address)
	require.NotNil(v.T, account, errMsg)
	require.Equal(v.T, codeDeployerAddress, account.CodeDeployerAddress, errMsg)
	return v
}

// ReturnData verifies if ReturnData is the same as the provided one
func (v *VMOutputVerifier) ReturnData(returnData ...[]byte) *VMOutputVerifier {
	require.Equal(v.T, len(returnData), len(v.VmOutput.ReturnData), "ReturnData")
	for idx := range v.VmOutput.ReturnData {
		require.Equal(v.T, returnData[idx], v.VmOutput.ReturnData[idx], "ReturnData")
	}
	return v
}

// StoreEntry holds the data for a storage assertion
type StoreEntry struct {
	address []byte
	key     []byte
	value   []byte
}

// CreateStoreEntry creates the data for a storage assertion
func CreateStoreEntry(address []byte) *StoreEntry {
	return &StoreEntry{address: address}
}

// WithKey sets the key for a storage assertion
func (storeEntry *StoreEntry) WithKey(key []byte) *StoreEntry {
	storeEntry.key = key
	return storeEntry
}

// WithValue sets the value for a storage assertion
func (storeEntry *StoreEntry) WithValue(value []byte) StoreEntry {
	storeEntry.value = value
	return *storeEntry
}

// Storage verifies if StorageUpdate(s) for the speficied accounts are the same as the provided ones
func (v *VMOutputVerifier) Storage(returnData ...StoreEntry) *VMOutputVerifier {

	storage := make(map[string]map[string]vmcommon.StorageUpdate)

	for _, storeEntry := range returnData {
		account := string(storeEntry.address)
		accountStorageMap, exists := storage[account]
		if !exists {
			accountStorageMap = make(map[string]vmcommon.StorageUpdate)
			storage[account] = accountStorageMap
		}
		accountStorageMap[string(storeEntry.key)] = vmcommon.StorageUpdate{Offset: storeEntry.key, Data: storeEntry.value}
	}

	for _, outputAccount := range v.VmOutput.OutputAccounts {
		accountStorageMap := storage[string(outputAccount.Address)]
		require.Equal(v.T, len(accountStorageMap), len(outputAccount.StorageUpdates), "Storage")
		for key, value := range accountStorageMap {
			require.Equal(v.T, value, *outputAccount.StorageUpdates[key], "Storage")
		}
		delete(storage, string(outputAccount.Address))
	}
	require.Equal(v.T, 0, len(storage), "Storage")

	return v
}

// TransferEntry holds the data for an output transfer assertion
type TransferEntry struct {
	vmcommon.OutputTransfer
	address []byte
}

// CreateTransferEntry creates the data for an output transfer assertion
func CreateTransferEntry(senderAddress []byte, receiverAddress []byte) *TransferEntry {
	return &TransferEntry{
		OutputTransfer: vmcommon.OutputTransfer{SenderAddress: senderAddress},
		address:        receiverAddress,
	}
}

// WithData create sets the data for an output transfer assertion
func (transferEntry *TransferEntry) WithData(data []byte) *TransferEntry {
	transferEntry.Data = data
	return transferEntry
}

// WithValue create sets the value for an output transfer assertion
func (transferEntry *TransferEntry) WithValue(value *big.Int) TransferEntry {
	transferEntry.Value = value
	return *transferEntry
}

// Transfers verifies if OutputTransfer(s) for the speficied accounts are the same as the provided ones
func (v *VMOutputVerifier) Transfers(transfers ...TransferEntry) *VMOutputVerifier {

	transfersMap := make(map[string][]vmcommon.OutputTransfer)

	for _, transferEntry := range transfers {
		account := string(transferEntry.address)
		accountTransfers, exists := transfersMap[account]
		if !exists {
			accountTransfers = make([]vmcommon.OutputTransfer, 0)
		}
		transfersMap[account] = append(accountTransfers, transferEntry.OutputTransfer)
	}

	for _, outputAccount := range v.VmOutput.OutputAccounts {
		transfersForAccount := transfersMap[string(outputAccount.Address)]
		require.Equal(v.T, len(transfersForAccount), len(outputAccount.OutputTransfers), "Transfers")
		for idx := range transfersForAccount {
			require.Equal(v.T, transfersForAccount[idx], outputAccount.OutputTransfers[idx], "Transfers")
		}
		delete(transfersMap, string(outputAccount.Address))
	}
	require.Equal(v.T, 0, len(transfersMap), "Transfers")

	return v
}

func getErrorForAccount(field string, address []byte) string {
	return fmt.Sprintf("%s %s", field, humanReadable(address))
}

func humanReadable(address []byte) string {
	var result []byte
	for _, c := range address {
		if unicode.IsPrint(rune(c)) {
			result = append(result, c)
		}
	}
	return string(result)
}
