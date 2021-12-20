package keeper

import (
	"github.com/line/lbm-sdk/codec"
	codectypes "github.com/line/lbm-sdk/codec/types"
	"github.com/line/lbm-sdk/store"
	sdk "github.com/line/lbm-sdk/types"
	accountkeeper "github.com/line/lbm-sdk/x/auth/keeper"
	accounttypes "github.com/line/lbm-sdk/x/auth/types"
	banktypes "github.com/line/lbm-sdk/x/bank/types"
	paramtypes "github.com/line/lbm-sdk/x/params/types"
	"github.com/line/ostracon/libs/log"
	ostproto "github.com/line/ostracon/proto/ostracon/types"
	"github.com/line/tm-db/v2/memdb"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func genAddress() sdk.AccAddress {
	b := make([]byte, 20)
	rand.Read(b)
	return sdk.BytesToAccAddress(b)
}

func setupKeeper(storeKey *sdk.KVStoreKey) BaseKeeper {
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	amino := codec.NewLegacyAmino()

	accountStoreKey := sdk.NewKVStoreKey(accounttypes.StoreKey)
	accountSubspace := paramtypes.NewSubspace(cdc, amino, accountStoreKey, accounttypes.ModuleName)
	accountKeeper := accountkeeper.NewAccountKeeper(cdc, accountStoreKey, accountSubspace, accounttypes.ProtoBaseAccount, nil)

	bankSubspace := paramtypes.NewSubspace(cdc, amino, storeKey, banktypes.StoreKey)
	return NewBaseKeeper(cdc, storeKey, accountKeeper, bankSubspace, nil)
}

func setupContext(t *testing.T, storeKey *sdk.KVStoreKey) sdk.Context {
	db := memdb.NewDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	return sdk.NewContext(stateStore, ostproto.Header{}, false, log.NewNopLogger())
}

func TestInactiveAddr(t *testing.T) {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	bankKeeper := setupKeeper(storeKey)
	ctx := setupContext(t, storeKey)

	addr := genAddress()

	require.Equal(t, 0, len(bankKeeper.inactiveAddrs))

	bankKeeper.addToInactiveAddr(ctx, addr)
	require.True(t, bankKeeper.isStoredInactiveAddr(ctx, addr))

	bankKeeper.addToInactiveAddr(ctx, addr)
	require.True(t, bankKeeper.isStoredInactiveAddr(ctx, addr))

	bankKeeper.deleteFromInactiveAddr(ctx, addr)
	require.False(t, bankKeeper.isStoredInactiveAddr(ctx, addr))

	addr2 := genAddress()
	require.False(t, bankKeeper.isStoredInactiveAddr(ctx, addr2))

	// expect no error
	bankKeeper.deleteFromInactiveAddr(ctx, addr2)

	// test loadAllInactiveAddrs
	bankKeeper.addToInactiveAddr(ctx, addr)
	bankKeeper.addToInactiveAddr(ctx, addr2)
	require.Equal(t, 0, len(bankKeeper.inactiveAddrs))
	bankKeeper.loadAllInactiveAddrs(ctx)
	require.Equal(t, 2, len(bankKeeper.inactiveAddrs))
}
