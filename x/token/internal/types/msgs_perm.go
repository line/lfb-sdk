package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/types"
)

var _ sdk.Msg = (*MsgGrantPermission)(nil)

func NewMsgGrantPermission(from, to sdk.AccAddress, perm Permission) MsgGrantPermission {
	return MsgGrantPermission{
		From:       from,
		To:         to,
		Permission: perm,
	}
}

type MsgGrantPermission struct {
	From       sdk.AccAddress `json:"from"`
	To         sdk.AccAddress `json:"to"`
	Permission Permission     `json:"permission"`
}

func (MsgGrantPermission) Route() string                    { return RouterKey }
func (MsgGrantPermission) Type() string                     { return "grant_permission" }
func (msg MsgGrantPermission) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.From} }
func (msg MsgGrantPermission) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgGrantPermission) ValidateBasic() sdk.Error {
	if msg.From.Empty() || msg.To.Empty() {
		return sdk.ErrInvalidAddress("addresses cannot be empty")
	}

	if msg.From.Equals(msg.To) {
		return sdk.ErrInvalidAddress("from, to address can not be the same")
	}

	if len(msg.Permission.GetAction()) <= 0 || len(msg.Permission.GetResource()) <= 0 {
		return types.ErrInvalidPermission("resource and action should not be empty")
	}
	return nil
}

var _ sdk.Msg = (*MsgRevokePermission)(nil)

func NewMsgRevokePermission(from sdk.AccAddress, perm Permission) MsgRevokePermission {
	return MsgRevokePermission{
		From:       from,
		Permission: perm,
	}
}

type MsgRevokePermission struct {
	From       sdk.AccAddress `json:"from"`
	Permission Permission     `json:"permission"`
}

func (MsgRevokePermission) Route() string                    { return RouterKey }
func (MsgRevokePermission) Type() string                     { return "revoke_permission" }
func (msg MsgRevokePermission) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.From} }
func (msg MsgRevokePermission) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRevokePermission) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress("address cannot be empty")
	}

	if len(msg.Permission.GetAction()) <= 0 || len(msg.Permission.GetResource()) <= 0 {
		return types.ErrInvalidPermission("resource and action should not be empty")
	}
	return nil
}
