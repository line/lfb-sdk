package keeper_test

import (
	"fmt"
	"context"
	"testing"

	ocproto "github.com/line/ostracon/proto/ostracon/types"

	"github.com/stretchr/testify/suite"

	"github.com/line/lbm-sdk/crypto/keys/secp256k1"
	"github.com/line/lbm-sdk/simapp"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/collection"
	"github.com/line/lbm-sdk/x/collection/keeper"
)

type KeeperTestSuite struct {
	suite.Suite
	ctx         sdk.Context
	goCtx       context.Context
	keeper      keeper.Keeper
	queryServer collection.QueryServer
	msgServer   collection.MsgServer

	vendor   sdk.AccAddress
	operator sdk.AccAddress
	customer sdk.AccAddress
	stranger sdk.AccAddress

	contractID string
	ftClassID string
	nftClassID string

	balance sdk.Int
}

func createRandomAccounts(accNum int) []sdk.AccAddress {
	seenAddresses := make(map[sdk.AccAddress]bool, accNum)
	addresses := make([]sdk.AccAddress, accNum)
	for i := 0; i < accNum; i++ {
		var address sdk.AccAddress
		for {
			pk := secp256k1.GenPrivKey().PubKey()
			address = sdk.BytesToAccAddress(pk.Address())
			if !seenAddresses[address] {
				seenAddresses[address] = true
				break
			}
		}
		addresses[i] = address
	}
	return addresses
}

func (s *KeeperTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)
	s.ctx = app.BaseApp.NewContext(checkTx, ocproto.Header{})
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	s.keeper = app.CollectionKeeper

	// s.queryServer = keeper.NewQueryServer(s.keeper)
	s.msgServer = keeper.NewMsgServer(s.keeper)

	addresses := []*sdk.AccAddress{
		&s.vendor,
		&s.operator,
		&s.customer,
		&s.stranger,
	}
	for i, address := range createRandomAccounts(len(addresses)) {
		*addresses[i] = address
	}

	s.balance = sdk.OneInt()

	// create a contract
	contractID, err := s.keeper.CreateContract(s.ctx, s.vendor, collection.Contract{
		Name: "test contract",
	})
	s.Require().NoError(err)
	s.contractID = *contractID

	for _, permission := range []collection.Permission{
		collection.Permission_Mint,
		collection.Permission_Burn,
	}{
		err := s.keeper.Grant(s.ctx, s.contractID, s.vendor, s.operator, permission)
		s.Require().NoError(err)
	}

	// create a fungible token class
	ftClassID, err := s.keeper.CreateTokenClass(s.ctx, &collection.FTClass{
		ContractId: s.contractID,
		Name: "test ft class",
	})
	s.Require().NoError(err)
	s.ftClassID = *ftClassID

	// create a non-fungible token class
	nftClassID, err := s.keeper.CreateTokenClass(s.ctx, &collection.NFTClass{
		ContractId: s.contractID,
		Name: "test ft class",
	})
	s.Require().NoError(err)
	s.nftClassID = *nftClassID

	// set balances
	// TODO: replace with mint
	for _, to := range []sdk.AccAddress{s.vendor, s.operator, s.customer} {
		s.keeper.SetBalance(s.ctx, s.contractID, to, s.ftClassID + fmt.Sprintf("%08x", 0), s.balance)
	}
	s.keeper.SetBalance(s.ctx, s.contractID, s.customer, s.nftClassID + fmt.Sprintf("%08x", 1), s.balance)

	// authorize
	err = s.keeper.AuthorizeOperator(s.ctx, s.contractID, s.customer, s.operator)
	s.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
