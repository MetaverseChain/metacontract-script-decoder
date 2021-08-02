package script

import (
	"bytes"
	"encoding/binary"
)

func decodeFT(scriptLen int, scriptPk []byte, txo *TxoData) bool {
	dataLen := 0
	protoVersionLen := 0
	genesisIdLen := 0
	sensibleIdLen := 0
	useTokenIdHash := false

	if scriptPk[scriptLen-76-76-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-76-76-1-1] == 0x4c &&
		scriptPk[scriptLen-76-76-1] == 152 {
		// v6
		// <type specific data> + <proto header>
		// <proto header> = <version(4 bytes)> + <type(4 bytes)> + <'sensible'(8 bytes)>
		// <token type specific data> = <token_name (20 bytes)> + <token_symbol (10 bytes)> + <is_genesis(1 byte)> + <decimailNum(1 byte)> + <address(20 bytes)> + <token amount(8 bytes)> + <genesisHash(20 bytes)> + <rabinPubKeyHashArrayHash(20 bytes)> + <genesisId(36 bytes)>
		protoVersionLen = 4
		genesisIdLen = 76
		sensibleIdLen = 36
		dataLen = 1 + 1 + 76 + genesisIdLen // 0x4c + pushdata + data + genesisId
		useTokenIdHash = true

	} else if scriptPk[scriptLen-72-76-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-72-76-1-1] == 0x4c &&
		scriptPk[scriptLen-72-76-1] == 148 {
		// v5
		// <type specific data> + <proto header>
		// <proto header> = <type(4 bytes)> + <'sensible'(8 bytes)>
		// <token type specific data> = <token_name (20 bytes)> + <token_symbol (10 bytes)> + <is_genesis(1 byte)> + <decimailNum(1 byte)> + <address(20 bytes)> + <token amount(8 bytes)> + <genesisHash(20 bytes)> + <rabinPubKeyHashArrayHash(20 bytes)> + <genesisId(36 bytes)>
		genesisIdLen = 76
		sensibleIdLen = 36
		dataLen = 1 + 1 + 1 + 72 + genesisIdLen // opreturn + 0x4c + pushdata + data + genesisId
		useTokenIdHash = true
	} else if scriptPk[scriptLen-72-36-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-72-36-1-1] == 0x4c &&
		scriptPk[scriptLen-72-36-1] == 108 {
		// v4
		// v1 ~ v4
		// <type specific data> + <proto header>
		// <proto header> = <type(4 bytes)> + <'sensible'(8 bytes)>
		// <token type specific data> = <token_name (20 bytes)> + <token_symbol (10 bytes)> + <is_genesis(1 byte)> + <decimailNum(1 byte)> + <address(20 bytes)> + <token amount(8 bytes)> + <genesisId(x bytes)>
		genesisIdLen = 36
		dataLen = 1 + 1 + 1 + 72 + genesisIdLen
	} else if scriptPk[scriptLen-72-20-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-72-20-1-1] == 0x4c &&
		scriptPk[scriptLen-72-20-1] == 92 {
		// ft v3
		genesisIdLen = 20
		dataLen = 1 + 1 + 1 + 72 + genesisIdLen
	} else if scriptPk[scriptLen-50-36-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-50-36-1-1] == 0x4c &&
		scriptPk[scriptLen-50-36-1] == 86 {
		// ft v2
		genesisIdLen = 36
		dataLen = 1 + 1 + 1 + 50 + genesisIdLen
	} else if scriptPk[scriptLen-92-20-1-1-1] == OP_RETURN &&
		scriptPk[scriptLen-92-20-1-1] == 0x4c &&
		scriptPk[scriptLen-92-20-1] == 112 {
		// ft v1
		genesisIdLen = 20
		dataLen = 1 + 1 + 1 + 92 + genesisIdLen
	} else {
		// error ft
		return false
	}

	protoTypeOffset := scriptLen - 8 - 4
	sensibleOffset := protoTypeOffset - protoVersionLen - sensibleIdLen

	genesisOffset := protoTypeOffset - protoVersionLen - genesisIdLen
	amountOffset := genesisOffset - 8
	addressOffset := amountOffset - 20
	decimalOffset := addressOffset - 1
	symbolOffset := decimalOffset - 1 - 10
	nameOffset := symbolOffset - 20

	txo.CodeType = CodeType_FT
	txo.Decimal = uint8(scriptPk[decimalOffset])
	txo.Symbol = string(bytes.TrimRight(scriptPk[symbolOffset:symbolOffset+10], "\x00"))
	txo.Name = string(bytes.TrimRight(scriptPk[nameOffset:nameOffset+20], "\x00"))

	txo.Amount = binary.LittleEndian.Uint64(scriptPk[amountOffset : amountOffset+8])

	txo.AddressPkh = make([]byte, 20)
	copy(txo.AddressPkh, scriptPk[addressOffset:addressOffset+20])

	txo.CodeHash = GetHash160(scriptPk[:scriptLen-dataLen])
	if useTokenIdHash {
		txo.SensibleId = make([]byte, sensibleIdLen)
		copy(txo.SensibleId, scriptPk[sensibleOffset:sensibleOffset+sensibleIdLen])

		// GenesisId is tokenIdHash
		txo.GenesisId = GetHash160(scriptPk[genesisOffset : genesisOffset+genesisIdLen])
	} else {
		txo.GenesisId = make([]byte, genesisIdLen)
		copy(txo.GenesisId, scriptPk[genesisOffset:genesisOffset+genesisIdLen])

		txo.SensibleId = txo.GenesisId
	}
	return true
}
