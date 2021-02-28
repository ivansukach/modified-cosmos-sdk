package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ivansukach/modified-cosmos-sdk/codec"
	cryptoAmino "github.com/ivansukach/modified-cosmos-sdk/crypto/codec"
	"github.com/ivansukach/modified-cosmos-sdk/testutil/testdata"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
