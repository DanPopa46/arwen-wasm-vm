package arwen

import (
	"fmt"
	"math/big"

	"sync"
	"unsafe"

	"github.com/ElrondNetwork/go-ext-wasm/wasmer"
)

const AddressLen = 32
const AddressLenEth = 20
const HashLen = 32
const BalanceLen = 32
const InitFunctionName = "init"
const InitFunctionNameEth = "solidity.ctor"

var (
	vmContextCounter uint8
	vmContextMap     map[uint8]VMContext
	vmContextMapMu   sync.Mutex
)

func AddHostContext(ctx VMContext) int {
	vmContextMapMu.Lock()
	id := vmContextCounter
	vmContextCounter++
	if vmContextMap == nil {
		vmContextMap = make(map[uint8]VMContext)
	}
	vmContextMap[id] = ctx
	vmContextMapMu.Unlock()
	return int(id)
}

func RemoveHostContext(idx int) {
	vmContextMapMu.Lock()
	delete(vmContextMap, uint8(idx))
	vmContextMapMu.Unlock()
}

func GetEthContext(pointer unsafe.Pointer) EthContext {
	var idx = *(*int)(pointer)

	vmContextMapMu.Lock()
	ctx := vmContextMap[uint8(idx)]
	vmContextMapMu.Unlock()

	return ctx.EthContext()
}

func GetErdContext(pointer unsafe.Pointer) HostContext {
	var idx = *(*int)(pointer)

	vmContextMapMu.Lock()
	ctx := vmContextMap[uint8(idx)]
	vmContextMapMu.Unlock()

	return ctx.CoreContext()
}

func GetBigIntContext(pointer unsafe.Pointer) BigIntContext {
	var idx = *(*int)(pointer)

	vmContextMapMu.Lock()
	ctx := vmContextMap[uint8(idx)]
	vmContextMapMu.Unlock()

	return ctx.BigInContext()
}

func GetCryptoContext(pointer unsafe.Pointer) CryptoContext {
	var idx = *(*int)(pointer)

	vmContextMapMu.Lock()
	ctx := vmContextMap[uint8(idx)]
	vmContextMapMu.Unlock()

	return ctx.CryptoContext()
}

func LoadBytes(from *wasmer.Memory, offset int32, length int32) []byte {
	result := make([]byte, length)
	if from.Length() < uint32(offset+length) {
		copy(result, from.Data()[offset:])
		return result
	}

	copy(result, from.Data()[offset:offset+length])
	return result
}

func StoreBytes(to *wasmer.Memory, offset int32, data []byte) error {
	length := int32(len(data))

	if to.Length() < uint32(offset+length) {
		err := to.Grow(1)
		if err != nil {
			return err
		}
	}

	var memoryData = to.Data()
	copy(memoryData[offset:offset+length], data)

	return nil
}

func ConvertReturnValue(wasmValue wasmer.Value) *big.Int {
	switch wasmValue.GetType() {
	case wasmer.TypeVoid:
		return big.NewInt(0)
	case wasmer.TypeI32:
		return big.NewInt(wasmValue.ToI64())
	case wasmer.TypeI64:
		return big.NewInt(wasmValue.ToI64())
	}

	panic("unsupported return type")
}

func GuardedMakeByteSlice2D(length int32) ([][]byte, error) {
	if length < 0 {
		return nil, fmt.Errorf("GuardedMakeByteSlice2D: negative length (%d)", length)
	}

	result := make([][]byte, length)
	return result, nil
}

func GuardedLoadBytes(from *wasmer.Memory, offset int32, length int32) ([]byte, error) {
	memoryView := from.Data()
	memoryLength := from.Length()
	requestedEnd := uint32(offset + length)
	isOffsetTooSmall := offset < 0
	isOffsetTooLarge := uint32(offset) > memoryLength
	isRequestedEndTooLarge := requestedEnd > memoryLength
	isLengthNegative := length < 0

	if isOffsetTooSmall || isOffsetTooLarge {
		return nil, fmt.Errorf("GuardedLoadBytes: bad bounds")
	}

	if isLengthNegative {
		return nil, fmt.Errorf("GuardedLoadBytes: negative length")
	}

	result := make([]byte, length)

	if isRequestedEndTooLarge {
		copy(result, memoryView[offset:])
	} else {
		copy(result, memoryView[offset:requestedEnd])
	}

	return result, nil
}

func GuardedStoreBytes(to *wasmer.Memory, offset int32, data []byte) error {
	memoryView := to.Data()
	memoryLength := to.Length()
	dataLength := int32(len(data))
	requestedEnd := uint32(offset + dataLength)
	isOffsetTooSmall := offset < 0
	isNewPageNecessary := requestedEnd > memoryLength

	if isOffsetTooSmall {
		return fmt.Errorf("GuardedStoreBytes: bad lower bounds")
	}

	if isNewPageNecessary {
		err := to.Grow(1)
		if err != nil {
			return err
		}

		memoryView = to.Data()
		memoryLength = to.Length()
	}

	isRequestedEndTooLarge := requestedEnd > memoryLength

	if isRequestedEndTooLarge {
		return fmt.Errorf("GuardedStoreBytes: bad upper bounds")
	}

	copy(memoryView[offset:requestedEnd], data)
	return nil
}
