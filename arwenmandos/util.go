package arwenmandos

import (
	"encoding/hex"
	"fmt"
	"math/big"

	mj "github.com/ElrondNetwork/arwen-wasm-vm/mandos-go/json/model"
	oj "github.com/ElrondNetwork/arwen-wasm-vm/mandos-go/orderedjson"
	worldhook "github.com/ElrondNetwork/arwen-wasm-vm/mock/world"
	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
)

func convertAccount(testAcct *mj.Account) *worldhook.Account {
	storage := make(map[string][]byte)
	for _, stkvp := range testAcct.Storage {
		key := string(stkvp.Key.Value)
		storage[key] = stkvp.Value.Value
	}

	if len(testAcct.Address.Value) != 32 {
		panic("bad test: account address should be 32 bytes long")
	}

	convertedESDTData := make(map[string]*worldhook.ESDTData)
	for _, mandosESDTData := range testAcct.ESDTData {
		convertedESDTData[string(mandosESDTData.TokenName.Value)] = &worldhook.ESDTData{
			Balance:      mandosESDTData.Balance.Value,
			BalanceDelta: big.NewInt(0),
			Frozen:       mandosESDTData.Frozen.Value > 0,
		}
	}

	return &worldhook.Account{
		Address:         testAcct.Address.Value,
		Nonce:           testAcct.Nonce.Value,
		Balance:         big.NewInt(0).Set(testAcct.Balance.Value),
		Storage:         storage,
		Code:            []byte(testAcct.Code.Value),
		AsyncCallData:   testAcct.AsyncCallData,
		ESDTData:        convertedESDTData,
		ShardID:         uint32(testAcct.Shard.Value),
		IsSmartContract: len(testAcct.Code.Value) > 0,
		CodeMetadata: (&vmcommon.CodeMetadata{
			Payable:     true,
			Upgradeable: true,
			Readable:    true,
		}).ToBytes(), // TODO: add explicit fields in mandos json
	}
}

func convertNewAddressMocks(testNAMs []*mj.NewAddressMock) []*worldhook.NewAddressMock {
	var result []*worldhook.NewAddressMock
	for _, testNAM := range testNAMs {
		result = append(result, &worldhook.NewAddressMock{
			CreatorAddress: testNAM.CreatorAddress.Value,
			CreatorNonce:   testNAM.CreatorNonce.Value,
			NewAddress:     testNAM.NewAddress.Value,
		})
	}
	return result
}

func convertBlockInfo(testBlockInfo *mj.BlockInfo) *worldhook.BlockInfo {
	if testBlockInfo == nil {
		return nil
	}

	var randomsSeed [48]byte
	if testBlockInfo.BlockRandomSeed != nil {
		copy(randomsSeed[:], testBlockInfo.BlockRandomSeed.Value)
	}

	result := &worldhook.BlockInfo{
		BlockTimestamp: testBlockInfo.BlockTimestamp.Value,
		BlockNonce:     testBlockInfo.BlockNonce.Value,
		BlockRound:     testBlockInfo.BlockRound.Value,
		BlockEpoch:     uint32(testBlockInfo.BlockEpoch.Value),
		RandomSeed:     &randomsSeed,
	}

	return result
}

// this is a small hack, so we can reuse mandos's JSON printing in error messages
func convertLogToTestFormat(outputLog *vmcommon.LogEntry) *mj.LogEntry {
	testLog := mj.LogEntry{
		Address:    mj.JSONCheckBytesReconstructed(outputLog.Address),
		Identifier: mj.JSONCheckBytesReconstructed(outputLog.Identifier),
		Data:       mj.JSONCheckBytesReconstructed(outputLog.Data),
		Topics:     make([]mj.JSONCheckBytes, len(outputLog.Topics)),
	}
	for i, topic := range outputLog.Topics {
		testLog.Topics[i] = mj.JSONCheckBytesReconstructed(topic)
	}

	return &testLog
}

func bigIntPretty(i *big.Int) string {
	return fmt.Sprintf("0x%x (%d)", i, i)
}

func byteArrayPretty(bytes []byte) string {
	if len(bytes) == 0 {
		return "[]"
	}

	if canInterpretAsString(bytes) {
		return fmt.Sprintf("0x%s (``%s)", hex.EncodeToString(bytes), string(bytes))
	}

	asInt := big.NewInt(0).SetBytes(bytes)
	return fmt.Sprintf("0x%s (%d)", hex.EncodeToString(bytes), asInt)
}

func canInterpretAsString(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}
	for _, b := range bytes {
		if b < 32 || b > 126 {
			return false
		}
	}
	return true
}

func generateTxHash(txIndex string) []byte {
	txIndexBytes := []byte(txIndex)
	if len(txIndexBytes) > 32 {
		return txIndexBytes[:32]
	}
	for i := len(txIndexBytes); i < 32; i++ {
		txIndexBytes = append(txIndexBytes, '.')
	}
	return txIndexBytes
}

// JSONCheckBytesString formats a list of JSONCheckBytes for printing to console.
func checkBytesListPretty(jcbs []mj.JSONCheckBytes) string {
	str := "["
	for i, jcb := range jcbs {
		if i > 0 {
			str += ", "
		}

		str += "\"" + oj.JSONString(jcb.Original) + "\""
	}
	return str + "]"
}
