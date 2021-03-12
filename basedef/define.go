package basedef

import (
	"fmt"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
)

var (
	MAX_DATA_LEN = 4096
)

const (
	CONTRACT_TYPE_PLT = iota
	CONTRACT_TYPE_NFT
)

const (
	TRANSACTION_STATUS_FAILED = iota
	TRANSACTION_STATUS_SUCCESS
)

const (
	TRANSACTION_TYPE_TRANSFER = iota
	TRANSACTION__TYPE_CONTRACTS
)

const (
	STATISTIC_DO = 1
	STATISTIC_NOT = 2
)

func FromatPercent(percent float64) string {
	decimal.DivisionPrecision = 2
	aaa := decimal.NewFromFloat(percent * 10000)
	aaa = aaa.Div(decimal.NewFromInt(100))
	bbb := fmt.Sprintf("%s%s", aaa.String(), "%")
	return bbb
}

