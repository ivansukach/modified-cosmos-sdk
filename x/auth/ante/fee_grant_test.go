package ante_test

import (
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	"github.com/ivansukach/modified-cosmos-sdk/codec/types"
	"github.com/ivansukach/modified-cosmos-sdk/testutil/testdata"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/ante"
	"github.com/ivansukach/modified-cosmos-sdk/x/auth/tx"
)

type setFeeGranter interface {
	SetFeeGranter(feeGranter sdk.AccAddress)
}

func (suite *AnteTestSuite) TestRejectFeeGranter() {
	suite.SetupTest(true) // setup
	txConfig := tx.NewTxConfig(codec.NewProtoCodec(types.NewInterfaceRegistry()), tx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()
	d := ante.NewRejectFeeGranterDecorator()
	antehandler := sdk.ChainAnteDecorators(d)

	_, err := antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().NoError(err)

	setGranterTx := txBuilder.(setFeeGranter)
	_, _, addr := testdata.KeyTestPubAddr()
	setGranterTx.SetFeeGranter(addr)

	_, err = antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().Error(err)
}
