package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	linktype "github.com/line/link/types"
	"github.com/line/link/x/token/internal/types"
)

var ChildExists = []byte{1}

func (k Keeper) Attach(ctx sdk.Context, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error {
	if err := k.attach(ctx, from, symbol, toTokenID, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAttachToken,
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyToTokenID, toTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) AttachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error {
	if !k.IsApproved(ctx, proxy, from, symbol) {
		return types.ErrCollectionNotApproved(types.DefaultCodespace, proxy.String(), from.String(), symbol)
	}

	if err := k.attach(ctx, from, symbol, toTokenID, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAttachFrom,
			sdk.NewAttribute(types.AttribyteKeyProxy, proxy.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyToTokenID, toTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) attach(ctx sdk.Context, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error {
	store := ctx.KVStore(k.storeKey)

	if toTokenID == tokenID {
		return types.ErrCannotAttachToItself(types.DefaultCodespace, symbol+tokenID)
	}

	token, err := k.GetCNFT(ctx, symbol, tokenID)
	if err != nil {
		return err
	}

	if !from.Equals(token.GetOwner()) {
		return types.ErrTokenNotOwnedBy(types.DefaultCodespace, token.GetDenom(), from)
	}

	// verify token should be a root
	childToParentKey := types.TokenChildToParentKey(token)
	if store.Has(childToParentKey) {
		return types.ErrTokenAlreadyAChild(types.DefaultCodespace, token.GetDenom())
	}

	// verify no circulation(toToken must not be a descendant of token)
	rootOfToToken, err := k.RootOf(ctx, symbol, toTokenID)
	if err != nil {
		return err
	}
	toToken, err := k.GetCNFT(ctx, symbol, toTokenID)
	if err != nil {
		return err
	}
	if rootOfToToken.GetTokenID() == tokenID {
		return types.ErrCannotAttachToADescendant(types.DefaultCodespace, token.GetDenom(), toToken.GetDenom())
	}

	parentToChildKey := types.TokenParentToChildKey(toToken)
	childrenStore := prefix.NewStore(store, parentToChildKey)
	parentToChildSubKey := types.TokenParentToChildSubKey(token)
	if childrenStore.Has(parentToChildSubKey) {
		panic("token is already a child of some other")
	}

	if !from.Equals(toToken.GetOwner()) {
		if err := k.moveCNFToken(ctx, store, from, toToken.GetOwner(), token); err != nil {
			return err
		}
	}

	store.Set(childToParentKey, k.mustEncodeTokenDenom(toToken.GetDenom()))
	childrenStore.Set(parentToChildSubKey, ChildExists)

	return nil
}

func (k Keeper) Detach(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, symbol string, tokenID string) sdk.Error {
	if err := k.detach(ctx, from, to, symbol, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDetachToken,
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeyTo, to.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})
	return nil
}

func (k Keeper) DetachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, to sdk.AccAddress, symbol string, tokenID string) sdk.Error {
	if !k.IsApproved(ctx, proxy, from, symbol) {
		return types.ErrCollectionNotApproved(types.DefaultCodespace, proxy.String(), from.String(), symbol)
	}

	if err := k.detach(ctx, from, to, symbol, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDetachFrom,
			sdk.NewAttribute(types.AttribyteKeyProxy, proxy.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeyTo, to.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) detach(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, symbol string, tokenID string) sdk.Error {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetCNFT(ctx, symbol, tokenID)
	if err != nil {
		return err
	}

	if !from.Equals(token.GetOwner()) {
		return types.ErrTokenNotOwnedBy(types.DefaultCodespace, token.GetDenom(), from)
	}

	childToParentKey := types.TokenChildToParentKey(token)
	if !store.Has(childToParentKey) {
		return types.ErrTokenNotAChild(types.DefaultCodespace, token.GetDenom())
	}

	bz := store.Get(childToParentKey)
	parentTokenDenom := k.mustDecodeTokenDenom(ctx, bz)
	ticker, suffix, tokenID := linktype.ParseDenom(parentTokenDenom)

	parentToken, err := k.GetCNFT(ctx, ticker+suffix, tokenID)
	if err != nil {
		return err
	}

	parentToChildKey := types.TokenParentToChildKey(parentToken)
	childrenStore := prefix.NewStore(store, parentToChildKey)
	parentToChildSubKey := types.TokenParentToChildSubKey(token)
	if !childrenStore.Has(parentToChildSubKey) {
		panic("token is not a child of some other")
	}

	if !from.Equals(to) {
		if err := k.moveCNFToken(ctx, store, from, to, token); err != nil {
			return err
		}
	}
	store.Delete(childToParentKey)
	childrenStore.Delete(parentToChildSubKey)

	return nil
}

func (k Keeper) RootOf(ctx sdk.Context, symbol string, tokenID string) (types.CollectiveNFT, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetCNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}

	for childToParentKey := types.TokenChildToParentKey(token); store.Has(childToParentKey); {
		bz := store.Get(childToParentKey)
		parentTokenDenom := k.mustDecodeTokenDenom(ctx, bz)
		ticker, suffix, tokenID := linktype.ParseDenom(parentTokenDenom)

		token, err = k.GetCNFT(ctx, ticker+suffix, tokenID)
		if err != nil {
			return nil, err
		}
		childToParentKey = types.TokenChildToParentKey(token)
	}

	return token, nil
}

func (k Keeper) ParentOf(ctx sdk.Context, symbol string, tokenID string) (types.CollectiveNFT, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetCNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}
	childToParentKey := types.TokenChildToParentKey(token)
	if store.Has(childToParentKey) {
		bz := store.Get(childToParentKey)
		parentTokenDenom := k.mustDecodeTokenDenom(ctx, bz)
		ticker, suffix, tokenID := linktype.ParseDenom(parentTokenDenom)

		return k.GetCNFT(ctx, ticker+suffix, tokenID)
	}
	return nil, nil
}

func (k Keeper) ChildrenOf(ctx sdk.Context, symbol string, tokenID string) (types.Tokens, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetCNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}

	result := make([]types.Token, 0)
	parentToChildKey := types.TokenParentToChildKey(token)
	iter := sdk.KVStorePrefixIterator(store, parentToChildKey)
	defer iter.Close()
	for {
		if !iter.Valid() {
			break
		}
		tokenDenom := types.ParentToChildSubKeyToToken(parentToChildKey, iter.Key())
		ticker, suffix, tokenID := linktype.ParseDenom(tokenDenom)

		childToken, err := k.GetCNFT(ctx, ticker+suffix, tokenID)
		if err != nil {
			panic("child token does not exist")
		}
		result = append(result, childToken)
		iter.Next()
	}
	return result, nil
}

func (k Keeper) GetCNFT(ctx sdk.Context, symbol, tokenID string) (types.CollectiveNFT, sdk.Error) {
	token, err := k.GetToken(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}
	tokenNFT, ok := token.(*types.BaseCollectiveNFT)
	if !ok {
		return nil, types.ErrTokenNotCNFT(types.DefaultCodespace, token.GetDenom())
	}
	return tokenNFT, nil
}

func (k Keeper) mustEncodeTokenDenom(denom string) []byte {
	return k.cdc.MustMarshalBinaryBare(denom)
}

func (k Keeper) mustDecodeTokenDenom(ctx sdk.Context, tokenByte []byte) string {
	var denom string
	k.cdc.MustUnmarshalBinaryBare(tokenByte, &denom)
	return denom
}
