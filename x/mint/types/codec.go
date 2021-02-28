package types

import (
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	cryptocodec "github.com/ivansukach/modified-cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
