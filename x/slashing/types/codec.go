package types

import (
	"github.com/Finschia/finschia-rdk/codec"
	"github.com/Finschia/finschia-rdk/codec/legacy"
	"github.com/Finschia/finschia-rdk/codec/types"
	cryptocodec "github.com/Finschia/finschia-rdk/crypto/codec"
	sdk "github.com/Finschia/finschia-rdk/types"
	"github.com/Finschia/finschia-rdk/types/msgservice"
	authzcodec "github.com/Finschia/finschia-rdk/x/authz/codec"
	fdncodec "github.com/Finschia/finschia-rdk/x/foundation/codec"
	govcodec "github.com/Finschia/finschia-rdk/x/gov/codec"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgUnjail{}, "cosmos-sdk/MsgUnjail")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnjail{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(fdncodec.Amino)
}
