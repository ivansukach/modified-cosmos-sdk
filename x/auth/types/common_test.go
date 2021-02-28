package types_test

import (
	"github.com/ivansukach/modified-cosmos-sdk/simapp"
)

var (
	app                   = simapp.Setup(false)
	appCodec, legacyAmino = simapp.MakeCodecs()
)
