package collection

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

// ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	for _, contractBalances := range data.Balances {
		// TODO: validate contract id
		if len(contractBalances.Balances) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("balances cannot be empty")
		}
		for _, balance := range contractBalances.Balances {
			if err := sdk.ValidateAccAddress(balance.Address); err != nil {
				return err
			}
			if err := balance.Amount.ValidateBasic(); err != nil {
				return err
			}
		}
	}

	// for _, c := range data.Classes {
	// 	if err := class.ValidateID(c.ContractId); err != nil {
	// 		return err
	// 	}
	// 	if err := validateName(c.Name); err != nil {
	// 		return err
	// 	}
	// 	if err := validateSymbol(c.Symbol); err != nil {
	// 		return err
	// 	}
	// 	if err := validateImageURI(c.ImageUri); err != nil {
	// 		return err
	// 	}
	// 	if err := validateMeta(c.Meta); err != nil {
	// 		return err
	// 	}
	// 	if err := validateDecimals(c.Decimals); err != nil {
	// 		return err
	// 	}
	// }

	for _, contractGrants := range data.Grants {
		// TODO: validate contract id
		if len(contractGrants.Grants) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("grants cannot be empty")
		}
		for _, grant := range contractGrants.Grants {
			if err := sdk.ValidateAccAddress(grant.Grantee); err != nil {
				return err
			}
			// if err := validatePermission(grant.Permission); err != nil {
			// 	return err
			// }
		}
	}

	for _, contractAuthorizations := range data.Authorizations {
		// TODO: validate contract id
		if len(contractAuthorizations.Authorizations) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("authorizations cannot be empty")
		}
		for _, authorization := range contractAuthorizations.Authorizations {
			if err := sdk.ValidateAccAddress(authorization.Approver); err != nil {
				return err
			}
			if err := sdk.ValidateAccAddress(authorization.Proxy); err != nil {
				return err
			}
		}
	}

	return nil
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
