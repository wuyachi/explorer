package utils

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
	"math/big"
)

func AbandonPrecision(value *big.Int) uint64 {
	precision := big.NewInt(1000000000)
	xx := new(big.Int).Div(value, precision)
	return xx.Uint64()
}

func AmountWithoutPrecision(amount uint64) string {
	aaa := decimal.NewFromInt(int64(amount))
	precision := decimal.NewFromInt(1000000000)
	newAmount := aaa.Div(precision)
	return newAmount.String()
}

func Hash2Bool(hash common.Hash) bool {
	if common.EmptyHash(hash) {
		return false
	}
	data := binary.BigEndian.Uint16(hash[30:])
	if data == 0 {
		return false
	}
	return true
}
