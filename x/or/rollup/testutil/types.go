package testutil

import (
	"github.com/Finschia/finschia-rdk/crypto/keys/ed25519"
	sdk "github.com/Finschia/finschia-rdk/types"
)

func AccAddressString() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// AccAddress returns a sample account address
func AccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}
