package types

import (
	codectypes "github.com/ivansukach/modified-cosmos-sdk/codec/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/ibc/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
