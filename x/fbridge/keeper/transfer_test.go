package keeper

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/Finschia/finschia-sdk/types"
	authtypes "github.com/Finschia/finschia-sdk/x/auth/types"
	"github.com/Finschia/finschia-sdk/x/fbridge/testutil"
	"github.com/Finschia/finschia-sdk/x/fbridge/types"
)

func TestHandleBridgeTransfer(t *testing.T) {
	key, ctx, encCfg, authKeeper, bankKeeper := testutil.PrepareFbridgeTest(t)

	sender := sdk.AccAddress("test")
	amt := sdk.NewInt(1000000)
	denom := "stake"
	token := sdk.Coins{sdk.Coin{Denom: denom, Amount: amt}}

	bridge := authtypes.NewEmptyModuleAccount("bridge")
	authKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(bridge.GetAddress()).AnyTimes()
	authKeeper.EXPECT().GetModuleAccount(ctx, types.ModuleName).Return(bridge).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, token).Return(nil)
	bankKeeper.EXPECT().BurnCoins(ctx, types.ModuleName, token).Return(nil)

	k := NewKeeper(encCfg.Codec, key, authKeeper, bankKeeper, denom, "gov")
	targetSeq := uint64(2)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, targetSeq)
	ctx.KVStore(key).Set(types.KeyNextSeqSend, bz)

	handledSeq, err := k.handleBridgeTransfer(ctx, sender, amt)
	require.NoError(t, err)
	require.Equal(t, targetSeq, handledSeq)
	afterSeq := k.GetNextSequence(ctx)
	require.Equal(t, targetSeq+1, afterSeq)
}

func TestIsValidEthereumAddress(t *testing.T) {
	tcs := map[string]struct {
		valid   bool
		address string
	}{
		"valid": {
			valid:   true,
			address: "0xf7bAc63fc7CEaCf0589F25454Ecf5C2ce904997c",
		},
		"invalid": {
			valid:   false,
			address: "0xf7bAc63fc7CEaCf0589F25454Ecf5C2ce905997c",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.valid, IsValidEthereumAddress(tc.address))
		})
	}
}
