package std

import (
	"github.com/ivansukach/modified-cosmos-sdk/codec"
	"github.com/ivansukach/modified-cosmos-sdk/codec/types"
	cryptocodec "github.com/ivansukach/modified-cosmos-sdk/crypto/codec"
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	txtypes "github.com/ivansukach/modified-cosmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
