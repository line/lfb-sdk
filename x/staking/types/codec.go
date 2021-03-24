package types

import (
	"github.com/line/lbm-sdk/v2/codec"
	"github.com/line/lbm-sdk/v2/codec/types"
	cryptocodec "github.com/line/lbm-sdk/v2/crypto/codec"
	sdk "github.com/line/lbm-sdk/v2/types"
	"github.com/line/lbm-sdk/v2/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/staking interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateValidator{}, "lbm-sdk/MsgCreateValidator", nil)
	cdc.RegisterConcrete(&MsgEditValidator{}, "lbm-sdk/MsgEditValidator", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "lbm-sdk/MsgDelegate", nil)
	cdc.RegisterConcrete(&MsgUndelegate{}, "lbm-sdk/MsgUndelegate", nil)
	cdc.RegisterConcrete(&MsgBeginRedelegate{}, "lbm-sdk/MsgBeginRedelegate", nil)
}

// RegisterInterfaces registers the x/staking interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
