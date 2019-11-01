package ethapi

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef unsigned char u8;
// typedef int i32;
// typedef int i32ptr;
// extern void ethuseGas(void *context, long long  gas);
// extern void ethgetAddress(void *context, i32ptr resultOffset);
// extern void ethgetExternalBalance(void *context, i32ptr addressOffset, i32ptr resultOffset);
// extern i32 ethgetBlockHash(void *context, long long number, i32ptr resultOffset);
// extern i32 ethcall(void *context, long long gas, i32ptr addressOffset, i32ptr valueOffset, i32ptr dataOffset, i32 dataLength);
// extern i32 ethgetCallDataSize(void *context);
// extern void ethcallDataCopy(void *context, i32ptr resultsOffset, i32ptr dataOffset, i32 length);
// extern i32 ethcallCode(void *context, long long gas, i32ptr addressOffset, i32ptr valueOffset, i32ptr dataOffset, i32 dataLength);
// extern i32 ethcallDelegate(void *context, long long gas, i32ptr addressOffset, i32ptr dataOffset, i32 dataLength);
// extern i32 ethcallStatic(void *context, long long gas, i32ptr addressOffset, i32ptr dataOffset, i32 dataLength);
// extern void ethstorageStore(void *context, i32ptr pathOffset, i32ptr valueOffset);
// extern void ethstorageLoad(void *context, i32ptr pathOffset, i32ptr resultOffset);
// extern void ethgetCaller(void *context, i32ptr resultOffset);
// extern void ethgetCallValue(void *context, i32ptr resultOffset);
// extern void ethcodeCopy(void *context, i32ptr resultOffset, i32 codeOffset, i32 length);
// extern i32 ethgetCodeSize(void *context);
// extern void ethgetBlockCoinbase(void *context, i32ptr resultOffset);
// extern i32 ethcreate(void *context, i32ptr valueoffset, i32ptr dataOffset, i32 length, i32ptr resultsOffset);
// extern void ethgetBlockDifficulty(void *context, i32ptr resultOffset);
// extern void ethexternalCodeCopy(void *context, i32ptr addressOffset, i32ptr resultOffset, i32 codeOffset, i32 length);
// extern i32 ethgetExternalCodeSize(void *context, i32ptr addressOffset);
// extern long long ethgetGasLeft(void *context);
// extern long long ethgetBlockGasLimit(void *context);
// extern void ethgetTxGasPrice(void *context, i32ptr valueOffset);
// extern void ethlogTopics(void *context, i32ptr dataOffset, i32 length, i32 numberOftopics, i32ptr topic1, i32ptr topic2, i32ptr topic3, i32ptr topic4);
// extern long long ethgetBlockNumber(void *context);
// extern void ethgetTxOrigin(void *context, i32ptr resultOffset);
// extern void ethfinish(void *context, i32ptr dataOffset, i32 length);
// extern void ethrevert(void *context, i32ptr dataOffset, i32 length);
// extern i32 ethgetReturnDataSize(void *context);
// extern void ethreturnDataCopy(void *context, i32ptr resultOffset, i32 dataOffset, i32 length);
// extern void ethselfDestruct(void *context, i32ptr addressOffset);
// extern long long ethgetBlockTimestamp(void *context);
import "C"
import (
	"github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	"github.com/ElrondNetwork/go-ext-wasm/wasmer"
	"math/big"
	"unsafe"
)

func EthereumImports(imports *wasmer.Imports) (*wasmer.Imports, error) {
	imports = imports.Namespace("ethereum")

	imports, err := imports.Append("useGas", ethuseGas, C.ethuseGas)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getAddress", ethgetAddress, C.ethgetAddress)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getExternalBalance", ethgetExternalBalance, C.ethgetExternalBalance)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockHash", ethgetBlockHash, C.ethgetBlockHash)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("call", ethcall, C.ethcall)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("callDataCopy", ethcallDataCopy, C.ethcallDataCopy)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallDataSize", ethgetCallDataSize, C.ethgetCallDataSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("callCode", ethcallCode, C.ethcallCode)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("callDelegate", ethcallDelegate, C.ethcallDelegate)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("callStatic", ethcallStatic, C.ethcallStatic)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageStore", ethstorageStore, C.ethstorageStore)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("storageLoad", ethstorageLoad, C.ethstorageLoad)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCaller", ethgetCaller, C.ethgetCaller)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCallValue", ethgetCallValue, C.ethgetCallValue)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("codeCopy", ethcodeCopy, C.ethcodeCopy)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getCodeSize", ethgetCodeSize, C.ethgetCodeSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockCoinbase", ethgetBlockCoinbase, C.ethgetBlockCoinbase)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("create", ethcreate, C.ethcreate)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockDifficulty", ethgetBlockDifficulty, C.ethgetBlockDifficulty)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("externalCodeCopy", ethexternalCodeCopy, C.ethexternalCodeCopy)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getExternalCodeSize", ethgetExternalCodeSize, C.ethgetExternalCodeSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getGasLeft", ethgetGasLeft, C.ethgetGasLeft)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockGasLimit", ethgetBlockGasLimit, C.ethgetBlockGasLimit)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getTxGasPrice", ethgetTxGasPrice, C.ethgetTxGasPrice)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("logTopics", ethlogTopics, C.ethlogTopics)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockNumber", ethgetBlockNumber, C.ethgetBlockNumber)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getTxOrigin", ethgetTxOrigin, C.ethgetTxOrigin)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("finish", ethfinish, C.ethfinish)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("revert", ethrevert, C.ethrevert)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getReturnDataSize", ethgetReturnDataSize, C.ethgetReturnDataSize)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("returnDataCopy", ethreturnDataCopy, C.ethreturnDataCopy)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("selfDestruct", ethselfDestruct, C.ethselfDestruct)
	if err != nil {
		return nil, err
	}

	imports, err = imports.Append("getBlockTimestamp", ethgetBlockTimestamp, C.ethgetBlockTimestamp)
	if err != nil {
		return nil, err
	}

	return imports, nil
}

//export ethuseGas
func ethuseGas(context unsafe.Pointer, useGas int64) {
	instCtx := wasmer.IntoInstanceContext(context)
	ethContext := arwen.GetEthContext(instCtx.Data())

	ethContext.UseGas(useGas)
}

//export ethgetAddress
func ethgetAddress(context unsafe.Pointer, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, hostContext.GetSCAddress())
}

//export ethgetExternalBalance
func ethgetExternalBalance(context unsafe.Pointer, addressOffset int32, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	address := arwen.LoadBytes(instCtx.Memory(), addressOffset, arwen.AddressLen)
	balance := hostContext.GetBalance(address)

	err := arwen.StoreBytes(instCtx.Memory(), resultOffset, balance)
	if err != nil {
	}
}

//export ethgetBlockHash
func ethgetBlockHash(context unsafe.Pointer, number int64, resultOffset int32) int32 {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	hash := hostContext.BlockHash(number)
	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, hash)

	if len(hash) == 0 {
		return 0
	}
	return 1
}

//export ethcall
func ethcall(context unsafe.Pointer, gasLimit int64, addressOffset int32, valueOffset int32, dataOffset int32, dataLength int32) int32 {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	send := hostContext.GetVMInput().CallerAddr
	dest := arwen.LoadBytes(instCtx.Memory(), addressOffset, arwen.AddressLen)
	value := arwen.LoadBytes(instCtx.Memory(), valueOffset, arwen.BalanceLen)
	data := arwen.LoadBytes(instCtx.Memory(), dataOffset, dataLength)

	_, err := hostContext.Transfer(dest, send, big.NewInt(0).SetBytes(value), data, gasLimit)
	if err != nil {
		return 1
	}

	return 0
}

//export ethcallDataCopy
func ethcallDataCopy(context unsafe.Pointer, resultOffset int32, dataOffset int32, length int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	ethContext := arwen.GetEthContext(instCtx.Data())

	callData := ethContext.CallData()
	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, callData[dataOffset:dataOffset+length])
}

//export ethgetCallDataSize
func ethgetCallDataSize(context unsafe.Pointer) int32 {
	instCtx := wasmer.IntoInstanceContext(context)
	ethContext := arwen.GetEthContext(instCtx.Data())

	return int32(len(ethContext.CallData()))
}

//export ethstorageStore
func ethstorageStore(context unsafe.Pointer, pathOffset int32, valueOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	key := arwen.LoadBytes(instCtx.Memory(), pathOffset, arwen.HashLen)
	data := arwen.LoadBytes(instCtx.Memory(), valueOffset, arwen.HashLen)

	_ = hostContext.SetStorage(hostContext.GetSCAddress(), key, data)
}

//export ethstorageLoad
func ethstorageLoad(context unsafe.Pointer, pathOffset int32, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	key := arwen.LoadBytes(instCtx.Memory(), pathOffset, arwen.HashLen)
	data := hostContext.GetStorage(hostContext.GetSCAddress(), key)

	currInput := make([]byte, arwen.HashLen)
	copy(currInput[arwen.HashLen-len(data):], data)

	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, currInput)
}

//export ethgetCaller
func ethgetCaller(context unsafe.Pointer, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	caller := hostContext.GetVMInput().CallerAddr

	err := arwen.StoreBytes(instCtx.Memory(), resultOffset, caller)
	if err != nil {
	}
}

//export ethgetCallValue
func ethgetCallValue(context unsafe.Pointer, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	value := hostContext.GetVMInput().CallValue.Bytes()
	length := len(value)
	invBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		invBytes[length-i-1] = value[i]
	}

	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, invBytes)
}

//export ethcodeCopy
func ethcodeCopy(context unsafe.Pointer, resultOffset int32, codeOffset int32, length int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	scAddress := hostContext.GetSCAddress()
	code := hostContext.GetCode(scAddress)

	if int32(len(code)) > codeOffset+length {
		_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, code[codeOffset:codeOffset+length])
	}
}

//export ethgetCodeSize
func ethgetCodeSize(context unsafe.Pointer) int32 {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	return int32(hostContext.GetCodeSize(hostContext.GetSCAddress()))
}

//export ethexternalCodeCopy
func ethexternalCodeCopy(context unsafe.Pointer, addressOffset int32, resultOffset int32, codeOffset int32, length int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	dest := arwen.LoadBytes(instCtx.Memory(), addressOffset, arwen.AddressLen)
	code := hostContext.GetCode(dest)

	if int32(len(code)) > codeOffset+length {
		_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, code[codeOffset:codeOffset+length])
	}
}

//export ethgetExternalCodeSize
func ethgetExternalCodeSize(context unsafe.Pointer, addressOffset int32) int32 {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	dest := arwen.LoadBytes(instCtx.Memory(), addressOffset, arwen.AddressLen)

	return int32(hostContext.GetCodeSize(dest))
}

//export ethgetGasLeft
func ethgetGasLeft(context unsafe.Pointer) int64 {
	instCtx := wasmer.IntoInstanceContext(context)
	ethContext := arwen.GetEthContext(instCtx.Data())

	return ethContext.GasLeft()
}

//export ethgetBlockGasLimit
func ethgetBlockGasLimit(context unsafe.Pointer) int64 {
	instCtx := wasmer.IntoInstanceContext(context)
	ethContext := arwen.GetEthContext(instCtx.Data())

	return ethContext.BlockGasLimit()
}

//export ethgetTxGasPrice
func ethgetTxGasPrice(context unsafe.Pointer, valueOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	gasPrice := hostContext.GetVMInput().GasPrice

	gasU128 := make([]byte, 16)
	copy(gasU128[16-len(gasPrice.Bytes()):], gasPrice.Bytes())

	_ = arwen.StoreBytes(instCtx.Memory(), valueOffset, gasU128)
}

//export ethlogTopics
func ethlogTopics(context unsafe.Pointer, dataOffset int32, length int32, numberOfTopics int32, topic1 int32, topic2 int32, topic3 int32, topic4 int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	data := arwen.LoadBytes(instCtx.Memory(), dataOffset, length)

	topics := make([]int32, 0)
	topics = append(topics, topic1)
	topics = append(topics, topic2)
	topics = append(topics, topic3)
	topics = append(topics, topic4)

	topicsData := make([][]byte, numberOfTopics)
	for i := int32(0); i < numberOfTopics; i++ {
		topicsData[i] = arwen.LoadBytes(instCtx.Memory(), topics[i], arwen.HashLen)
	}

	hostContext.WriteLog(hostContext.GetSCAddress(), topicsData, data)
}

//export ethgetTxOrigin
func ethgetTxOrigin(context unsafe.Pointer, resultOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	_ = arwen.StoreBytes(instCtx.Memory(), resultOffset, hostContext.GetVMInput().CallerAddr)
}

//export ethfinish
func ethfinish(context unsafe.Pointer, resultOffset int32, length int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	data := arwen.LoadBytes(instCtx.Memory(), resultOffset, length)
	hostContext.Finish(data)
}

//export ethrevert
func ethrevert(context unsafe.Pointer, dataOffset int32, length int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	data := arwen.LoadBytes(instCtx.Memory(), dataOffset, length)
	hostContext.Finish(data)
	hostContext.SignalUserError()
}

//export ethselfDestruct
func ethselfDestruct(context unsafe.Pointer, addressOffset int32) {
	instCtx := wasmer.IntoInstanceContext(context)
	hostContext := arwen.GetEthContext(instCtx.Data())

	address := arwen.LoadBytes(instCtx.Memory(), addressOffset, arwen.HashLen)

	hostContext.SelfDestruct(address, hostContext.GetVMInput().CallerAddr)
}

//export ethgetBlockNumber
func ethgetBlockNumber(context unsafe.Pointer) int64 {
	//instCtx := wasmer.IntoInstanceContext(context)
	//hostContext := arwen.arwen.GetEthContext(instCtx.Data())

	//TODO: implement
	return 0 //hostContext.BlockNonce()
}

//export ethgetBlockTimestamp
func ethgetBlockTimestamp(context unsafe.Pointer) int64 {
	//instCtx := wasmer.IntoInstanceContext(context)
	//hostContext := arwen.arwen.GetEthContext(instCtx.Data())

	//return hostContext.BlockTimeStamp()
	return 0
}

//export ethgetReturnDataSize
func ethgetReturnDataSize(context unsafe.Pointer) int32 {
	//TODO: implement
	return 0
}

//export ethreturnDataCopy
func ethreturnDataCopy(context unsafe.Pointer, resultOffset int32, dataOffset int32, length int32) {
	//TODO: implement
}

//export ethgetBlockCoinbase
func ethgetBlockCoinbase(context unsafe.Pointer, resultOffset int32) {
	//TODO: implement
}

//export ethgetBlockDifficulty
func ethgetBlockDifficulty(context unsafe.Pointer, resultOffset int32) {
	//TODO: implement
}

//export ethcallCode
func ethcallCode(context unsafe.Pointer, gasLimit int64, addressOffset int32, valueOffset int32, dataOffset int32, dataLength int32) int32 {
	//TODO: implement
	return 0
}

//export ethcallDelegate
func ethcallDelegate(context unsafe.Pointer, gasLimit int64, addressOffset int32, dataOffset int32, dataLength int32) int32 {
	//TODO: implement
	return 0
}

//export ethcallStatic
func ethcallStatic(context unsafe.Pointer, gasLimit int64, addressOffset int32, dataOffset int32, dataLength int32) int32 {
	//TODO: implement
	return 0
}

//export ethcreate
func ethcreate(context unsafe.Pointer, valueOffset int32, dataOffset int32, length int32, resultOffset int32) int32 {
	//TODO: implement
	return 0
}
