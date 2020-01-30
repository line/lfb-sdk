package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	linktype "github.com/line/link/types"
	"github.com/line/link/x/token/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"testing"
)

const (
	defaultName     = "description"
	defaultSymbol   = "linktkn"
	defaultTokenURI = "token-uri"
	defaultDecimals = 6
	defaultAmount   = 1000
)

func TestGetAllTokens(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	require.NotNil(t, ak.GetAccount(ctx, addr1))
	require.Equal(t, uint64(0), ak.GetAccount(ctx, addr1).GetAccountNumber())

	t.Log("issue a token")
	{
		token := types.NewFT(defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultDecimals), true)
		require.NoError(t, keeper.IssueFT(ctx, token, sdk.NewInt(defaultAmount), addr1))
		tokens := keeper.GetAllTokens(ctx)
		require.Equal(t, defaultName, tokens[0].GetName())
		require.Equal(t, defaultSymbol, tokens[0].GetSymbol())
		require.Equal(t, int64(defaultDecimals), tokens[0].(types.FT).GetDecimals().Int64())
		require.Equal(t, true, tokens[0].(types.FT).GetMintable())
		require.Equal(t, int64(defaultAmount), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol).Int64())
	}
	t.Log("issue tokens and get tokens")
	{
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, defaultSymbol+"1", defaultTokenURI+"1", sdk.NewInt(defaultDecimals), true), sdk.NewInt(100), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, defaultSymbol+"2", defaultTokenURI+"2", sdk.NewInt(defaultDecimals), true), sdk.NewInt(200), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, defaultSymbol+"3", defaultTokenURI+"3", sdk.NewInt(defaultDecimals), true), sdk.NewInt(300), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, defaultSymbol+"4", defaultTokenURI+"4", sdk.NewInt(defaultDecimals), true), sdk.NewInt(400), addr1))
		tokens := keeper.GetAllTokens(ctx)
		{
			require.Equal(t, defaultName, tokens[0].GetName())
			require.Equal(t, defaultSymbol, tokens[0].GetSymbol())
			require.Equal(t, true, tokens[0].(types.FT).GetMintable())
			require.Equal(t, defaultTokenURI, tokens[0].GetTokenURI())
		}
		{
			require.Equal(t, defaultSymbol+"1", tokens[1].GetSymbol())
			require.Equal(t, defaultSymbol+"2", tokens[2].GetSymbol())
			require.Equal(t, defaultSymbol+"3", tokens[3].GetSymbol())
			require.Equal(t, defaultSymbol+"4", tokens[4].GetSymbol())
		}

		{
			require.Equal(t, defaultTokenURI+"1", tokens[1].GetTokenURI())
			require.Equal(t, defaultTokenURI+"2", tokens[2].GetTokenURI())
			require.Equal(t, defaultTokenURI+"3", tokens[3].GetTokenURI())
			require.Equal(t, defaultTokenURI+"4", tokens[4].GetTokenURI())
		}

		{
			require.Equal(t, int64(100), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol+"1").Int64())
			require.Equal(t, int64(200), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol+"2").Int64())
			require.Equal(t, int64(300), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol+"3").Int64())
			require.Equal(t, int64(400), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol+"4").Int64())
		}
	}
}

func TestIssueTokenAndSendTokens(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak, bk := input.Cdc, input.Ctx, input.Keeper, input.Ak, input.Bk

	// Register account 1
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr1))
		require.Equal(t, uint64(0), ak.GetAccount(ctx, addr1).GetAccountNumber())
	}

	// Register account 2
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr2))
		require.Equal(t, uint64(1), ak.GetAccount(ctx, addr2).GetAccountNumber())
	}

	t.Log("Issue a token")
	{
		setupToken := types.NewFT(defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultDecimals), true)
		require.NoError(t, keeper.IssueFT(ctx, setupToken, sdk.NewInt(900), addr1))

		token, err := keeper.GetToken(ctx, defaultSymbol, "")
		require.NoError(t, err)
		require.Equal(t, defaultName, token.GetName())
		require.Equal(t, defaultSymbol, token.GetSymbol())
		require.Equal(t, defaultTokenURI, token.GetTokenURI())
		require.Equal(t, int64(defaultDecimals), token.(types.FT).GetDecimals().Int64())
		require.Equal(t, true, token.(types.FT).GetMintable())
		require.Equal(t, int64(900), keeper.accountKeeper.GetAccount(ctx, addr1).GetCoins().AmountOf(defaultSymbol).Int64())

		require.NoError(t, keeper.MintTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(99))), addr1))

		require.Equal(t, int64(999), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(999), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
	}
	t.Log("Issue a token again FAIL")
	{
		token := types.NewFT(defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultDecimals), true)
		require.Error(t, keeper.IssueFT(ctx, token, sdk.NewInt(900), addr1))
	}

	t.Log("Transfer Token")
	{
		require.NoError(t, bk.SendCoins(ctx, addr1, addr2, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(100)))))
		require.Equal(t, int64(899), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(100), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(999), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Transfer Token again")
	{
		require.NoError(t, bk.SendCoins(ctx, addr1, addr2, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(100)))))
		require.Equal(t, int64(799), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(200), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(999), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Mint Token")
	{
		require.NoError(t, keeper.MintTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(100))), addr1))
		require.Equal(t, int64(899), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(200), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1099), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Burn Token")
	{
		require.NoError(t, keeper.BurnTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(100))), addr1))
		require.Equal(t, int64(799), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(200), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(999), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Burn Token again amount > has --> fail")
	{
		require.Error(t, keeper.BurnTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(800))), addr1))
		require.Equal(t, int64(799), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(200), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(999), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

}

func TestIssueNFTAndSendTokens(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak, bk := input.Cdc, input.Ctx, input.Keeper, input.Ak, input.Bk

	// Register account 1
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr1))
		require.Equal(t, uint64(0), ak.GetAccount(ctx, addr1).GetAccountNumber())
	}

	// Register account 2
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr2))
		require.Equal(t, uint64(1), ak.GetAccount(ctx, addr2).GetAccountNumber())
	}

	t.Log("Issue a nft")
	{
		token := types.NewNFT(defaultName, defaultSymbol, defaultTokenURI, addr1)
		require.NoError(t, keeper.IssueNFT(ctx, token, addr1))
	}
	t.Log("Issue a nft again FAIL")
	{
		token := types.NewNFT(defaultName, defaultSymbol, defaultTokenURI, addr1)
		require.Error(t, keeper.IssueNFT(ctx, token, addr1))
	}

	t.Log("Get the token and check")
	{
		token, err := keeper.GetToken(ctx, defaultSymbol, "")
		require.NoError(t, err)
		require.Equal(t, defaultName, token.GetName())
		require.Equal(t, defaultSymbol, token.GetSymbol())
		require.Equal(t, defaultTokenURI, token.(types.NFT).GetTokenURI())
		require.Equal(t, int64(1), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
	}
	t.Log("Mint token -> fail. it is nft")
	{
		require.Error(t, keeper.MintTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(99))), addr1))
		require.Equal(t, int64(1), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
	}

	t.Log("Send insufficient")
	{
		require.Error(t, bk.SendCoins(ctx, addr1, addr2, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(10)))))
		require.Equal(t, int64(1), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(0), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Send from addr 1 to addr 2")
	{
		require.NoError(t, bk.SendCoins(ctx, addr1, addr2, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(1)))))
		require.Equal(t, int64(0), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Burn from account 1 -> fail.")
	{
		require.Error(t, keeper.BurnTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(1))), addr1))
		require.Equal(t, int64(0), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(1), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}

	t.Log("Burn from account 2 ( the owner)")
	{
		require.NoError(t, keeper.BurnTokens(ctx, sdk.NewCoins(sdk.NewCoin(defaultSymbol, sdk.NewInt(1))), addr2))
		require.Equal(t, int64(0), bk.GetCoins(ctx, addr1).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(0), bk.GetCoins(ctx, addr2).AmountOf(defaultSymbol).Int64())
		require.Equal(t, int64(0), keeper.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(defaultSymbol).Int64())
	}
}

func TestCollectionAndPermission(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak, _ := input.Cdc, input.Ctx, input.Keeper, input.Ak, input.Bk

	const (
		resource01 = "reso01"
		resource02 = "reso02"
	)

	// Register account 1
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr1))
		require.Equal(t, uint64(0), ak.GetAccount(ctx, addr1).GetAccountNumber())
	}

	// Register account 2
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
		require.NotNil(t, ak.GetAccount(ctx, addr2))
		require.Equal(t, uint64(1), ak.GetAccount(ctx, addr2).GetAccountNumber())
	}
	issuePerm := types.NewIssuePermission(resource01)
	{
		require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(resource01, "name"), addr1))
		require.True(t, keeper.HasPermission(ctx, addr1, issuePerm))
		require.Error(t, keeper.CreateCollection(ctx, types.NewCollection(resource01, "name"), addr1))
		collection, err := keeper.GetCollection(ctx, resource01)
		require.NoError(t, err)
		require.Equal(t, resource01, collection.GetSymbol())

		{
			require.NoError(t, keeper.IssueFT(ctx, types.NewCollectiveFT(collection, defaultName, "00000001", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
			require.NoError(t, keeper.MintCollectionTokens(ctx, linktype.NewCoinWithTokenIDs(linktype.NewCoinWithTokenID(resource01, "00000001", sdk.NewInt(defaultAmount))), addr1))
			supply, err := keeper.GetSupply(ctx, resource01, "00000001")
			require.NoError(t, err)
			require.Equal(t, int64(defaultAmount+defaultAmount), supply.Int64())

			collection, err := keeper.GetCollection(ctx, resource01)
			require.NoError(t, err)
			require.NoError(t, keeper.IssueFT(ctx, types.NewCollectiveFT(collection, defaultName, "", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
			token, err := keeper.GetToken(ctx, resource01, "00000002")
			require.NoError(t, err)
			require.Equal(t, resource01, token.GetSymbol())
			require.Equal(t, "00000002", token.GetTokenID())

		}
		{
			collection, err := keeper.GetCollection(ctx, resource01)
			require.NoError(t, err)
			require.NoError(t, keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, "a0000001", defaultTokenURI, addr1), addr1))

			token, err := keeper.GetToken(ctx, resource01, "a0000001")
			require.NoError(t, err)
			require.Equal(t, resource01, token.GetSymbol())
			require.Equal(t, "a0000001", token.GetTokenID())

			collection, err = keeper.GetCollection(ctx, resource01)
			require.NoError(t, err)
			require.NoError(t, keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, "a000", defaultTokenURI, addr1), addr1))
			token, err = keeper.GetToken(ctx, resource01, "a0000002")
			require.NoError(t, err)
			require.Equal(t, resource01, token.GetSymbol())
			require.Equal(t, "a0000002", token.GetTokenID())

			count, err := keeper.GetNFTCount(ctx, resource01, "a0000")
			require.NoError(t, err)
			require.Equal(t, int64(2), count.Int64())

			collection, err = keeper.GetCollection(ctx, resource01)
			require.NoError(t, err)
			require.NoError(t, keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, "", defaultTokenURI, addr1), addr1))
			token, err = keeper.GetToken(ctx, resource01, "a0010000")
			require.NoError(t, err)
			require.Equal(t, resource01, token.GetSymbol())
			require.Equal(t, "a0010000", token.GetTokenID())
		}

	}
	{
		require.NoError(t, keeper.GrantPermission(ctx, addr1, addr2, issuePerm))
		require.True(t, keeper.HasPermission(ctx, addr1, issuePerm))
		require.True(t, keeper.HasPermission(ctx, addr2, issuePerm))
	}

	issuePerm2 := types.NewIssuePermission(resource02)
	{
		require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(resource02, "name"), addr1))
		require.True(t, keeper.HasPermission(ctx, addr1, issuePerm2))
		require.Error(t, keeper.CreateCollection(ctx, types.NewCollection(resource02, "name"), addr1))
		collection, err := keeper.GetCollection(ctx, resource02)
		require.NoError(t, err)
		require.Equal(t, resource02, collection.GetSymbol())
	}
	{
		collections := keeper.GetAllCollections(ctx)
		require.Equal(t, 2, len(collections))
		require.Equal(t, resource01, collections[0].GetSymbol())
		require.Equal(t, resource02, collections[1].GetSymbol())
	}
}

func TestGetPrefixedTokens(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	const (
		symbolPrefixLink = "link"
		symbolPrefixCony = "cony"
		symbolPrefixLine = "line"
	)

	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	require.NotNil(t, ak.GetAccount(ctx, addr1))
	require.Equal(t, uint64(0), ak.GetAccount(ctx, addr1).GetAccountNumber())

	{
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixLink+"1", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixLink+"2", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixLink+"3", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixCony+"1", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixCony+"2", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixLine+"1", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
		require.NoError(t, keeper.IssueFT(ctx, types.NewFT(defaultName, symbolPrefixLine+"2", defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1))
	}
	{
		tokens := keeper.GetAllTokens(ctx)
		require.Equal(t, 7, len(tokens))
	}
}

func TestAttachDetachScenario(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	const (
		defaultTokenURI = ""
		rightSymbol     = "symbol1"
		diffSymbol      = "symbol2"
		token1Id        = "id000001"
		token2Id        = "id000002"
		token3Id        = "id000003"
		token4Id        = "id000004"
		token5Id        = "id000005"
		token6Id        = "id000006"
		token7Symbol    = rightSymbol
	)

	//
	// preparation: create account
	//
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
	}

	// prepare collection
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(rightSymbol, "name"), addr1))
	collection, err := keeper.GetCollection(ctx, rightSymbol)
	require.NoError(t, err)

	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(diffSymbol, "name"), addr1))
	collection2, err := keeper.GetCollection(ctx, diffSymbol)
	require.NoError(t, err)

	// issue 6 tokens
	// token1 = symbol1id1 by addr1
	// token2 = symbol1id2 by addr1
	// token3 = symbol1id3 by addr1
	// token4 = symbol1id4 by addr1
	// token5 = symbol2id5 by addr1
	// token6 = symbol1id6 by addr2
	// token7 = symbol1 by addr1
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token1Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token2Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token3Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token4Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection2, defaultName, token5Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token6Id, defaultTokenURI, addr2), addr2)
	_ = keeper.IssueNFT(ctx, types.NewNFT(defaultName, token7Symbol, defaultTokenURI, addr1), addr1)

	//
	// attach success cases
	//

	// attach token1 <- token2 (basic case) : success
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token2Id))

	// attach token2 <- token3 (attach to a child): success
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token2Id, token3Id))

	// attach token1 <- token4 (attach to a root): success
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token4Id))

	// verify the relations

	// root of token1 is nil
	rootOfToken1, err1 := keeper.RootOf(ctx, rightSymbol, token1Id)
	require.NoError(t, err1)
	require.Nil(t, rootOfToken1)

	// root of token2 is token1
	rootOfToken2, err2 := keeper.RootOf(ctx, rightSymbol, token2Id)
	require.NoError(t, err2)
	require.Equal(t, rootOfToken2.GetTokenID(), token1Id)

	// root of token3 is token1
	rootOfToken3, err3 := keeper.RootOf(ctx, rightSymbol, token3Id)
	require.NoError(t, err3)
	require.Equal(t, rootOfToken3.GetTokenID(), token1Id)

	// root of token4 is token1
	rootOfToken4, err4 := keeper.RootOf(ctx, rightSymbol, token4Id)
	require.NoError(t, err4)
	require.Equal(t, rootOfToken4.GetTokenID(), token1Id)

	// parent of token1 is nil
	parentOfToken1, err5 := keeper.ParentOf(ctx, rightSymbol, token1Id)
	require.NoError(t, err5)
	require.Nil(t, parentOfToken1)

	// parent of token2 is token1
	parentOfToken2, err6 := keeper.ParentOf(ctx, rightSymbol, token2Id)
	require.NoError(t, err6)
	require.Equal(t, parentOfToken2.GetTokenID(), token1Id)

	// parent of token3 is token2
	parentOfToken3, err7 := keeper.ParentOf(ctx, rightSymbol, token3Id)
	require.NoError(t, err7)
	require.Equal(t, parentOfToken3.GetTokenID(), token2Id)

	// parent of token4 is token1
	parentOfToken4, err8 := keeper.ParentOf(ctx, rightSymbol, token4Id)
	require.NoError(t, err8)
	require.Equal(t, parentOfToken4.GetTokenID(), token1Id)

	// children of token1 are token2, token4
	childrenOfToken1, err9 := keeper.ChildrenOf(ctx, rightSymbol, token1Id)
	require.NoError(t, err9)
	require.Equal(t, len(childrenOfToken1), 2)
	require.True(t, (childrenOfToken1[0].GetTokenID() == token2Id && childrenOfToken1[1].GetTokenID() == token4Id) || (childrenOfToken1[0].GetTokenID() == token4Id && childrenOfToken1[1].GetTokenID() == token2Id))

	// child of token2 is token3
	childrenOfToken2, err10 := keeper.ChildrenOf(ctx, rightSymbol, token2Id)
	require.NoError(t, err10)
	require.Equal(t, len(childrenOfToken2), 1)
	require.True(t, childrenOfToken2[0].GetTokenID() == token3Id)

	// child of token3 is empty
	childrenOfToken3, err11 := keeper.ChildrenOf(ctx, rightSymbol, token3Id)
	require.NoError(t, err11)
	require.Equal(t, len(childrenOfToken3), 0)

	// child of token4 is empty
	childrenOfToken4, err12 := keeper.ChildrenOf(ctx, rightSymbol, token4Id)
	require.NoError(t, err12)
	require.Equal(t, len(childrenOfToken4), 0)

	//
	// attach error cases
	//

	// attach non-root token : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token6Id, token2Id), types.ErrTokenAlreadyAChild(types.DefaultCodespace, rightSymbol+token2Id).Error())

	// attach non-exist token : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token5Id), types.ErrCollectionTokenNotExist(types.DefaultCodespace, rightSymbol, token5Id).Error())
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token5Id, token1Id), types.ErrCollectionTokenNotExist(types.DefaultCodespace, rightSymbol, token5Id).Error())

	// attach non-mine token : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token6Id), types.ErrTokenNotOwnedBy(types.DefaultCodespace, rightSymbol+token6Id, addr1).Error())

	// attach non-CNFT : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, ""), types.ErrTokenNotIDNF(types.DefaultCodespace, rightSymbol).Error())

	// attach to itself : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token1Id), types.ErrCannotAttachToItself(types.DefaultCodespace, rightSymbol+token1Id).Error())

	// attach to a descendant : failure
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token2Id, token1Id), types.ErrCannotAttachToADescendant(types.DefaultCodespace, rightSymbol+token1Id, rightSymbol+token2Id).Error())
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token3Id, token1Id), types.ErrCannotAttachToADescendant(types.DefaultCodespace, rightSymbol+token1Id, rightSymbol+token3Id).Error())
	require.EqualError(t, keeper.Attach(ctx, addr1, rightSymbol, token4Id, token1Id), types.ErrCannotAttachToADescendant(types.DefaultCodespace, rightSymbol+token1Id, rightSymbol+token4Id).Error())

	//
	// detach error cases
	//

	// detach not a child : failure
	require.EqualError(t, keeper.Detach(ctx, addr1, addr1, rightSymbol, token1Id), types.ErrTokenNotAChild(types.DefaultCodespace, rightSymbol+token1Id).Error())

	// detach non-mine token : failure
	require.EqualError(t, keeper.Detach(ctx, addr1, addr1, rightSymbol, token6Id), types.ErrTokenNotOwnedBy(types.DefaultCodespace, rightSymbol+token6Id, addr1).Error())

	// detach non-exist token : failure
	require.EqualError(t, keeper.Detach(ctx, addr1, addr1, rightSymbol, token5Id), types.ErrCollectionTokenNotExist(types.DefaultCodespace, rightSymbol, token5Id).Error())

	//
	// detach success cases
	//

	// detach single child
	require.NoError(t, keeper.Detach(ctx, addr1, addr1, rightSymbol, token4Id))

	// detach a child having child
	require.NoError(t, keeper.Detach(ctx, addr1, addr1, rightSymbol, token2Id))

	// detach child and transfer to other
	require.NoError(t, keeper.Detach(ctx, addr1, addr2, rightSymbol, token3Id))

	//
	// verify the relations
	//
	// parent of token2 is nil
	parentOfToken2, err6 = keeper.ParentOf(ctx, rightSymbol, token2Id)
	require.NoError(t, err6)
	require.Nil(t, parentOfToken2)

	// parent of token3 is nil
	parentOfToken3, err7 = keeper.ParentOf(ctx, rightSymbol, token3Id)
	require.NoError(t, err7)
	require.Nil(t, parentOfToken3)

	// parent of token4 is nil
	parentOfToken4, err8 = keeper.ParentOf(ctx, rightSymbol, token4Id)
	require.NoError(t, err8)
	require.Nil(t, parentOfToken4)

	// children of token1 is empty
	childrenOfToken1, err1 = keeper.ChildrenOf(ctx, rightSymbol, token1Id)
	require.NoError(t, err1)
	require.Equal(t, len(childrenOfToken1), 0)

	// owner of token3 is addr2
	token3, err13 := keeper.GetToken(ctx, rightSymbol, token3Id)
	require.NoError(t, err13)

	require.Equal(t, (token3.(types.CollectiveNFT)).GetOwner(), addr2)
}

func TestTransferCFTScenario(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	const (
		defaultTokenURI = ""
		Symbol          = "symbol1"
		tokenID         = "00000001"
	)

	//
	// preparation: create account
	//
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
	}

	// issue idf token
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(Symbol, "name"), addr1))
	collection, err := keeper.GetCollection(ctx, Symbol)
	require.NoError(t, err)
	_ = keeper.IssueFT(ctx, types.NewCollectiveFT(collection, defaultName, tokenID, defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount), addr1)

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferCFT(ctx, addr1, addr2, Symbol, tokenID, sdk.NewInt(10)))

	//
	// transfer failure cases
	//
	// Insufficient coins
	require.EqualError(t, keeper.TransferCFT(ctx, addr1, addr2, Symbol, tokenID, sdk.NewInt(defaultAmount+10)), sdk.ErrInsufficientCoins("insufficient account funds; 990symbol100000001 < 1010symbol100000001").Error())
}

func TestTransferNFTScenario(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	const (
		defaultTokenURI = ""
		Symbol          = "symbol1"
	)

	//
	// preparation: create account
	//
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
	}

	// issue nf token
	_ = keeper.IssueNFT(ctx, types.NewNFT(defaultName, Symbol, defaultTokenURI, addr1), addr1)

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferNFT(ctx, addr1, addr2, Symbol))

	//
	// transfer failure cases
	//
	// Insufficient coins
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, "Symbol2"), types.ErrTokenNotExist(types.DefaultCodespace, "Symbol2").Error())
}

func TestTransferCNFTScenario(t *testing.T) {
	input := SetupTestInput(t)
	_, ctx, keeper, ak := input.Cdc, input.Ctx, input.Keeper, input.Ak

	const (
		defaultTokenURI = ""
		rightSymbol     = "symbol1"
		diffSymbol      = "symbol2"
		token1Id        = "id000001"
		token2Id        = "id000002"
		token3Id        = "id000003"
		token4Id        = "id000004"
		token5Id        = "id000005"
		token6Id        = "id000006"
		token7Symbol    = rightSymbol
	)

	//
	// preparation: create account
	//
	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr1)
		ak.SetAccount(ctx, acc)
	}
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	{
		acc := ak.NewAccountWithAddress(ctx, addr2)
		ak.SetAccount(ctx, acc)
	}

	// prepare collection
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(rightSymbol, "name"), addr1))
	collection, err := keeper.GetCollection(ctx, rightSymbol)
	require.NoError(t, err)

	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(diffSymbol, "name"), addr1))
	collection2, err := keeper.GetCollection(ctx, diffSymbol)
	require.NoError(t, err)

	// issue 6 tokens
	// token1 = symbol1id1 by addr1
	// token2 = symbol1id2 by addr1
	// token3 = symbol1id3 by addr1
	// token4 = symbol1id4 by addr1
	// token5 = symbol2id5 by addr1
	// token6 = symbol1id6 by addr2
	// token7 = symbol1 by addr1
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token1Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token2Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token3Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token4Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection2, defaultName, token5Id, defaultTokenURI, addr1), addr1)
	_ = keeper.IssueNFT(ctx, types.NewCollectiveNFT(collection, defaultName, token6Id, defaultTokenURI, addr2), addr2)
	_ = keeper.IssueNFT(ctx, types.NewNFT(defaultName, token7Symbol, defaultTokenURI, addr1), addr1)

	// attach token1 <- token2 (basic case) : success
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token2Id))
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token2Id, token3Id))
	require.NoError(t, keeper.Attach(ctx, addr1, rightSymbol, token1Id, token4Id))

	//
	// transfer failure cases
	//

	// transfer non-exist token : failure
	require.EqualError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token5Id), types.ErrCollectionTokenNotExist(types.DefaultCodespace, rightSymbol, token5Id).Error())

	// transfer a child : failure
	require.EqualError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token2Id), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, rightSymbol+token2Id).Error())
	require.EqualError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token3Id), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, rightSymbol+token3Id).Error())
	require.EqualError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token4Id), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, rightSymbol+token4Id).Error())

	// transfer non-mine : failure
	require.EqualError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token6Id), types.ErrTokenNotOwnedBy(types.DefaultCodespace, rightSymbol+token6Id, addr1).Error())

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token1Id))
	require.NoError(t, keeper.TransferCNFT(ctx, addr2, addr1, rightSymbol, token1Id))
	require.NoError(t, keeper.TransferCNFT(ctx, addr1, addr2, rightSymbol, token1Id))

	// verify the owner of transferred tokens
	// owner of token1 is addr2
	token1, err1 := keeper.GetToken(ctx, rightSymbol, token1Id)
	require.NoError(t, err1)
	require.Equal(t, token1.(types.CollectiveNFT).GetOwner(), addr2)

	// owner of token2 is addr2
	token2, err2 := keeper.GetToken(ctx, rightSymbol, token2Id)
	require.NoError(t, err2)
	require.Equal(t, token2.(types.CollectiveNFT).GetOwner(), addr2)

	// owner of token3 is addr2
	token3, err3 := keeper.GetToken(ctx, rightSymbol, token3Id)
	require.NoError(t, err3)
	require.Equal(t, token3.(types.CollectiveNFT).GetOwner(), addr2)

	// owner of token4 is addr2
	token4, err4 := keeper.GetToken(ctx, rightSymbol, token4Id)
	require.NoError(t, err4)
	require.Equal(t, token4.(types.CollectiveNFT).GetOwner(), addr2)
}

// This test is from cosmos/x/bank/internal/keeper_test.go
func TestTransferFT(t *testing.T) {
	input := SetupTestInput(t)
	ctx := input.Ctx

	addr := sdk.AccAddress([]byte("addr1"))
	addr2 := sdk.AccAddress([]byte("addr2"))
	acc := input.Ak.NewAccountWithAddress(ctx, addr)

	// Test GetCoins/SetCoins
	input.Ak.SetAccount(ctx, acc)
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr).IsEqual(sdk.NewCoins()))

	require.NoError(t, input.Keeper.bankKeeper.SetCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10))))
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10))))

	// Test HasCoins
	require.True(t, input.Keeper.bankKeeper.HasCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10))))
	require.True(t, input.Keeper.bankKeeper.HasCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 5))))
	require.False(t, input.Keeper.bankKeeper.HasCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 15))))
	require.False(t, input.Keeper.bankKeeper.HasCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("barcoin", 5))))

	require.NoError(t, input.Keeper.bankKeeper.SetCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 15))))

	// Test SendCoins
	_, err := input.Keeper.bankKeeper.SubtractCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 5)))
	require.NoError(t, err)
	_, err = input.Keeper.bankKeeper.AddCoins(ctx, addr2, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 5)))
	require.NoError(t, err)
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10))))
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr2).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 5))))

	_, err = input.Keeper.bankKeeper.SubtractCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 50)))
	require.Implements(t, (*sdk.Error)(nil), err)
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 10))))
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr2).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("foocoin", 5))))

	_, err = input.Keeper.bankKeeper.AddCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("barcoin", 30)))
	require.NoError(t, err)
	_, err = input.Keeper.bankKeeper.SubtractCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("barcoin", 10), sdk.NewInt64Coin("foocoin", 5)))
	require.NoError(t, err)
	_, err = input.Keeper.bankKeeper.AddCoins(ctx, addr2, sdk.NewCoins(sdk.NewInt64Coin("barcoin", 10), sdk.NewInt64Coin("foocoin", 5)))
	require.NoError(t, err)
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("barcoin", 20), sdk.NewInt64Coin("foocoin", 5))))
	require.True(t, input.Keeper.bankKeeper.GetCoins(ctx, addr2).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("barcoin", 10), sdk.NewInt64Coin("foocoin", 10))))

	// validate coins with invalid denoms or negative values cannot be sent
	// NOTE: We must use the Coin literal as the constructor does not allow
	// negative values.
	_, err = input.Keeper.bankKeeper.SubtractCoins(ctx, addr, sdk.Coins{sdk.Coin{Denom: "FOOCOIN", Amount: sdk.NewInt(-5)}})
	require.Error(t, err)
}
