package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/stretchr/testify/require"
)

const testName = "test"

func TestModuleAccountMarshal(t *testing.T) {
	moduleAcc := NewEmptyModuleAccount(testName, Minter, Burner, Staking)

	// need a codec for marshaling
	cdc := codec.New()
	codec.RegisterCrypto(cdc)

	exp, err := cdc.MarshalBinaryBare(moduleAcc)
	require.Nil(t, err)

	b, err := moduleAcc.MarshalAminoBare(false)
	require.Nil(t, err)
	require.Equal(t, exp, b)

	moduleAcc2 := ModuleAccount{}
	err = cdc.UnmarshalBinaryBare(b, &moduleAcc2)
	require.Nil(t, err)
	require.Equal(t, *moduleAcc, moduleAcc2)

	// error on bad bytes
	moduleAcc2 = ModuleAccount{}
	err = cdc.UnmarshalBinaryBare(b[:len(b)/2], &moduleAcc2)
	require.NotNil(t, err)
}

func TestBaseAccountUnmarshal(t *testing.T) {
	moduleAcc := NewEmptyModuleAccount(testName, Minter, Burner, Staking)

	// need a codec for marshaling
	cdc := codec.New()
	codec.RegisterCrypto(cdc)

	b, err := cdc.MarshalBinaryBare(moduleAcc)
	require.Nil(t, err)

	moduleAcc2 := ModuleAccount{}
	err = cdc.UnmarshalBinaryBare(b, &moduleAcc2)
	require.Nil(t, err)
	require.Equal(t, *moduleAcc, moduleAcc2)

	moduleAcc3 := ModuleAccount{}
	_, err = moduleAcc3.UnmarshalAminoBare(b)
	require.Nil(t, err)
	require.Equal(t, *moduleAcc, moduleAcc3)
}

func TestModuleAccountMarshalYAML(t *testing.T) {
	moduleAcc := NewEmptyModuleAccount(testName, Minter, Burner, Staking)
	bs, err := yaml.Marshal(moduleAcc)
	require.NoError(t, err)

	want := "|\n  address: cosmos1n7rdpqvgf37ktx30a2sv2kkszk3m7ncmg5drhe\n  coins: []\n  public_key: \"\"\n  account_number: 0\n  sequence: 0\n  name: test\n  permissions:\n  - minter\n  - burner\n  - staking\n"
	require.Equal(t, want, string(bs))
}

func TestHasPermissions(t *testing.T) {
	macc := NewEmptyModuleAccount(testName, Staking, Minter, Burner)
	cases := []struct {
		permission string
		expectHas  bool
	}{
		{Staking, true},
		{Minter, true},
		{Burner, true},
		{"other", false},
	}

	for i, tc := range cases {
		hasPerm := macc.HasPermission(tc.permission)
		if tc.expectHas {
			require.True(t, hasPerm, "test case #%d", i)
		} else {
			require.False(t, hasPerm, "test case #%d", i)
		}
	}
}

func TestValidate(t *testing.T) {
	addr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	baseAcc := authtypes.NewBaseAccount(addr, sdk.Coins{}, nil, 0, 0)
	tests := []struct {
		name   string
		acc    authexported.GenesisAccount
		expErr error
	}{
		{
			"valid module account",
			NewEmptyModuleAccount("test"),
			nil,
		},
		{
			"invalid name and address pair",
			NewModuleAccount(baseAcc, "test"),
			fmt.Errorf("address %s cannot be derived from the module name 'test'", addr),
		},
		{
			"empty module account name",
			NewModuleAccount(baseAcc, "    "),
			errors.New("module account name cannot be blank"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.acc.Validate()
			require.Equal(t, tt.expErr, err)
		})
	}
}

func TestModuleAccountJSON(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	coins := sdk.NewCoins(sdk.NewInt64Coin("test", 5))
	baseAcc := authtypes.NewBaseAccount(addr, coins, nil, 10, 50)
	acc := NewModuleAccount(baseAcc, "test", "burner")

	bz, err := json.Marshal(acc)
	require.NoError(t, err)

	bz1, err := acc.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, string(bz1), string(bz))

	var a ModuleAccount
	require.NoError(t, json.Unmarshal(bz, &a))
	require.Equal(t, acc.String(), a.String())
}
