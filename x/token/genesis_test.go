package token_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/line/lbm-sdk/crypto/keys/secp256k1"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/token"
)

func TestValidateGenesis(t *testing.T) {
	addr := sdk.BytesToAccAddress(secp256k1.GenPrivKey().PubKey().Address())
	testCases := map[string]struct {
		gs    *token.GenesisState
		valid bool
	}{
		"default genesis": {
			token.DefaultGenesisState(),
			true,
		},
		"invalid class nonce": {
			&token.GenesisState{
				ClassState: &token.ClassGenesisState{
					Nonce: sdk.NewUint(math.MaxUint64).Incr(),
				},
			},
			false,
		},
		"invalid address in a balance": {
			&token.GenesisState{
				Balances: []token.ContractBalances{
					{
						ContractId: "deadbeef",
						Balances: []token.Balance{
							{
								Amount:  sdk.OneInt(),
							},
						},
					},
				},
			},
			false,
		},
		"invalid amount in a balance": {
			&token.GenesisState{
				Balances: []token.ContractBalances{
					{
						ContractId: "deadbeef",
						Balances: []token.Balance{
							{
								Address: addr.String(),
								Amount:  sdk.ZeroInt(),
							},
						},
					},
				},
			},
			false,
		},
		"empty tokens in a balance": {
			&token.GenesisState{
				Balances: []token.ContractBalances{
					{
						ContractId: "deadbeef",
					},
				},
			},
			false,
		},
		"invalid id of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:     "invalid",
					Name:   "test",
					Symbol: "TT",
				}},
			},
			false,
		},
		"invalid name of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:     "deadbeef",
					Name:   string(make([]rune, 21)),
					Symbol: "TT",
				}},
			},
			false,
		},
		"invalid symbol of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:     "deadbeef",
					Name:   "test",
					Symbol: "tt",
				}},
			},
			false,
		},
		"invalid image uri of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:       "deadbeef",
					Name:     "test",
					Symbol:   "TT",
					ImageUri: string(make([]rune, 1001)),
				}},
			},
			false,
		},
		"invalid meta of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:     "deadbeef",
					Name:   "test",
					Symbol: "TT",
					Meta:   string(make([]rune, 1001)),
				}},
			},
			false,
		},
		"invalid decimals of class": {
			&token.GenesisState{
				Classes: []token.TokenClass{{
					ContractId:       "deadbeef",
					Name:     "test",
					Symbol:   "TT",
					Decimals: -1,
				}},
			},
			false,
		},
		"empty grants": {
			&token.GenesisState{
				Grants: []token.ContractGrants{{
					ContractId: "deadbeef",
				}},
			},
			false,
		},
		"invalid grantee of grant": {
			&token.GenesisState{
				Grants: []token.ContractGrants{{
					ContractId: "deadbeef",
					Grants: []token.Grant{{
						Permission: token.Permission_Mint.String(),
					}},
				}},
			},
			false,
		},
		"invalid action of grant": {
			&token.GenesisState{
				Grants: []token.ContractGrants{{
					ContractId: "deadbeef",
					Grants: []token.Grant{{
						Grantee: addr.String(),
					}},
				}},
			},
			false,
		},
		"empty authorizations": {
			&token.GenesisState{
				Authorizations: []token.ContractAuthorizations{{
					ContractId: "deadbeef",
				}},
			},
			false,
		},
		"invalid holder of authorization": {
			&token.GenesisState{
				Authorizations: []token.ContractAuthorizations{{
					ContractId: "deadbeef",
					Authorizations: []token.Authorization{{
						Operator:    addr.String(),
					}},
				}},
			},
			false,
		},
		"invalid operator of authorization": {
			&token.GenesisState{
				Authorizations: []token.ContractAuthorizations{{
					ContractId: "deadbeef",
					Authorizations: []token.Authorization{{
						Holder: addr.String(),
					}},
				}},
			},
			false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := token.ValidateGenesis(*tc.gs)
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
