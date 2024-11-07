package vm

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	GG16VInputUnpackErr = 1000001
	GG16VVerifyErr      = 1000002
	GG16WitnessErr      = 1000003
	GG16OtherErr        = 1000004

	GPVInputUnpackErr = 1000011
	GPVVerifyErr      = 1000012
	GPVWitnessErr     = 1000013
	GPVOther          = 1000014

	EVReadIndexErr                   = 1000022
	EVParseIndexErr                  = 1000023
	EVUnpackInputErr                 = 1000024
	EVParseInputHeightErr            = 1000025
	EVInputGreaterThanIndexHeightErr = 1000026
	EVParseInputHashErr              = 1000027
	EVReadSideChainDataErr           = 1000028
	EVGzipDecompressErr              = 1000030
	EVInvalidInput                   = 1000039
	EVOtherErr                       = 1000040
)

var (
	boolArg abi.Arguments
)

func init() {
	boolType, _ := abi.NewType("bool", "", nil)

	boolArg = abi.Arguments{
		{Type: boolType},
	}

}

func ErrCodeErr(code int) error {
	err := errors.New(fmt.Sprintf("%d", code))
	return err
}

func EncodeBool(b bool) []byte {
	r, _ := boolArg.Pack(b)
	return r
}
