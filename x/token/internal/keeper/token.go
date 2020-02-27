package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/token/internal/types"
)

func (k Keeper) GetToken(ctx sdk.Context, symbol string) (types.Token, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenSymbolKey(symbol))
	if bz == nil {
		return nil, types.ErrTokenNotExist(types.DefaultCodespace, symbol)
	}
	return k.mustDecodeToken(bz), nil
}

func (k Keeper) SetToken(ctx sdk.Context, token types.Token) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	tokenKey := types.TokenSymbolKey(token.GetSymbol())
	if store.Has(tokenKey) {
		return types.ErrTokenExist(types.DefaultCodespace, token.GetSymbol())
	}
	store.Set(tokenKey, k.cdc.MustMarshalBinaryBare(token))

	k.setSupply(ctx, types.DefaultSupply(token.GetSymbol()))

	return nil
}

func (k Keeper) UpdateToken(ctx sdk.Context, token types.Token) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	tokenKey := types.TokenSymbolKey(token.GetSymbol())
	if !store.Has(tokenKey) {
		return types.ErrTokenNotExist(types.DefaultCodespace, token.GetSymbol())
	}
	store.Set(tokenKey, k.cdc.MustMarshalBinaryBare(token))
	return nil
}

func (k Keeper) GetAllTokens(ctx sdk.Context) (tokens types.Tokens) {
	appendToken := func(token types.Token) (stop bool) {
		tokens = append(tokens, token)
		return false
	}
	k.iterateTokens(ctx, "", appendToken)
	return tokens
}

func (k Keeper) iterateTokens(ctx sdk.Context, prefix string, process func(types.Token) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.TokenSymbolKey(prefix))
	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}
		val := iter.Value()
		token := k.mustDecodeToken(val)
		if process(token) {
			return
		}
		iter.Next()
	}
}

func (k Keeper) mustDecodeToken(tokenByte []byte) (token types.Token) {
	err := k.cdc.UnmarshalBinaryBare(tokenByte, &token)
	if err != nil {
		panic(err)
	}
	return token
}
