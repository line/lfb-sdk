package token

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

var _ sdk.Msg = (*MsgSend)(nil)

// ValidateBasic implements Msg.
func (m MsgSend) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgSend) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.From)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgSend) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgSend) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgOperatorSend)(nil)

// ValidateBasic implements Msg.
func (m MsgOperatorSend) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgOperatorSend) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgOperatorSend) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgOperatorSend) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgOperatorSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgRevokeOperator)(nil)

// ValidateBasic implements Msg.
func (m MsgRevokeOperator) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid holder address: %s", m.Holder)
	}
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgRevokeOperator) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgRevokeOperator) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgRevokeOperator) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgRevokeOperator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgAuthorizeOperator)(nil)

// ValidateBasic implements Msg.
func (m MsgAuthorizeOperator) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid holder address: %s", m.Holder)
	}
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgAuthorizeOperator) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgAuthorizeOperator) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgAuthorizeOperator) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgAuthorizeOperator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgIssue)(nil)

// ValidateBasic implements Msg.
func (m MsgIssue) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", m.Owner)
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateName(m.Name); err != nil {
		return err
	}

	if err := validateSymbol(m.Symbol); err != nil {
		return err
	}

	if err := validateImageURI(m.Uri); err != nil {
		return err
	}

	if err := validateMeta(m.Meta); err != nil {
		return err
	}

	if err := validateDecimals(m.Decimals); err != nil {
		return err
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgIssue) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgIssue) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgIssue) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgIssue) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgGrantPermission)(nil)

// ValidateBasic implements Msg.
func (m MsgGrantPermission) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid granter address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.To)
	}

	if err := validateLegacyPermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgGrantPermission) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.From)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgGrantPermission) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgGrantPermission) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgGrantPermission) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgRevokePermission)(nil)

// ValidateBasic implements Msg.
func (m MsgRevokePermission) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateLegacyPermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgRevokePermission) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.From)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgRevokePermission) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgRevokePermission) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgRevokePermission) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgMint)(nil)

// ValidateBasic implements Msg.
func (m MsgMint) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.From)
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgMint) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.From)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgMint) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgMint) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgBurn)(nil)

// ValidateBasic implements Msg.
func (m MsgBurn) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgBurn) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.From)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgBurn) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgBurn) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgOperatorBurn)(nil)

// ValidateBasic implements Msg.
func (m MsgOperatorBurn) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgOperatorBurn) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgOperatorBurn) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgOperatorBurn) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgOperatorBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgModify)(nil)

// ValidateBasic implements Msg.
func (m MsgModify) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.Owner)
	}

	checkedFields := map[string]bool{}
	for _, change := range m.Changes {
		if checkedFields[change.Field] {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate fields: %s", change.Field)
		}
		checkedFields[change.Field] = true

		if err := validateChange(change); err != nil {
			return err
		}
	}
	if len(checkedFields) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("no field provided")
	}

	return nil
}

// GetSigners implements Msg
func (m MsgModify) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgModify) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgModify) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgModify) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
