package keeper_test

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/Finschia/finschia-sdk/x/collection"
)

func (s *KeeperTestSuite) TestQueryBalance() {
	// empty request
	_, err := s.queryServer.Balance(s.ctx, nil)
	s.Require().Error(err)

	tokenID := s.issuedNFTs[s.bytesToString(s.customer)][0].TokenId
	testCases := map[string]struct {
		contractID string
		address    sdk.AccAddress
		tokenID    string
		valid      bool
		postTest   func(res *collection.QueryBalanceResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			address:    s.customer,
			tokenID:    tokenID,
			valid:      true,
			postTest: func(res *collection.QueryBalanceResponse) {
				expected := collection.Coin{
					TokenId: tokenID,
					Amount:  math.OneInt(),
				}
				fmt.Println(res.Balance.Amount)
				s.Require().Equal(expected, res.Balance)
			},
		},
		"valid request with zero amount": {
			contractID: s.contractID,
			address:    s.stranger,
			tokenID:    tokenID,
			valid:      true,
			postTest: func(res *collection.QueryBalanceResponse) {
				expected := collection.Coin{
					TokenId: tokenID,
					Amount:  math.ZeroInt(),
				}
				s.Require().Equal(expected, res.Balance)
			},
		},
		"invalid contract id": {
			address: s.vendor,
			tokenID: tokenID,
		},
		"invalid address": {
			contractID: s.contractID,
			tokenID:    tokenID,
		},
		"invalid token id": {
			contractID: s.contractID,
			address:    s.vendor,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryBalanceRequest{
				ContractId: tc.contractID,
				Address:    s.bytesToString(tc.address),
				TokenId:    tc.tokenID,
			}
			res, err := s.queryServer.Balance(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryAllBalances() {
	// empty request
	_, err := s.queryServer.AllBalances(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		address    sdk.AccAddress
		valid      bool
		count      uint64
		postTest   func(res *collection.QueryAllBalancesResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			address:    s.customer,
			valid:      true,
			postTest: func(res *collection.QueryAllBalancesResponse) {
				s.Require().Equal(s.numNFTs, len(res.Balances))
			},
		},
		"valid request with limit": {
			contractID: s.contractID,
			address:    s.customer,
			valid:      true,
			count:      1,
			postTest: func(res *collection.QueryAllBalancesResponse) {
				s.Require().Equal(1, len(res.Balances))
			},
		},
		"invalid contract id": {
			address: s.customer,
		},
		"invalid address": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			pageReq := &query.PageRequest{}
			if tc.count != 0 {
				pageReq.Limit = tc.count
			}
			req := &collection.QueryAllBalancesRequest{
				ContractId: tc.contractID,
				Address:    s.bytesToString(tc.address),
				Pagination: pageReq,
			}
			res, err := s.queryServer.AllBalances(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryNFTSupply() {
	// empty request
	_, err := s.queryServer.NFTSupply(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		tokenType  string
		valid      bool
		postTest   func(res *collection.QueryNFTSupplyResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTSupplyResponse) {
				s.Require().EqualValues(s.numNFTs*3, res.Supply.Int64())
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTSupplyResponse) {
				s.Require().Equal(math.ZeroInt(), res.Supply)
			},
		},
		"token type not found": {
			contractID: s.contractID,
			tokenType:  "deadbeef",
			valid:      true,
			postTest: func(res *collection.QueryNFTSupplyResponse) {
				s.Require().Equal(math.ZeroInt(), res.Supply)
			},
		},
		"invalid contract id": {
			tokenType: s.nftClassID,
		},
		"invalid token type": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryNFTSupplyRequest{
				ContractId: tc.contractID,
				TokenType:  tc.tokenType,
			}
			res, err := s.queryServer.NFTSupply(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryNFTMinted() {
	// empty request
	_, err := s.queryServer.NFTMinted(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		tokenType  string
		valid      bool
		postTest   func(res *collection.QueryNFTMintedResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTMintedResponse) {
				s.Require().EqualValues(s.numNFTs*3, res.Minted.Int64())
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTMintedResponse) {
				s.Require().Equal(math.ZeroInt(), res.Minted)
			},
		},
		"token type not found": {
			contractID: s.contractID,
			tokenType:  "deadbeef",
			valid:      true,
			postTest: func(res *collection.QueryNFTMintedResponse) {
				s.Require().Equal(math.ZeroInt(), res.Minted)
			},
		},
		"invalid contract id": {
			tokenType: s.nftClassID,
		},
		"invalid token type": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryNFTMintedRequest{
				ContractId: tc.contractID,
				TokenType:  tc.tokenType,
			}
			res, err := s.queryServer.NFTMinted(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryNFTBurnt() {
	// empty request
	_, err := s.queryServer.NFTBurnt(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		tokenType  string
		valid      bool
		postTest   func(res *collection.QueryNFTBurntResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTBurntResponse) {
				s.Require().Equal(math.ZeroInt(), res.Burnt)
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryNFTBurntResponse) {
				s.Require().Equal(math.ZeroInt(), res.Burnt)
			},
		},
		"token type not found": {
			contractID: s.contractID,
			tokenType:  "deadbeef",
			valid:      true,
			postTest: func(res *collection.QueryNFTBurntResponse) {
				s.Require().Equal(math.ZeroInt(), res.Burnt)
			},
		},
		"invalid contract id": {
			tokenType: s.nftClassID,
		},
		"invalid token type": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryNFTBurntRequest{
				ContractId: tc.contractID,
				TokenType:  tc.tokenType,
			}
			res, err := s.queryServer.NFTBurnt(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryContract() {
	// empty request
	_, err := s.queryServer.Contract(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		valid      bool
		postTest   func(res *collection.QueryContractResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			valid:      true,
			postTest: func(res *collection.QueryContractResponse) {
				s.Require().Equal(s.contractID, res.Contract.Id)
			},
		},
		"invalid contract id": {},
		"no such an id": {
			contractID: "deadbeef",
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryContractRequest{
				ContractId: tc.contractID,
			}
			res, err := s.queryServer.Contract(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryTokenClassTypeName() {
	// empty request
	_, err := s.queryServer.TokenClassTypeName(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		classID    string
		valid      bool
		postTest   func(res *collection.QueryTokenClassTypeNameResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			classID:    s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryTokenClassTypeNameResponse) {
				s.Require().Equal(proto.MessageName(&collection.NFTClass{}), res.Name)
			},
		},
		"invalid contract id": {
			classID: s.nftClassID,
		},
		"invalid class id": {
			contractID: s.contractID,
		},
		"collection not found": {
			contractID: "deadbeef",
			classID:    s.nftClassID,
		},
		"class not found": {
			contractID: s.contractID,
			classID:    "00bab10c",
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryTokenClassTypeNameRequest{
				ContractId: tc.contractID,
				ClassId:    tc.classID,
			}
			res, err := s.queryServer.TokenClassTypeName(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryTokenType() {
	// empty request
	_, err := s.queryServer.TokenType(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		tokenType  string
		valid      bool
		postTest   func(res *collection.QueryTokenTypeResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			tokenType:  s.nftClassID,
			valid:      true,
			postTest: func(res *collection.QueryTokenTypeResponse) {
				s.Require().Equal(s.contractID, res.TokenType.ContractId)
				s.Require().Equal(s.nftClassID, res.TokenType.TokenType)
			},
		},
		"invalid contract id": {
			tokenType: s.nftClassID,
		},
		"invalid token type": {
			contractID: s.contractID,
		},
		"collection not found": {
			contractID: "deadbeef",
			tokenType:  s.nftClassID,
		},
		"token type not found": {
			contractID: s.contractID,
			tokenType:  "deadbeef",
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryTokenTypeRequest{
				ContractId: tc.contractID,
				TokenType:  tc.tokenType,
			}
			res, err := s.queryServer.TokenType(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryToken() {
	// empty request
	_, err := s.queryServer.Token(s.ctx, nil)
	s.Require().Error(err)

	nftTokenID := collection.NewNFTID(s.nftClassID, 1)
	testCases := map[string]struct {
		contractID string
		tokenID    string
		valid      bool
		postTest   func(res *collection.QueryTokenResponse)
	}{
		"valid nft request": {
			contractID: s.contractID,
			tokenID:    nftTokenID,
			valid:      true,
			postTest: func(res *collection.QueryTokenResponse) {
				s.Require().Equal("/lbm.collection.v1.OwnerNFT", res.Token.TypeUrl)
				token := collection.TokenFromAny(&res.Token)
				nft, ok := token.(*collection.OwnerNFT)
				s.Require().True(ok)
				s.Require().Equal(s.contractID, nft.ContractId)
				s.Require().Equal(nftTokenID, nft.TokenId)
			},
		},
		"invalid contract id": {
			tokenID: nftTokenID,
		},
		"invalid token id": {
			contractID: s.contractID,
		},
		"collection not found": {
			contractID: "deadbeef",
			tokenID:    nftTokenID,
		},
		"no such a non-fungible token": {
			contractID: s.contractID,
			tokenID:    collection.NewNFTID("deadbeef", 1),
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryTokenRequest{
				ContractId: tc.contractID,
				TokenId:    tc.tokenID,
			}
			res, err := s.queryServer.Token(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryGranteeGrants() {
	// empty request
	_, err := s.queryServer.GranteeGrants(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		grantee    sdk.AccAddress
		valid      bool
		postTest   func(res *collection.QueryGranteeGrantsResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			grantee:    s.vendor,
			valid:      true,
			postTest: func(res *collection.QueryGranteeGrantsResponse) {
				s.Require().Equal(4, len(res.Grants))
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			grantee:    s.vendor,
			valid:      true,
			postTest: func(res *collection.QueryGranteeGrantsResponse) {
				s.Require().Equal(0, len(res.Grants))
			},
		},
		"invalid contract id": {
			grantee: s.vendor,
		},
		"invalid grantee": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryGranteeGrantsRequest{
				ContractId: tc.contractID,
				Grantee:    s.bytesToString(tc.grantee),
			}
			res, err := s.queryServer.GranteeGrants(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryIsOperatorFor() {
	// empty request
	_, err := s.queryServer.IsOperatorFor(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		operator   sdk.AccAddress
		holder     sdk.AccAddress
		valid      bool
		postTest   func(res *collection.QueryIsOperatorForResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			operator:   s.operator,
			holder:     s.customer,
			valid:      true,
			postTest: func(res *collection.QueryIsOperatorForResponse) {
				s.Require().True(res.Authorized)
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			operator:   s.operator,
			holder:     s.vendor,
			valid:      true,
			postTest: func(res *collection.QueryIsOperatorForResponse) {
				s.Require().False(res.Authorized)
			},
		},
		"invalid contract id": {
			operator: s.operator,
			holder:   s.customer,
		},
		"invalid operator": {
			contractID: s.contractID,
			holder:     s.customer,
		},
		"invalid holder": {
			contractID: s.contractID,
			operator:   s.operator,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			req := &collection.QueryIsOperatorForRequest{
				ContractId: tc.contractID,
				Operator:   s.bytesToString(tc.operator),
				Holder:     s.bytesToString(tc.holder),
			}
			res, err := s.queryServer.IsOperatorFor(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestQueryHoldersByOperator() {
	// empty request
	_, err := s.queryServer.HoldersByOperator(s.ctx, nil)
	s.Require().Error(err)

	testCases := map[string]struct {
		contractID string
		operator   sdk.AccAddress
		valid      bool
		count      uint64
		postTest   func(res *collection.QueryHoldersByOperatorResponse)
	}{
		"valid request": {
			contractID: s.contractID,
			operator:   s.operator,
			valid:      true,
			postTest: func(res *collection.QueryHoldersByOperatorResponse) {
				s.Require().Equal(1, len(res.Holders))
			},
		},
		"valid request with limit": {
			contractID: s.contractID,
			operator:   s.operator,
			valid:      true,
			count:      1,
			postTest: func(res *collection.QueryHoldersByOperatorResponse) {
				s.Require().Equal(1, len(res.Holders))
			},
		},
		"collection not found": {
			contractID: "deadbeef",
			operator:   s.operator,
			valid:      true,
			postTest: func(res *collection.QueryHoldersByOperatorResponse) {
				s.Require().Equal(0, len(res.Holders))
			},
		},
		"invalid contract id": {
			operator: s.operator,
		},
		"invalid operator": {
			contractID: s.contractID,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			pageReq := &query.PageRequest{}
			if tc.count != 0 {
				pageReq.Limit = tc.count
			}
			req := &collection.QueryHoldersByOperatorRequest{
				ContractId: tc.contractID,
				Operator:   s.bytesToString(tc.operator),
				Pagination: pageReq,
			}
			res, err := s.queryServer.HoldersByOperator(s.ctx, req)
			if !tc.valid {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(res)
			tc.postTest(res)
		})
	}
}
