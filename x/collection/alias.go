package collection

import (
	"github.com/line/link/x/collection/internal/keeper"
	"github.com/line/link/x/collection/internal/types"
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	DefaultParamspace = types.DefaultParamspace
)

type (
	Token  = types.Token
	Tokens = types.Tokens

	FT  = types.FT
	NFT = types.NFT

	TokenType = types.TokenType

	Collection = types.BaseCollection

	MsgCreateCollection = types.MsgCreateCollection
	MsgIssueFT          = types.MsgIssueFT
	MsgIssueNFT         = types.MsgIssueNFT
	MsgMintNFT          = types.MsgMintNFT
	MsgBurnNFT          = types.MsgBurnNFT
	MsgBurnNFTFrom      = types.MsgBurnNFTFrom
	MsgMintFT           = types.MsgMintFT
	MsgBurnFT           = types.MsgBurnFT
	MsgBurnFTFrom       = types.MsgBurnFTFrom

	MsgGrantPermission  = types.MsgGrantPermission
	MsgRevokePermission = types.MsgRevokePermission

	MsgModify = types.MsgModify

	MsgTransferFT  = types.MsgTransferFT
	MsgTransferNFT = types.MsgTransferNFT

	MsgTransferFTFrom  = types.MsgTransferFTFrom
	MsgTransferNFTFrom = types.MsgTransferNFTFrom

	MsgAttach = types.MsgAttach
	MsgDetach = types.MsgDetach

	MsgAttachFrom = types.MsgAttachFrom
	MsgDetachFrom = types.MsgDetachFrom

	MsgApproveCollection    = types.MsgApprove
	MsgDisapproveCollection = types.MsgDisapprove

	Permissions = types.Permissions
	Permission  = types.Permission

	MintNFTParam = types.MintNFTParam

	Keeper = keeper.Keeper
)

var (
	NewMsgCreateCollection = types.NewMsgCreateCollection
	NewMsgIssueFT          = types.NewMsgIssueFT
	NewMsgIssueNFT         = types.NewMsgIssueNFT
	NewMsgMintNFT          = types.NewMsgMintNFT
	NewMsgBurnNFT          = types.NewMsgBurnNFT
	NewMsgMintFT           = types.NewMsgMintFT
	NewMsgBurnFT           = types.NewMsgBurnFT
	NewMsgGrantPermission  = types.NewMsgGrantPermission
	NewMsgModify           = types.NewMsgModify
	NewChangesWithMap      = types.NewChangesWithMap
	NewMsgTransferFT       = types.NewMsgTransferFT
	NewMsgTransferNFT      = types.NewMsgTransferNFT
	NewMsgTransferNFTFrom  = types.NewMsgTransferNFTFrom
	NewMsgAttach           = types.NewMsgAttach
	NewMsgDetach           = types.NewMsgDetach
	NewMsgApprove          = types.NewMsgApprove
	NewMintNFTParam        = types.NewMintNFTParam
	NewCoin                = types.NewCoin
	NewPermissions         = types.NewPermissions

	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewKeeper = keeper.NewKeeper
)
