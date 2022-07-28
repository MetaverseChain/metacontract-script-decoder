package script

import (
	"bytes"
	"encoding/binary"
)

//decodeMvcFT
// <op_pushdata> + <type specific data> + <proto header> + <data_len(4 bytes)> + <version(1 bytes)>
// <proto header> = <proto_version(4 bytes)> + <proto_type(4 bytes)> + <'metacontract'(12 bytes)>
// <token type specific data> = <name(40 bytes)> + <symbol(20 bytes)> + <decimal(1 bytes)> + <address(20 bytes)> + <token amount(8 bytes)> + <genesisHash(20 bytes)> + <sensibleID(36 bytes)>
func decodeMvcFT(scriptLen int, pkScript []byte, txo *TxoData) bool {
	//<name(40 bytes)> +
	//<symbol(20 bytes)> +
	//<decimal(1 bytes)> +
	//<address(20 bytes)> +
	//<token amount(8 bytes)> +
	//<genesisHash(20 bytes)> +
	//<sensibleID(36 bytes)> +
	//<proto_version(4 bytes)> +
	//<proto_type(4 bytes)> +
	//<'metacontract'(12 bytes)>+
	//<data_len(4 bytes)> +
	//<version(1 bytes)>
	dataLen := 40 + 20 + 1 + 20 + 8 + 20 + 36 + 4 + 4 + 12 + 4 + 1
	protoVersionLen := 4
	protoTypeLen := 4
	genesisIdLen := 20
	addressLen := 20
	sensibleIdLen := 36
	decimalLen := 1
	symbolLen := 20
	nameLen := 40
	amountLen := 8

	if !(pkScript[scriptLen-dataLen-1-1-1] == OP_RETURN &&
		pkScript[scriptLen-dataLen-1-1] == 0x4c &&
		pkScript[scriptLen-dataLen-1] == 0xaa) {
		// error ft
		return false

	}
	protoTypeOffset := scriptLen - 17 - protoTypeLen
	sensibleOffset := protoTypeOffset - protoVersionLen - sensibleIdLen
	genesisOffset := sensibleOffset - genesisIdLen
	amountOffset := genesisOffset - amountLen
	addressOffset := amountOffset - addressLen
	decimalOffset := addressOffset - decimalLen
	symbolOffset := decimalOffset - symbolLen
	nameOffset := symbolOffset - nameLen

	txo.CodeType = CodeType_FT

	ft := &FTData{
		Decimal: uint8(pkScript[decimalOffset]),
		Symbol:  string(bytes.TrimRight(pkScript[symbolOffset:symbolOffset+symbolLen], "\x00")),
		Name:    string(bytes.TrimRight(pkScript[nameOffset:nameOffset+nameLen], "\x00")),
		Amount:  binary.LittleEndian.Uint64(pkScript[amountOffset : amountOffset+amountLen]),
	}
	txo.FT = ft

	txo.HasAddress = true
	copy(txo.AddressPkh[:], pkScript[addressOffset:addressOffset+addressLen])

	copy(txo.CodeHash[:], GetHash160(pkScript[:scriptLen-dataLen]))
	ft.SensibleId = make([]byte, genesisIdLen)
	copy(ft.SensibleId, pkScript[genesisOffset:genesisOffset+genesisIdLen])

	txo.GenesisIdLen = uint8(genesisIdLen)
	copy(txo.GenesisId[:], ft.SensibleId)
	return true
}