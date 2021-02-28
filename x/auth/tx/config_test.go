package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ivansukach/modified-cosmos-sdk/codec"
	codectypes "github.com/ivansukach/modified-cosmos-sdk/codec/types"
	"github.com/ivansukach/modified-cosmos-sdk/std"
	"github.com/ivansukach/modified-cosmos-sdk/testutil/testdata"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
